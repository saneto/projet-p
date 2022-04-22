package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"encoding/json"
	"net/http"
	"github.com/olahol/melody"
	"github.com/gin-gonic/gin"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/auth"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/models"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/utils/formaterror"
	"github.com/olivere/elastic"
	"time"
	"os"
	"log"
	"strings"
)

func (server *Server) getUserChat(c *gin.Context) {	
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	
	user := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		errList["Unauthorized"] = "User not found"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	chats := []models.Chat{}

	
	err = server.DB.Model(&user).Related(&chats,  "Chats").Error
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		errList = formattedError
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	err =server.DB.Debug().Preload("Messages").Preload("Users").Find(&chats).Error
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		errList = formattedError
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": chats,
	})
}


func (server *Server) createChat(c *gin.Context) {
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	// check if the user exists
	user := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		errList["Unauthorized"] = "Unauthorized 2"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}


	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	requestBody := map[string]string{}
	_ = json.Unmarshal(body, &requestBody)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
		return
	}

	user2 := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("email = ?", requestBody["email"]).Take(&user2).Error
	if err != nil {
		errList["Unmarshal_error"] = "Not found"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	chat := models.Chat{}
	err = server.DB.Debug().Where("name = ? or name = ? ", user2.Email+user.Email,  user.Email+user2.Email).Find(&chat).Error
	fmt.Println(err)
	if err != nil {
		chat.Name = user.Email+user2.Email
		err = server.DB.Debug().Model(&chat).Create(&chat).Association("Users").Clear().Error
		err = server.DB.Debug().Model(&chat).Association("Users").Append(&user,&user2).Error
		fmt.Println(chat.ID)
		if err != nil {
			formattedError := formaterror.FormatError(err.Error())
			errList = formattedError
			c.JSON(http.StatusNotFound, gin.H{
				"status": http.StatusNotFound,
				"error":  errList,
			})
			return
		}
	}else{
		chat.Users= []*models.User{
			&user, &user2,
		}
	}


	d := models.ELKMessage{
		Body : "",
		CreatedAt : time.Now(),
		SenderId: user.ID,
		SenderName: user.Lastname,
		Groupe: false,
		ConversationId :chat.ID,
		ConversationParticipantIds: []uint64{user.ID,user2.ID}}


	_, err = server.Elk.Index().
		Index(os.Getenv("ELK_MESSAGE_INDEX")).
		Type(os.Getenv("ELK_MESSAGE_TYPE")).
		BodyJson(&d).
		Do(server.Context)

	d2 := models.ELKNotification{
		Message : "L'utilisateur "+ user.Firstname+" a pris contact avec vous",
		Type: "Chat",
		TargetName: user2.Firstname+" "+user2.Lastname,
		CreatedAt : time.Now(),
		TargetID: 	user2.ID,
		TargetEmail: user2.Email,
	}
	
	server.sendNotification(d2)
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(chat)

	server.Melody.BroadcastFilter(reqBodyBytes.Bytes(), func(q *melody.Session) bool {
		_, exist := q.Get(user2.Email)
		return exist
	})	

	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": chat,
	})

}

