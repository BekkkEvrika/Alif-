package api

import "github.com/gin-gonic/gin"

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func InternalError(c *gin.Context, mess string) {
	e := Error{Code: "500", Message: mess}
	c.JSON(500, e)
}

func NotFound(c *gin.Context, mess string) {
	e := Error{Code: "404", Message: mess}
	c.JSON(404, e)
}

func Success(c *gin.Context, mess string) {
	e := Error{Code: "200", Message: mess}
	c.JSON(200, e)
}

func BadRequest(c *gin.Context, mess string) {
	e := Error{Code: "400", Message: mess}
	c.JSON(400, e)
}
