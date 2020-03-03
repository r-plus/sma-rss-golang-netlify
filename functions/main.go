package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// inject compiled golang version
var version string

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	q := request.QueryStringParameters
	artistNumber := q["sma"]
	fmt.Printf("query=%s, num=%s", q, artistNumber)

	if artistNumber == "" {
		return &events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       "Please add ?sma=XXX (XXX is artist number) querystring.",
		}, nil
	}

	feed, err := makeSMAAtomFeed(artistNumber)
	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       feed,
	}, nil
}

func makeSMAAtomFeed(artistID string) (string, error) {
	resp, err := http.Get("http://www.sma.co.jp/artist/json/info/" + artistID)
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	prefix := "callback"
	jsonp := bytes.NewBuffer(body)
	decoder := json.NewDecoder(&JSONPWrapper{Prefix: prefix, Underlying: jsonp})
	var json SMAJsonp
	decoder.Decode(&json)
	return AtomFeed{ArtistID: artistID, Info: json.Info}.MakeFeed(), nil
}

func main() {
	fmt.Printf("compiled go version: %s\n", version)
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
