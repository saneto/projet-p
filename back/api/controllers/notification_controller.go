package controllers

import (
	"bytes"
	"github.com/olahol/melody"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
	"log"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/models"
)



func (server *Server) sendNotification(notification  models.ELKNotification){
	reqBodyBytes := new(bytes.Buffer)
	if(notification.Type!="Message"){	
		server.insertELKNotification(notification)
	}	
	json.NewEncoder(reqBodyBytes).Encode(notification)
	server.Melody.BroadcastFilter(reqBodyBytes.Bytes(), func(q *melody.Session) bool {
		_, exist := q.Get(notification.TargetEmail)
		return exist
	})	
}


func (server *Server) sendGrouPNotification(notification  models.ELKNotification, emails []string ){
	reqBodyBytes := new(bytes.Buffer)
	if(notification.Type!="Message" && notification.Type!="GroupMessage" ){	
		server.insertELKNotification(notification)
	}	
	json.NewEncoder(reqBodyBytes).Encode(notification)
	server.Melody.BroadcastFilter(reqBodyBytes.Bytes(), func(q *melody.Session) bool {
		for _,v := range emails {
			_, exist := q.Get(v)
			if exist == true{
				return exist
			}
		}
		return false
	})	
}


func (server *Server) newConnection(s *melody.Session){	
	var c interface {}
	c =s.Request.URL.Query()["apiKey"][0]
	s.Set(s.Request.URL.Query()["apiKey"][0], c)

}

func  (server *Server) wsNotification(c *gin.Context){	
	server.Melody.HandleRequest(c.Writer, c.Request)
}

func (server *Server) insertELKNotification(notification models.ELKNotification) error{
	put1, err := server.Elk.Index().
		Index(os.Getenv("ELK_NOTIF_INDEX")).
		Type(os.Getenv("ELK_NOTIF_TYPE")).
		BodyJson(notification).
		Do(server.Context)
	if err != nil {
		return err
	}
	fmt.Printf("Indexed notif %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	return nil
}



func (server *Server) getELKNotification(c *gin.Context){
	ret := make(map[int]interface{})
	bq := elastic.NewBoolQuery()
	limit := 10
	offset := 0
	userID := c.Param("id")
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		/*errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return*/
	}
	fmt.Print(body)

	//bq.Sort(, false)
	bq.Must(elastic.NewMatchQuery("target_id", userID))
	sr, err := server.Elk.Search().
			Index(os.Getenv("ELK_NOTIF_INDEX")).
			Type(os.Getenv("ELK_NOTIF_TYPE")). // search in type
			Query(bq).
			SortBy(elastic.NewFieldSort("created_at").Desc().SortMode("max")).
			From(offset). // Starting from this result
			Size(limit).  // Limit of responds
			Pretty(true). 
			Do(server.Context)         // execute
	log.Println(err)
	log.Println(sr)
	if err != nil {
		errList["Invalid_body"] = "Error in request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	for k, v := range sr.Hits.Hits { 
		ret[k] = v.Source
	}


	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": ret ,
	})
}