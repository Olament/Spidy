# Spidy
A simple concurrent golang web crawler

Notice that this crawler is work-in-progress. Many core features may subject to change in the the future. Please use with caution.

## Example
From `main.go`
```golang
crawler := controller.Controller {
	&schedule.SimpleSchedule{}, // select schedule 
	10,                         // number of workers
}

crawler.Run(
	controller.Request {
	Url:       "www.google.com", // URL
	ParseFunc: parser.Parse,     // Parser to parse the content from URL
},
	controller.Request{ .        // support multiple seed request
	Url:       "www.bing.com",
	ParseFunc: parser.Parse,
})
```

## Features to implement
- [ ] More choice of UA
- [ ] Improve performance of Schedule
- [ ] Implement Parser
- [ ] Implement Pipeline