func (server *Server) createGroupeChat(c *gin.Context){
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	requestBody := map[string]string{}
	_ = json.Unmarshal(body, &requestBody)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
		return
	}

	// check if the user is the trip owner
	trip := models.Trip{}
	tripId ,_ := strconv.ParseUint(requestBody["tripId"], 10, 64)
	err = server.DB.Debug().Model(models.Trip{}).Where("author_id = ? and id = ?", uid, tripId ).Take(&trip).Error
	fmt.Println(err)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized to create groupe"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	chat := models.Chat{}
	err = server.DB.Debug().Model(models.Chat{}).Where("create_by = ? and name = ?", uid,requestBody["tripName"] ).Take(&chat).Error
	fmt.Println(err)
	if err != nil {
		chat.CreateBy = uid
		chat.Name = requestBody["tripName"]
		err = server.DB.Debug().Model(models.Chat{}).Create(&chat).Error
		if err != nil {
			formattedError := formaterror.FormatError(err.Error())
			errList = formattedError
			c.JSON(http.StatusNotFound, gin.H{
				"status": http.StatusNotFound,
				"error":  errList,
			})
			return
		}
	}



	var userArray  []uint64
	a := strings.Split(requestBody["particiapantID"], ",")
	var creatby uint64
	for _, v := range a {
		creatby,_ = strconv.ParseUint( v, 10, 64)
		userArray = append(userArray, creatby)
	}

	d := models.ELKMessage{
		Body : "",
		CreatedAt : time.Now(),
		SenderId: uid,
		Groupe: true,
		SenderName: requestBody["name"],
		ConversationId :chat.ID,
		ConversationName : requestBody["tripName"],
		ConversationParticipantIds: userArray,
	}
	put1, err := server.Elk.Index().
		Index(os.Getenv("ELK_MESSAGE_INDEX")).
		Type(os.Getenv("ELK_MESSAGE_TYPE")).
		BodyJson(&d).
		Do(server.Context)
		
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Indexed message %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	valuen := models.ELKNotification{
		Message : "Une discussion pour le voyage "+requestBody["tripName"]+" a été créée",
		Type: "GroupChat",
		TargetName: requestBody["tripName"],
		CreatedAt : time.Now(),
		TargetID:  chat.ID,
		TargetEmail: requestBody["particiapantID"],
	}
	email := strings.Split(requestBody["emails"], ",") 	
	server.sendGrouPNotification(valuen, email)
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": put1,
	})
}

func  (server *Server) getMessage(c *gin.Context){
	message := models.Message{}

	fmt.Println("err")
	messages, err := message.GetAllMessage(server.DB)
	if err != nil {
		errList["No_trip"] = "No Trip Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": messages,
	})
}


func  (server *Server) wsMessage(c *gin.Context){	
	server.Melody.HandleRequest(c.Writer, c.Request)
}

func (server *Server) pushMsg(s *melody.Session, messagesent []byte) {

	msg_data  := map[string]string{}
	err := json.Unmarshal(messagesent, &msg_data)
	if err != nil {
	  	fmt.Println(err)
	}
	
	
	si,_ := strconv.ParseUint(msg_data["discussionId"], 10, 64)
	creatby, _ := strconv.ParseUint( msg_data["createBy"], 10, 64)
	targetid, _ := strconv.ParseUint( msg_data["targetID"], 10, 64)

	///SQL method
	message := models.Message{}
	message.Body = msg_data["content"]
	message.CreatedBy = creatby 
	chat := models.Chat{}
	err = server.DB.Debug().Model(models.Chat{}).Where("id = ?", msg_data["discussionId"]).Take(&chat).Error
	server.DB.Model(&chat).Association("Messages").Append(&message)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(msg_data)

	/// ELK METHOD
	d := models.ELKMessage{
		Body : msg_data["content"],
		CreatedAt : time.Now(),
		SenderId: creatby,
		Groupe: false,
		SenderName: msg_data["targetName"],
		ConversationId :si,
		ConversationParticipantIds: []uint64{creatby,targetid}}


	put1, err := server.Elk.Index().
		Index(os.Getenv("ELK_MESSAGE_INDEX")).
		Type(os.Getenv("ELK_MESSAGE_TYPE")).
		BodyJson(&d).
		Do(server.Context)
		
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Indexed message %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	valuen := models.ELKNotification{
		Message : msg_data["content"],
		Type: "Message",
		TargetName: msg_data["createBy"],
		CreatedAt : time.Now(),
		TargetID:  si,
		TargetEmail: msg_data["targetUsername"],
	}
	
	server.sendNotification(valuen)


	/*
	msg_data["type"] = "message"
	//newData, err := json.Marshal(msg_data)
	server.Melody.BroadcastFilter(newData, func(q *melody.Session) bool {
		_, exist := q.Get(msg_data["targetUsername"])
		return exist
	})	*/
}


