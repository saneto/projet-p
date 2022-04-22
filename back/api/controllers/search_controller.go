package controllers

import (
	"encoding/json"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
	"sync"
	"context"
	"bytes"
	"strconv"
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/models"
	"os"

	"github.com/olivere/elastic"
)
func (server *Server) insertTestValue(c *gin.Context){
	
	var (
		r  map[string]interface{}
		wg sync.WaitGroup
	  )
	cfg := elasticsearch.Config{
	
	  }
	es, err := elasticsearch.NewClient(cfg)
	//es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Println("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Println("Error getting response: %s", err)
	}

	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	  }
	  // Print client and server version numbers.
	  log.Printf("Client: %s", elasticsearch.Version)
	  log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	  log.Println(strings.Repeat("~", 37))
	
	  // 2. Index documents concurrently
	  //
	  	for i, title := range []string{"Test Odqs ddne", "Test Two"} {
			wg.Add(1)
		
			go func(i int, title string) {
				defer wg.Done()
			
				// Build the request body.
				var b strings.Builder

				b.WriteString(`	{ 
						"title" : "dsqd",
						"country" : "France",
						"description" : "paris",
						"nb_days" : "2",
						"nb_traveler" : "0",
						"budget" : "2",
						"start_date" : "2020-10-01",
						"end_date" : "2020-10-12",
						"type":" ` )
				b.WriteString(strconv.Itoa(i))
				b.WriteString(` ",	` )
				b.WriteString(`"analyzer": "analyzer-name" }`)

				// Set up the request object.
				req := esapi.IndexRequest{
					Index:      "testv2",
					DocumentID: strconv.Itoa(i + 1),
					Body:       strings.NewReader(b.String()),
					Refresh:    "true",
				}
				
				log.Println(b.String())
				// Perform the request with the client.
				res, err := req.Do(context.Background(), es)
				if err != nil {
					log.Fatalf("Error getting response: %s", err)
				}
				defer res.Body.Close()
			
				if res.IsError() {
					log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
				} else {
					// Deserialize the response into a map.
					var r map[string]interface{}
					if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
					} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
					}
				}
			}(i, title)
	 	 }
	  wg.Wait()

	  log.Println(strings.Repeat("-", 37))
	  var buf bytes.Buffer
	  query := map[string]interface{}{
		"query": map[string]interface{}{
		  "match": map[string]interface{}{
			"title": "test",
		  },
		},
	  }
	  log.Println(query)
	  if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	  }
	
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": r,
	})
}

func (server *Server) searchTrips(c *gin.Context){
	var (
		r  map[string]interface{}
	)
	cfg := elasticsearch.Config{}

	// var buf bytes.Buffer
	es, err := elasticsearch.NewClient(cfg)
	//es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Println("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Println("Error getting response: %s", err)
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
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		var buf bytes.Buffer
		res, err = es.Search(
			es.Search.WithContext(context.Background()),
			es.Search.WithIndex("trip"),
			es.Search.WithBody(&buf),
			es.Search.WithTrackTotalHits(true),
			es.Search.WithPretty(),
			)
	}else{
		i := 1
		s := `{
			"query": {
			"bool": {
			"must": [`
		for k, v := range requestBody { 
			log.Println(k) 
			log.Println(v)
			if i < len(requestBody){
				s+=` { "match": { "`+k+`" : "`+v+`" }} , `
			}else{
				s+=`{ "match": { "`+k+`" : "`+v+`" }}  `
			}
			i++
		}
		s += `] }}}`
		res, err = es.Search(
			es.Search.WithContext(context.Background()),
			es.Search.WithIndex("triptestv"),
			es.Search.WithBody(strings.NewReader(s)),
			es.Search.WithTrackTotalHits(true),
			es.Search.WithPretty(),
		) 
	}

	
	
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			//log.Fatalf("Error parsing the response body: %s", err)
		} else {
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": r["hits"],
	})
}

