package controller

import (
	"crawler/pipeline"
	"log"
)

type Controller struct {
	schedule Schedule
	workerCount int
}

type Schedule interface {
	Submit(request Request)
	Init(chan Request) // init schedule with input channel
}

func (controller *Controller) Run(seeds ...Request) {
	/* init channel to communicate with schedule and worker */
	input := make(chan Request)
	output := make(chan Result)

	/* add input channel to schedule */
	controller.schedule.Init(input)

	/* create workers */
	for i := 0; i < controller.workerCount; i++ {
		initWorker(input, output)
	}

	/* submit workers we created to schedule */
	for _, request := range seeds {
		controller.schedule.submit(request)
	}

	for {
		result := <- output
		for _, content := range result.Content {
			pipeline.Process(content)
		}

		for _, request := range result.Requests {
			controller.schedule.submit(request)
		}
	}
}

func initWorker(input chan Request, output chan Result) {
	go func() {
		request := <- input
		result, err := worker(request)

		if err != nil {
			log.Printf("Error: %s", err)
		}

		output <- result
	}()
}
