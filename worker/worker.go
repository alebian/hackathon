package worker

import (
	"fmt"
	worker "github.com/contribsys/faktory_worker_go"
	"io/ioutil"
	"net/http"
)

func StartFaktory() {
	mgr := worker.NewManager()

	// register job types and the function to execute them
	mgr.Register("LambdaEnqueuer", EnqueueLambda)
	//mgr.Register("AnotherJob", anotherFunc)

	// use up to N goroutines to execute jobs
	mgr.Concurrency = 20

	// pull jobs from these queues, in this order of precedence
	mgr.Queues = []string{"critical", "default", "bulk"}

	// Start processing jobs, this method does not return
	mgr.Run()

	fmt.Println("Started worker")
}

func EnqueueLambda(ctx worker.Context, args ...interface{}) error {
	fmt.Println("Working on job", ctx.Jid())

	lambdaName := args[0].(string)
	name := args[1].(string)

	url := fmt.Sprintf("https://6kr1f1mp3h.execute-api.us-east-1.amazonaws.com/api2/%s?name=%s", lambdaName, name)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed request:", err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	fmt.Println("Finished working on job", ctx.Jid())
	return nil
}
