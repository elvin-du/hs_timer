package main

import (
	"github.com/gin-gonic/gin"
	//	"net/http"
	"strconv"
	"time"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.LoadHTMLFiles("index.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"t": ""})
	})
	router.GET("/convert", Convert)

	router.Run(":10002")
}

func Convert(ctx *gin.Context) {
	tStr := ctx.Request.FormValue("t")
	t, err := strconv.ParseInt(tStr, 10, 64)
	if nil != err {
		ctx.Error(err,err.Error())
	}

	ctx.HTML(200, "index.html", gin.H{"t": I2Time(t)})
}

func I2Time(v int64) time.Time {
	return time.Unix(v/int64(time.Millisecond), int64(0))
}
