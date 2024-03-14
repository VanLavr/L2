package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type RequestBuilder struct {
	Method string
	URL    string
	Header http.Header
	Body   string
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{Header: make(http.Header)}
}

func (r *RequestBuilder) SetMethod(method string) *RequestBuilder {
	r.Method = strings.ToUpper(method)
	return r
}

func (r *RequestBuilder) SetURL(url string) *RequestBuilder {
	r.URL = url
	return r
}

func (r *RequestBuilder) SetHeader(key, value string) *RequestBuilder {
	r.Header.Add(key, value)
	return r
}

func (r *RequestBuilder) SetBody(body string) *RequestBuilder {
	r.Body = body
	return r
}

func (r *RequestBuilder) Build() (*http.Request, error) {
	req, err := http.NewRequest(r.Method, r.URL, strings.NewReader(r.Body))
	if err != nil {
		return nil, err
	}

	req.Header = r.Header

	return req, nil
}

func main() {
	request, err := NewRequestBuilder().
		SetURL("http://example.com").
		SetMethod("post").
		SetHeader("Content-Type", "application/json").
		SetBody("{\"message\": \"hello\"}").
		Build()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", request)
}
