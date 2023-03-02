package log

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Logger struct {
	gin.ResponseWriter
	ReqID   string
	Request *http.Request
	Body    string
}

func (w Logger) Write(b []byte) (int, error) {
	WriteLn("# " + w.ReqID + " URL: " + w.Request.URL.String())
	WriteLn("# " + w.ReqID + " METHOD: " + w.Request.Method)
	WriteLn("# " + w.ReqID + " BODY: " + w.Body)
	WriteLn("# " + w.ReqID + " RESPONSE: " + string(b))
	return w.ResponseWriter.Write(b)
}

var reqId int

func LogMiddleware(c *gin.Context) {
	reqId++
	WriteLn("REQUEST #" + strconv.Itoa(reqId))
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	blw := &Logger{ResponseWriter: c.Writer, ReqID: strconv.Itoa(reqId), Request: c.Request, Body: string(bodyBytes)}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	c.Writer = blw
	c.Next()
}
