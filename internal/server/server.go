package server

import (
	"fmt"
	"log"
	"time"

	"github.com/dollarkillerx/async_utils"
	"github.com/dollarkillerx/creeper/internal/conf"
	"github.com/dollarkillerx/creeper/internal/models"
	"github.com/meilisearch/meilisearch-go"
	"github.com/rs/xid"
)

type Server struct {
	meilsearchClient *meilisearch.Client
	logChannel       chan models.Message
}

func New(meilsearchClient *meilisearch.Client) *Server {
	s := &Server{
		meilsearchClient: meilsearchClient,
		logChannel:       make(chan models.Message, 1000),
	}

	go s.core()

	return s
}

// core 核心服务
func (s *Server) core() {
	buf := make([]models.Message, conf.CONFIG.FlashSize)
	timeAfter := time.NewTicker(time.Second * time.Duration(conf.CONFIG.FlashSec))

	for {
		select {
		case <-timeAfter.C:
			if len(buf) != 0 {
				err := s.logs(buf)
				if err != nil {
					log.Println(err)
				}

				buf = make([]models.Message, conf.CONFIG.FlashSize)
			}
		case data, ex := <-s.logChannel:
			if !ex {
				return
			}

			buf = append(buf, data)
			if len(buf) >= conf.CONFIG.FlashSize {
				err := s.logs(buf)
				if err != nil {
					log.Println(err)
				}

				buf = make([]models.Message, conf.CONFIG.FlashSize)
			}
		}
	}
}

func (s *Server) logs(msg []models.Message) error {
	insMap := map[string][]models.Message{}
	for i := range msg {
		idx := i
		insMap[msg[idx].Index] = append(insMap[msg[idx].Index], msg[idx])
	}

	over := make(chan struct{})

	poolFunc := async_utils.NewPoolFunc(conf.CONFIG.MaxFlashPoolSize, func() {
		close(over)
	})

	for i := range insMap {
		k := i
		insData := insMap[k]
		poolFunc.Send(func() {
			_, err := s.meilsearchClient.Index(k).AddDocuments(insData)
			if err != nil {
				log.Printf("Insert :%s  Error: %s  Error Count: %d \n", k, err.Error(), len(insData))
			}
		})
	}

	poolFunc.Over()
	<-over

	return nil
}

// Log 记录日志
func (s *Server) Log(msg models.Message) {
	msg.ID = xid.New().String()

	go func() {
		s.logChannel <- msg
	}()
}

func (s *Server) DelIndex(idx string) error {
	_, err := s.meilsearchClient.DeleteIndex(idx)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// LogSlimming 日志瘦身
func (s *Server) LogSlimming(idx string, retentionDays int64) error {
	// 获取 保留天数内的日志

	search, err := s.meilsearchClient.Index(idx).Search("", &meilisearch.SearchRequest{
		Filter: fmt.Sprintf("create_at > %d", time.Now().Add(-24*time.Hour*time.Duration(retentionDays)).Unix()),
	})
	if err != nil {
		log.Println("LogSlimming: ", err)
		return err
	}

	// 删除 idx
	s.meilsearchClient.DeleteIndex(idx)
	// 插入数据
	if len(search.Hits) != 0 {
		_, err := s.meilsearchClient.Index(idx).AddDocuments(search.Hits)
		if err != nil {
			log.Println("LogSlimming AddDocuments: ", err)
			return err
		}
	}

	return nil
}

// SearchLog 查询日志
func (s *Server) SearchLog(keyWord string, idx string, offset int64, limit int64, startTime int64, endTime int64) (int64, []interface{}, error) {
	if limit <= 0 {
		limit = 500
	}
	if offset < 0 {
		offset = 0
	}

	var filter interface{}
	if startTime != 0 && endTime != 0 {
		filter = fmt.Sprintf("create_at => %d AND create_at <= %d", startTime, endTime)
	}

	if startTime != 0 {
		filter = fmt.Sprintf("create_at => %d ", startTime)
	}

	if endTime != 0 {
		filter = fmt.Sprintf("create_at <= %d ", endTime)
	}

	searchRes, err := s.meilsearchClient.Index(idx).Search(keyWord,
		&meilisearch.SearchRequest{
			Filter: filter,
			Limit:  limit,
			Offset: offset,
		},
	)
	if err != nil {
		return 0, nil, err
	}

	return searchRes.NbHits, searchRes.Hits, nil
}
