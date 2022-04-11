package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context)  {
	data := make(map[string]interface{})
	data["key"] = "value"
	c.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"msg" : "ok",
		"data" : data,
	})
}

func Push(c *gin.Context) {
	// 验证job_id唯一性

	// hash job_id 进bucket

	c.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"msg" : "ok",
		"data" : "push",
	})
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
