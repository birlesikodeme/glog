package glog

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Ginglog(debug bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// some evil middlewares modify this values
		path := c.Request.URL.Path

		var req string
		if debug {
			req = strings.TrimSpace(getRequestBody(c))
		}

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		status := c.Writer.Status()
		method := c.Request.Method
		ip := c.ClientIP()

		if len(c.Errors) > 0 {
			Errorf("[%d] %s %s %s from: %s took: %s error: %s", status, method, path, req, ip, latency, c.Errors.String())
		} else {
			Infof("[%d] %s %s %s from: %s took: %s", status, method, path, req, ip, latency)
		}

	}
}

func getRequestBody(c *gin.Context) string {
	buf := bytes.NewBuffer(nil)
	var save io.ReadCloser
	save, c.Request.Body, _ = drainBody(c.Request.Body)
	io.Copy(buf, c.Request.Body)
	c.Request.Body = save
	return buf.String()
}

func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}
