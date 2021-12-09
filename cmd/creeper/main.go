package main

import (
	"github.com/dollarkillerx/creeper/internal/api"
	"github.com/dollarkillerx/creeper/internal/conf"
	"github.com/dollarkillerx/creeper/internal/server"
	"github.com/meilisearch/meilisearch-go"

	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	conf.InitConfig()

	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   conf.CONFIG.MeilisearchAddr,
		APIKey: conf.CONFIG.MeilisearchToken,
	})

	s, err := server.New(client)
	if err != nil {
		log.Fatalln(err)
	}

	ser := api.New(s)
	if ser.Run(); err != nil {
		log.Fatalln(err)
	}
}
