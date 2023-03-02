package api

import (
	"AlifTask/log"
	"AlifTask/models"
	"github.com/gin-gonic/gin"
)

func AccountExistsHandler(c *gin.Context) {
	cl := models.Client{}
	if err := c.ShouldBind(&cl); err != nil {
		log.WriteLn(err.Error())
		BadRequest(c, "Bad Request")
		return
	}
	if !cl.Exists() {
		NotFound(c, "Not Found!")
		return
	}
	Success(c, cl.Name)
}

func ReplenishmentHandler(c *gin.Context) {
	rep := models.Replenishment{}
	if err := c.ShouldBind(&rep); err != nil {
		log.WriteLn(err.Error())
		BadRequest(c, "Bad request")
		return
	}
	if err := rep.Create(); err != nil {
		log.WriteLn(err.Error())
		InternalError(c, "Internal error")
		return
	}
	Success(c, "Success")
}

func ReportHandler(c *gin.Context) {
	cl := models.Client{}
	if err := c.ShouldBind(&cl); err != nil {
		log.WriteLn(err.Error())
		BadRequest(c, "Bad request")
		return
	}
	if rp, err := models.GetCountSum(cl.Account); err != nil {
		log.WriteLn(err.Error())
		InternalError(c, "Internal error")
		return
	} else {
		c.JSON(200, rp)
	}
}

func GetBalanceHandler(c *gin.Context) {
	cl := models.Client{}
	if err := c.ShouldBind(&cl); err != nil {
		log.WriteLn(err.Error())
		BadRequest(c, "Bad request")
		return
	}
	if bal, err := cl.GetBalance(); err != nil {
		log.WriteLn(err.Error())
		NotFound(c, "Not found")
		return
	} else {
		c.JSON(200, gin.H{"balance": bal})
	}
}
