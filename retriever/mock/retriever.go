package mock

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut time.Duration
}

func (r Retriever) Get(url string) string  {
	response, err := http.Get(url)
	defer response.Body.Close()
	if err!=nil{
		panic(err)
	}

	body, err := httputil.DumpResponse(response, true)
	if err!=nil{
		panic(err)
	}
	return string(body)
}