package controller

import (
	"crawler/fetcher"
	"log"
)

func worker(request Request) (results Result, err error) {
	log.Printf("Fetch: %s\n", request.Url)
	content, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("Error: %s on %s", err, request.Url)
		return Result{}, err
	}

	return request.ParseFunc(content), nil
}
