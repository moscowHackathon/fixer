package service

import (
	"log"
	"net/http"
)

func Start(channelID string) (GetResponse, error) {

	client, err := NewClient("http://localhost:8080", &http.Client{})
	if err != nil {
		log.Printf("Cannot initialize client: %v", err)
		return GetResponse{}, err
	}

	response, err := client.Start(GetRequestAll{ChanId:channelID})

	if err != nil {
		log.Printf("Client error occurred: %v", err)
		return GetResponse{}, err
	}
	log.Printf("Start response:\n%v", response)
	log.Println("--------------------------------------------------------")

	return response, nil
}

func Question(channelID string)(GetResponse, error)  {
	client, err := NewClient("http://localhost:8080", &http.Client{})
	if err != nil {
		log.Printf("Cannot initialize client: %v", err)
		return GetResponse{}, err
	}

	response, err := client.Question(GetRequestAll{ChanId:channelID})

	if err != nil {
		log.Printf("Client error occurred: %v", err)
		return GetResponse{}, err
	}
	log.Printf("Start response:\n%v", response)
	log.Println("--------------------------------------------------------")

	return response, nil

}