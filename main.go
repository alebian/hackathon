package main

import (
	faktory "github.com/contribsys/faktory/client"
	ourWorker "hackathon/worker"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Job struct {
	Lambda string `json:"lambda"`
	Args   string `json:"args"`
}

func main() {
	work, ok := os.LookupEnv("WORKER")
	if ok {
		if work == "TRUE" {
			ourWorker.StartFaktory()

		} else {
			r := gin.Default()

			r.POST("/job", func(c *gin.Context) {
				var err error

				var json Job
				if err := c.BindJSON(&json); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				client, err := faktory.Open()
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				job := faktory.NewJob("LambdaEnqueuer", json.Lambda, json.Args)
				err = client.Push(job)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				c.JSON(201, gin.H{
					"message": "ok",
				})
			})

			r.GET("/ping", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "pong",
				})
			})
			r.Run() // listen and serve on 0.0.0.0:8080
		}
	}
}
