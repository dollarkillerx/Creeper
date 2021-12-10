package api

import (
	"github.com/dollarkillerx/creeper/internal/conf"
	"github.com/dollarkillerx/creeper/internal/models"
	"github.com/dollarkillerx/creeper/internal/request"
	"github.com/dollarkillerx/creeper/internal/response"
	"github.com/gin-gonic/gin"

	"time"
)

func (a *ApiServer) router() {
	api := a.app.Group("/api")
	v1 := api.Group("/v1").Use(authToken)
	{
		v1.GET("/index", a.allIndex)
		v1.POST("/del_index", a.delIndex)
		v1.POST("/log_slimming", a.logSlimming)
		v1.POST("/log", a.log)
		v1.POST("/search", a.search)
		v1.POST("/web_search", a.webSearch)
	}

	if conf.CONFIG.Token != "" {
		a.app.GET("/", gin.BasicAuth(gin.Accounts{
			"token": conf.CONFIG.Token,
		}), a.webUi)
	} else {
		a.app.GET("/", a.webUi)
	}
}

func (a *ApiServer) allIndex(ctx *gin.Context) {
	ctx.JSON(200, response.UniversalReturn{
		Data: a.ser.AllIndex(),
	})
}

func (a *ApiServer) delIndex(ctx *gin.Context) {
	var req request.DelIndexRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, response.UniversalReturn{
			Code:    -1,
			Message: "400 参数错误",
		})
		return
	}

	if req.Index == "" {
		ctx.JSON(400, response.UniversalReturn{
			Code:    -1,
			Message: "400 参数错误",
		})
		return
	}

	err = a.ser.DelIndex(req.Index)
	if err != nil {
		ctx.JSON(500, response.UniversalReturn{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(200, response.UniversalReturn{})
}

// logSlimming 日志瘦身
func (a *ApiServer) logSlimming(ctx *gin.Context) {
	var req request.LogSlimmingRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, response.UniversalReturn{
			Code:    -1,
			Message: "400 参数错误",
		})
		return
	}

	if req.Index == "" || req.RetentionDays == 0 {
		ctx.JSON(400, response.UniversalReturn{
			Code:    -1,
			Message: "400 参数错误",
		})
		return
	}

	err = a.ser.LogSlimming(req.Index, req.RetentionDays)
	if err != nil {
		ctx.JSON(500, response.UniversalReturn{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(200, response.UniversalReturn{})
}

func (a *ApiServer) log(ctx *gin.Context) {
	var req request.LogRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, response.UniversalReturn{
			Code:    -1,
			Message: "400 参数错误",
		})
		return
	}

	if req.Index == "" || req.Message == "" {
		ctx.JSON(400, response.UniversalReturn{
			Code:    -1,
			Message: "400 参数错误",
		})
		return
	}

	a.ser.Log(models.Message{
		Index:   req.Index,
		Message: req.Message,
	})

	ctx.JSON(200, response.UniversalReturn{})
}

func (a *ApiServer) search(ctx *gin.Context) {
	var req request.SearchRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, response.UniversalReturn{
			Code:    -1,
			Message: "400 参数错误",
		})
		return
	}

	if req.Index == "" {
		ctx.JSON(400, response.UniversalReturn{
			Code:    -1,
			Message: "400 参数错误",
		})
		return
	}

	total, data, err := a.ser.SearchLog(req.KeyWord, req.Index, req.Offset, req.Limit, req.StartTime, req.EndTime)
	if err != nil {
		ctx.JSON(500, response.UniversalReturn{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, response.UniversalReturn{
		Data: response.LogRespModel{
			Total: total,
			List:  data,
		},
	})
}

var timeTemplate1 = "20060102"

func (a *ApiServer) webSearch(ctx *gin.Context) {
	var req request.SearchRequestV2
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, response.UniversalReturn{
			Code:    -1,
			Message: "400 参数错误",
		})
		return
	}

	if req.Index == "" {
		ctx.JSON(400, response.UniversalReturn{
			Code:    -1,
			Message: "400 参数错误",
		})
		return
	}

	var startTimeInt int64
	var endTimeInt int64

	if req.StartTime != "" {
		startTime, err := time.ParseInLocation(timeTemplate1, req.StartTime, time.Local)
		if err != nil {
			ctx.JSON(400, response.UniversalReturn{
				Code:    -1,
				Message: "400 startTime 参数错误",
			})
			return
		}

		startTimeInt = startTime.Unix()
	}

	if req.EndTime != "" {
		endTime, err := time.ParseInLocation(timeTemplate1, req.EndTime, time.Local)
		if err != nil {
			ctx.JSON(400, response.UniversalReturn{
				Code:    -1,
				Message: "400 endTime 参数错误",
			})
			return
		}
		endTimeInt = endTime.Unix()
	}

	total, data, err := a.ser.SearchLog(req.KeyWord, req.Index, req.Offset, req.Limit, startTimeInt, endTimeInt)
	if err != nil {
		ctx.JSON(500, response.UniversalReturn{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, response.UniversalReturn{
		Data: response.LogRespModel{
			Total: total,
			List:  data,
		},
	})
}
