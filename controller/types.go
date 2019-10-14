package controller

/* the request body that controller send to fetcher */
type Request struct {
	Url string
	ParseFunc func([]byte) Result // the function to parse results
}

type Result struct {
	Requests []Request
	Content []interface{}
}