func (server *Server) getELKMessages(c *gin.Context){
	ret := make(map[int]interface{})
	chat := make(map[int]map[int]interface{})
	bq := elastic.NewBoolQuery()
	limit := 10
	offset := 0
	userID := c.Param("id")
	user := models.User{}

	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		errList["No_user"] = "No User Found"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	topTagsHitsvAgg := elastic.NewTopHitsAggregation().Sort("created_at", false).Size(30).FetchSource(true)
	topTagsAggv :=  elastic.NewTermsAggregation().Field("conversation_id").SubAggregation("top_tag_hits", topTagsHitsvAgg)
	bq.Must(elastic.NewMatchQuery("conversation_participant_ids", userID))
	sr, err := server.Elk.Search().
			Index(os.Getenv("ELK_MESSAGE_INDEX")).
			Type(os.Getenv("ELK_MESSAGE_TYPE")). 
			Query(bq).
			SortBy(elastic.NewFieldSort("created_at").Desc().SortMode("max")).
			Aggregation("chat", topTagsAggv).
			From(offset). 
			Size(limit). 
			Pretty(true). 
			Do(server.Context)        
	for k, v := range sr.Hits.Hits { 
		ret[k] = v.Source
	}

	groups, _ := sr.Aggregations.Terms("chat")
	for k, v := range groups.Buckets { 
		chat[k] =  make(map[int]interface{})
		topHits, _ :=v.TopHits("top_tag_hits") 
		for k2, v2 := range topHits.Hits.Hits { 
			chat[k][k2] = v2.Source
		}
	}
	log.Println(ret)
	log.Println(err)

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": ret,
		"chat" : chat,
		"users" : users,
	})
}


func (server *Server) appendGroupeMessage(c *gin.Context) {
	_, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	requestBody := map[string]string{}
	_ = json.Unmarshal(body, &requestBody)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
		return
	}

	si ,_ := strconv.ParseUint(requestBody["discussionId"], 10, 64)
	creatby, _ := strconv.ParseUint( requestBody["createBy"], 10, 64)
	email := strings.Split(requestBody["targetUsername"], ",") 
	var userArray  []uint64
	a := strings.Split(requestBody["targetID"], ",")
	var convId uint64
	for _, v := range a {
		convId,_ = strconv.ParseUint( v, 10, 64)
		userArray = append(userArray, convId)
	}

	d := models.ELKMessage{
		Body : requestBody["content"],
		CreatedAt : time.Now(),
		SenderId: creatby,
		SenderName: requestBody["senderName"],
		ConversationName : requestBody["targetName"] ,
		ConversationId :si,
		Groupe: true,
		ConversationParticipantIds:userArray}

	put1, err := server.Elk.Index().
		Index(os.Getenv("ELK_MESSAGE_INDEX")).
		Type(os.Getenv("ELK_MESSAGE_TYPE")).
		BodyJson(&d).
		Do(server.Context)
		
	valuen := models.ELKNotification{
		Message : requestBody["content"],
		Type: "GroupMessage",
		TargetName: requestBody["senderName"],
		CreatedAt : time.Now(),
		TargetID:  si,
		TargetEmail: requestBody["targetUsername"],
	}
		
	server.sendGrouPNotification(valuen, email)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Indexed message %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": put1,
	})
}

func (server *Server) getELKConversation(c *gin.Context){
	_, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	requestBody := map[string]string{}
	_ = json.Unmarshal(body, &requestBody)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
		return
	}

	ret := make(map[int]interface{})
	bq := elastic.NewBoolQuery()
	limit := 40
	userID := requestBody["id"]
	offset ,_ := strconv.Atoi(requestBody["lastIndex"])

	bq.Must(elastic.NewMatchQuery("conversation_id", userID))
	sr, err := server.Elk.Search().
			Index(os.Getenv("ELK_MESSAGE_INDEX")).
			Type(os.Getenv("ELK_MESSAGE_TYPE")). 
			Query(bq).
			SortBy(elastic.NewFieldSort("created_at").Desc().SortMode("max")).
			From(offset). 
			Size(limit). 
			Pretty(true). 
			Do(server.Context)        
	for k, v := range sr.Hits.Hits { 
		ret[k] = v.Source
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": ret,
	})
}
