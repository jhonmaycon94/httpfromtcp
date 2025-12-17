package request

import (
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func parseRequestLine(requestLine string) (*RequestLine, error) {
	parts := strings.Split(requestLine, " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("Expected 3 parts, but got %d part(s)", len(parts))
	}
	requestLineStruct := new(RequestLine)
	requestLineStruct.Method = parts[0]
	requestLineStruct.RequestTarget = parts[1]
	// requestLineStruct.HttpVersion = parts[2]

	httpVersion := strings.Split(parts[2], "/")
	requestLineStruct.HttpVersion = httpVersion[1]

	return requestLineStruct, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	request, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	requestLine, _, _ := strings.Cut(string(request), "\r\n")
	requestLineStruct, err := parseRequestLine(requestLine)
	if err != nil {
		return nil, err
	}

	requestStruct := new(Request)
	requestStruct.RequestLine = *requestLineStruct
	return requestStruct, nil
}
