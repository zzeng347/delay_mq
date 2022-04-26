package http

import (
	"delay_mq_v2/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"time"
)

func Hello(c *gin.Context)  {
	data := `{"key": "value"}`
	c.JSON(http.StatusOK, JSON{
		Code: ErrorCode,
		Message : "error test",
		Data: data,
	})

	return
}

func JsonRaw(c *gin.Context)  {
	data := `{"key": "value"}`
	Success(c, data)
	return
}

func Map(c *gin.Context)  {
	data := map[string]string{"key": "value"}
	Success(c, data)
	return
}

func Biz(c *gin.Context)  {
	var (
		err error
		job *model.PushJobReq
	)
	// 模拟延迟
	time.Sleep(3 * time.Second)

	if err = c.ShouldBindWith(&job, binding.JSON); err != nil {
		Fail(c, err.Error())
		return
	}
	//Fail(c, "error test")
	//return
	Success(c, job)
	return
}

func Push(c *gin.Context) {
	var (
		err error
		job *model.PushJobReq
	)

	if err = c.ShouldBindWith(&job, binding.JSON); err != nil {
		Fail(c, err.Error())
		return
	}

	if err = hSrv.PushJob(job); err != nil {
		Fail(c, err.Error())
		return
	}
	Success(c, job)
	return
}

func Pop(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"msg" : "ok",
		"data" : "pop",
	})
}

func Finish(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"msg" : "ok",
		"data" : "finish",
	})
}

func Delete(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"msg" : "ok",
		"data" : "delete",
	})
}
