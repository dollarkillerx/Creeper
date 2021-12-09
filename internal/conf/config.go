package conf

import (
	"log"
	"os"
	"strconv"
)

type creeperConfig struct {
	ListenAddr string `json:"listen_addr"`
	Token      string `json:"token"`

	MeilisearchAddr  string `json:"meilisearch_addr"`
	MeilisearchToken string `json:"meilisearch_token"`

	FlashSec         int64 `json:"flash_sec"`           // 刷新时间 sec
	FlashSize        int   `json:"flash_size"`          // 刷新大小
	MaxFlashPoolSize int   `json:"max_flash_pool_size"` // 最大刷新线程数
}

var CONFIG = &creeperConfig{}

func InitConfig() {
	listenAddr := os.Getenv("ListenAddr")
	if listenAddr != "" {
		CONFIG.ListenAddr = listenAddr
	} else {
		CONFIG.ListenAddr = "0.0.0.0:8745"
	}

	token := os.Getenv("Token")
	if token != "" {
		CONFIG.Token = token
	}

	meilisearchAddr := os.Getenv("MeilisearchAddr")
	if meilisearchAddr != "" {
		CONFIG.MeilisearchAddr = meilisearchAddr
	} else {
		log.Fatalln("MeilisearchAddr is null")
	}

	meilisearchToken := os.Getenv("MeilisearchToken")
	if meilisearchToken != "" {
		CONFIG.MeilisearchToken = meilisearchToken
	} else {
		log.Fatalln("MeilisearchToken is null")
	}

	flashSec := os.Getenv("FlashSec")
	if flashSec != "" {
		atoi, err := strconv.Atoi(flashSec)
		if err == nil {
			CONFIG.FlashSec = int64(atoi)
		}
	}

	flashSize := os.Getenv("FlashSize")
	if flashSize != "" {
		atoi, err := strconv.Atoi(flashSize)
		if err == nil {
			CONFIG.FlashSize = atoi
		}
	}

	maxFlashPoolSize := os.Getenv("MaxFlashPoolSize")
	if maxFlashPoolSize != "" {
		atoi, err := strconv.Atoi(maxFlashPoolSize)
		if err == nil {
			CONFIG.MaxFlashPoolSize = atoi
		}
	}

	if CONFIG.FlashSec < 1 {
		CONFIG.FlashSec = 3
	}

	if CONFIG.FlashSize < 10 {
		CONFIG.FlashSize = 1000
	}

	if CONFIG.MaxFlashPoolSize < 10 {
		CONFIG.MaxFlashPoolSize = 100
	}

	return
}
