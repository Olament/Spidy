package crawler

import (
	"crawler/controller"
	"crawler/parser"
	"crawler/schedule"
)

func main() {
	crawler := controller.Controller{
		&schedule.SimpleSchedule{},
		10,
	}

	crawler.Run(controller.Request{
		"www.google.com",
		parser.Parse,
	})
}