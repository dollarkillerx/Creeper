package creeper_sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/dollarkillerx/creeper/internal/request"
	"github.com/dollarkillerx/creeper/internal/response"
	"github.com/dollarkillerx/urllib"
)

type CreeperSdk struct {
	addr  string
	token string
}

func New(addr string, token string) *CreeperSdk {
	return &CreeperSdk{
		addr:  addr,
		token: token,
	}
}

func (c *CreeperSdk) Log(index string, message string) error {
	if strings.TrimSpace(message) == "" {
		return nil
	}
	code, data, err := urllib.Post(fmt.Sprintf("%s/api/v1/log", c.addr)).KeepAlives().SetHeader("token", c.token).SetJsonObject(request.LogRequest{
		Index:   index,
		Message: message,
	}).ByteOriginal()
	if err != nil {
		return err
	}

	if code != 200 {
		return errors.New(string(data))
	}

	return nil
}

type Index struct {
	Data []string `json:"data"`
}

func (c *CreeperSdk) Index() ([]string, error) {
	code, data, err := urllib.Get(fmt.Sprintf("%s/api/v1/index", c.addr)).SetHeader("token", c.token).KeepAlives().ByteOriginal()
	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, errors.New(string(data))
	}

	var idx Index
	err = json.Unmarshal(data, &idx)
	if err != nil {
		return nil, err
	}

	return idx.Data, nil
}

func (c *CreeperSdk) DelIndex(index string) error {
	code, data, err := urllib.Post(fmt.Sprintf("%s/api/v1/del_index", c.addr)).KeepAlives().SetHeader("token", c.token).
		SetJsonObject(request.DelIndexRequest{
			Index: index,
		}).ByteOriginal()
	if err != nil {
		return err
	}

	if code != 200 {
		return errors.New(string(data))
	}

	return nil
}

func (c *CreeperSdk) LogSlimming(index string, retentionDays int64) error {
	code, data, err := urllib.Post(fmt.Sprintf("%s/api/v1/log_slimming", c.addr)).KeepAlives().SetHeader("token", c.token).
		SetJsonObject(request.LogSlimmingRequest{
			Index:         index,
			RetentionDays: retentionDays,
		}).ByteOriginal()
	if err != nil {
		return err
	}

	if code != 200 {
		return errors.New(string(data))
	}

	return nil
}

type searchResp struct {
	Data response.LogRespModel `json:"data"`
}

func (c *CreeperSdk) Search(index string, keyWord string, offset int64, limit int64, startTime int64, endTime int64) (count int64, list []interface{}, err error) {
	code, data, err := urllib.Post(fmt.Sprintf("%s/api/v1/search", c.addr)).KeepAlives().SetHeader("token", c.token).
		SetJsonObject(request.SearchRequest{
			Index:     index,
			KeyWord:   keyWord,
			Offset:    offset,
			Limit:     limit,
			StartTime: startTime,
			EndTime:   endTime,
		}).ByteOriginal()

	if err != nil {
		return 0, nil, err
	}

	if code != 200 {
		return 0, nil, errors.New(string(data))
	}

	var resp searchResp
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return 0, nil, err
	}

	return resp.Data.Total, resp.Data.List, nil
}
