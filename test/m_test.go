package test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/dollarkillerx/creeper/internal/models"
	"github.com/meilisearch/meilisearch-go"
	"github.com/rs/xid"
)

func TestMeilisearchInsert3(t *testing.T) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://127.0.0.1:7700",
		APIKey: "root",
	})

	client.DeleteIndex("movies3")

	index := client.Index("movies3")

	var p []models.Message
	for i := 0; i < 100; i++ {
		time.Sleep(time.Millisecond * 20)
		p = append(p, models.Message{
			ID:       xid.New().String(),
			Message:  fmt.Sprintf("hello: %d", i),
			CreateAt: time.Now().UnixNano(),
		})
	}

	update, err := index.AddDocuments(p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(update.UpdateID)
}

func TestMeilisearchLogSlimming(t *testing.T) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://127.0.0.1:7700",
		APIKey: "root",
	})

	index := client.Index("movies3")

	_, err := index.UpdateFilterableAttributes(&[]string{"create_at"})
	if err != nil {
		log.Fatalln(err)
	}

	var fil interface{}
	fil = "create_at > 1639039544267782960 AND create_at < 1639039544510021600"

	searchRes, err := index.Search("",
		&meilisearch.SearchRequest{
			Filter: fil,
			Limit:  0,
			Offset: 1,
		},
	)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(searchRes.Hits))

	fmt.Println(searchRes.NbHits)

	marshal, err := json.Marshal(searchRes.Hits)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(marshal))
}

// Meilisearch 必须要有ID
func TestMeilisearchInsert2(t *testing.T) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://127.0.0.1:7700",
		APIKey: "root",
	})

	index := client.Index("movies2")

	p := []map[string]interface{}{
		{
			"id":        xid.New().String(),
			"message":   "helloworld",
			"create_at": time.Now().Unix(),
		},
		{
			"id":        xid.New().String(),
			"message":   "helloworld2",
			"create_at": time.Now().Unix(),
		},
		{
			"id":        xid.New().String(),
			"message":   "helloworld3",
			"create_at": time.Now().Unix(),
		},
	}

	update, err := index.AddDocuments(p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(update.UpdateID)
}

func TestMeilisearchDel(t *testing.T) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://127.0.0.1:7700",
		APIKey: "root",
	})

	client.Index("movies2").DeleteAllDocuments()
}

func TestMeilisearchInsert(t *testing.T) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://127.0.0.1:7700",
		APIKey: "root",
	})

	index := client.Index("movies")

	documents := []map[string]interface{}{
		{"id": 122133, "title": "Carol", "genres": []string{"Romance", "Drama"}},
		{"id": 11232323, "title": "Wonder Woman", "genres": []string{"Action", "Adventure"}},
		{"id": 431235423213, "title": "Life of Pi", "genres": []string{"Adventure", "Drama"}},
		{"id": 56123213, "title": "Mad Max: Fury Road", "genres": []string{"Adventure", "Science Fiction"}},
		{"id": 67123213, "title": "Moana", "genres": []string{"Fantasy", "Action"}},
		{"id": 67862131232137, "title": "Philadelphia", "genres": []string{"Drama"}},
	}
	update, err := index.AddDocuments(documents)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(update.UpdateID)
}

func TestMeilisearchSearch(t *testing.T) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://127.0.0.1:7700",
		APIKey: "root",
	})

	//updateId, err := client.Index("movies").UpdateFilterableAttributes(&[]string{"id", "genres"})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(updateId)

	//rankingRules := []string{
	//	"id",
	//}
	//client.Index("movies").UpdateRankingRules(&rankingRules)
	//client.Index("movies").UpdateSortableAttributes(&rankingRules)

	searchRes, err := client.Index("movies").Search("",
		&meilisearch.SearchRequest{
			//AttributesToHighlight: []string{"*"},
			Filter: "id > 1",
			//Limit: 10,
			Sort: []string{"id:desc"},
		})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(searchRes.Hits)
	fmt.Println(len(searchRes.Hits))
}

func TestMeilisearchSearch2(t *testing.T) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://127.0.0.1:7700",
		APIKey: "root",
	})

	//updateId, err := client.Index("movies").UpdateFilterableAttributes(&[]string{"id", "genres"})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(updateId)

	//rankingRules := []string{
	//	"id",
	//}
	//client.Index("movies").UpdateRankingRules(&rankingRules)
	//client.Index("movies").UpdateSortableAttributes(&rankingRules)

	searchRes, err := client.Index("movies2").Search("",
		&meilisearch.SearchRequest{
			//AttributesToHighlight: []string{"*"},
			//Filter: "id > 1",
			Limit: 10,
			//Sort: []string{"id:desc"},
		},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(searchRes.Hits)
	fmt.Println(len(searchRes.Hits))
}

func TestTime(t *testing.T) {
	timer := time.NewTimer(time.Second)

	for {
		select {
		case <-timer.C:
			fmt.Println("aaa")
		}
	}
}

func TestTime2(t *testing.T) {
	timer := time.NewTicker(time.Second)

	for {
		select {
		case <-timer.C:
			fmt.Println("aaa")
		}
	}
}
