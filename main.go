package main

import (
	"fmt"
	faktory "github.com/contribsys/faktory/client"
	ourWorker "hackathon/worker"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	work, ok := os.LookupEnv("WORKER")
	if ok {
		if work == "TRUE" {
			ourWorker.StartFaktory()

		} else {
			r := gin.Default()
			r.GET("/ping", func(c *gin.Context) {

				var err error
				client, err := faktory.Open()
				if err != nil {
					fmt.Println("Failed to open faktory client")
				}
				job := faktory.NewJob("SomeJob", 1, 2, 3)
				err = client.Push(job)
				if err != nil {
					fmt.Println("Failed to push job to faktory")
				}

				c.JSON(200, gin.H{
					"message": "pong",
				})
			})
			r.Run() // listen and serve on 0.0.0.0:8080
		}
	}
}
