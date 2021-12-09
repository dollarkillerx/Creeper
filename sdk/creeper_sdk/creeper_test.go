package creeper_sdk

import (
	"fmt"
	"log"
	"testing"
)

func TestCreeperIndex(t *testing.T) {
	sdk := New("http://127.0.0.1:8745", "my_token")
	index, err := sdk.Index()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(index)
}

func TestCreeperLog(t *testing.T) {
	sdk := New("http://127.0.0.1:8745", "my_token")
	for i := 0; i < 2000; i++ {
		sdk.Log("creeper", fmt.Sprintf("hello world %d", i))
	}
}

func TestCreeperShow(t *testing.T) {
	sdk := New("http://127.0.0.1:8745", "my_token")
	total, resp, err := sdk.Search("creeper", "12", 0, 0, 0, 0)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("total: ", total)
	log.Println("resp: ", resp)
}
