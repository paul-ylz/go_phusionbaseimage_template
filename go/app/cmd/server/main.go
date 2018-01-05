package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
)

func main() {
	r := gin.Default()
	r.Use(requestLogger())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "go / phusion/baseimage template",
		})
	})
	r.Run()
}

func requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Fatal(err)
		}
		// Have to save twice because after it is read
		// the buffer won't be accessible to the handler
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
		log.Printf(readBody(rdr1))
		c.Request.Body = rdr2
		c.Next()
	}
}

func readBody(reader io.Reader) string {
	var buf bytes.Buffer
	buf.ReadFrom(reader)
	return buf.String()
}