func (server *Server) insertTrips(c *gin.Context, t *models.Trip) error{
	
	var (
		r  map[string]interface{}
	  )
	cfg := elasticsearch.Config{
	
	  }
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Println("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Println("Error getting response: %s", err)
	}


	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	  }
	  // Print client and server version numbers.
	//  log.Printf("Client: %s", elasticsearch.Version)
	 // log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	 // log.Println(strings.Repeat("~", 37))
	

	// Build the request body.
	var b strings.Builder

	b.WriteString(`{`)
	b.WriteString(`	"title" : "`)
	b.WriteString(t.Title)
	b.WriteString(`",
					"country" : "`)
	b.WriteString(t.Country)
	b.WriteString(`",
					"description" : "`)
	b.WriteString(t.Description)
	b.WriteString(`",
					"nb_days" : "`)
	b.WriteString(strconv.Itoa(t.NbDays))
	b.WriteString(`",
					"nb_traveler" : "`)
	b.WriteString(strconv.Itoa(t.NbTraveler))
	b.WriteString(` ",
					"budget" : "`)
	b.WriteString(strconv.Itoa(t.Budget))
	b.WriteString(`",
					"start_date" : "`)
	b.WriteString(t.StartDate.Format("2006-01-02"))
	b.WriteString(`",
					"end_date" : "`)
	b.WriteString(t.EndDate.Format("2006-01-02"))
	b.WriteString(`",
					"photo" : "`)
	b.WriteString(t.Base64)
	b.WriteString(`", "lodging" : "`)
	b.WriteString(t.Lodging)
	b.WriteString(`"	}`)
	req := esapi.IndexRequest{
		Index:      "triptestv",
		DocumentID: fmt.Sprint(t.ID),
		Body:       strings.NewReader(b.String()),
		Refresh:    "true",
	}
	
	//log.Println(b.String())
	// Perform the request with the client.
	res2, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res2.Body.Close()

	if res2.IsError() {
		log.Printf("[%s] Error indexing document ID=%d", res.Status(), t.ID)
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
		} else {
		// Print the response status and indexed document version.
		log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
			
	

	return nil
}


func (server *Server) insertELKTrip(t *models.Trip) error{
	put1, err := server.Elk.Index().
		Index(os.Getenv("ELK_TRIPS_INDEX")).
		Type(os.Getenv("ELK_TRIPS_TYPE")).
		Id(fmt.Sprint(t.ID)).
		BodyJson(t).
		Do(server.Context)
	if err != nil {
		return err
	}
	fmt.Printf("Indexed trip %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	return nil
}



func (server *Server) updateELKTrip(t *models.Trip, u *models.User) error{

	update, err := server.Elk.Update().
				Index(os.Getenv("ELK_TRIPS_INDEX")).
				Type(os.Getenv("ELK_TRIPS_TYPE")).
				Id(fmt.Sprint(t.ID)).
				Script(elastic.NewScriptInline("ctx._source.users += param.users").Param("users","dsqd")).
				Do(server.Context)
	if err != nil {
		// Handle error
		panic(err)
		fmt.Println(err)
	}
	fmt.Printf("New version of trip %q is now %d\n", update.Id, update.Version)
	return nil
}


func (server *Server) searchELKTrip(c *gin.Context){
	bq := elastic.NewBoolQuery()
	limit := 100
	offset := 0

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
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		
	}
	for k, v := range requestBody { 
		if k=="end_date" {
			bq.Filter(elastic.NewRangeQuery(k).
			To(v))
		} else if k=="start_date" {
			bq.Filter(elastic.NewRangeQuery(k).
			From(v))
		}else if k =="offset" {
			offset,_ = strconv.Atoi(v)
		}else if k =="limit" {
			limit,_ = strconv.Atoi(v)
		}else{
			bq.Must(elastic.NewMatchQuery(k, v))
		}
	}

	sr, err := server.Elk.Search().
			Index(os.Getenv("ELK_TRIPS_INDEX")).
			Type(os.Getenv("ELK_TRIPS_TYPE")). // search in type
			Query(bq).
			From(offset). // Starting from this result
			Size(limit).  // Limit of responds
			Do(server.Context)         // execute
	log.Println(err)
	log.Println(bq)
	if err != nil {
		errList["Invalid_body"] = "Error in request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	ret := make(map[int]interface{})
	for k, v := range sr.Hits.Hits { 
		ret[k] = v.Source
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": ret ,
	})
}

