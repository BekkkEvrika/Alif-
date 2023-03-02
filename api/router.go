package api

import (
	"AlifTask/log"
	"AlifTask/models"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	key = "test"
)

func MainRouter() http.Handler {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.Use(log.LogMiddleware)
	v1 := r.Group("/api/v1")
	v1.Use(AuthMiddleware)
	v1.POST("/account/exists", AccountExistsHandler)
	v1.POST("/replenishment", ReplenishmentHandler)
	v1.POST("/report", ReportHandler)
	v1.POST("account/balance", GetBalanceHandler)
	return r
}

func AuthMiddleware(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Request.Header.Get("X-UserId"))
	digest := c.Request.Header.Get("X-Digest")
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	if ShaHMAC(string(bodyBytes), key) == digest {
		cl := models.Client{ID: userId}
		if !cl.Exists() {
			c.AbortWithError(401, fmt.Errorf("Not Authorized! "))
			return
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		c.Next()
	} else {
		log.WriteLn(ShaHMAC(string(bodyBytes), key))
		log.WriteLn(string(bodyBytes))
		c.AbortWithError(401, fmt.Errorf("Not Authorized! "))
	}
}
