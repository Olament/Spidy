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

	crawler.Run(
		controller.Request{
		Url:       "www.google.com",
		ParseFunc: parser.Parse,
	},
		controller.Request{
		Url:       "www.bing.com",
		ParseFunc: parser.Parse,
	})
}