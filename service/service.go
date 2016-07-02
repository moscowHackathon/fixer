package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"io"
	"io/ioutil"
)

const ContentTypeJSON = "application/json"


type HttpTransporter interface {
	Do(*http.Request) (*http.Response, error)
}

type fixerService struct {
	baseUrl string
	transport HttpTransporter
}


type IFixerService interface {
	Start(req GetRequest) (GetResponse, error)
}

type GetRequest struct {
	ChanId      string  `json:"chan_id"`
	UserMessage string  `json:"user_message"`
}

type GetResponse struct {
	//nextQuestion string `json:"user_message"`
	ID int64 `json:"ID"`
}


func NewClient(baseUrl string, transport HttpTransporter) (IFixerService, error) {
	serviceClient := &fixerService{baseUrl: baseUrl, transport: transport}
	return serviceClient, nil
}

////RESTCall rest call
func (this fixerService) restCall(method, urlStr string, request interface{}, response interface{}) error {

	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequest(method, urlStr, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", ContentTypeJSON)
	httpReq.Header.Set("Accept-Encoding", "identity")

	httpResp, err := this.transport.Do(httpReq)

	if httpResp != nil {
		defer ClearBody(httpResp.Body)
	}

	if err != nil {
		return err
	}

	var cliError error

	if httpResp.StatusCode != 200 {
		if err := json.NewDecoder(httpResp.Body).Decode(&cliError); err != nil {
			return err
		}
		return cliError
	}

	return json.NewDecoder(httpResp.Body).Decode(response)
}


func ClearBody(body io.ReadCloser) {
	if body == nil {
		return
	}
	var bodyBytes []byte
	bodyBytes, _ = ioutil.ReadAll(body)
	ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	body.Close()
}

func (this fixerService)Start(request GetRequest) (GetResponse, error) {
	response := GetResponse{}

	err := this.restCall("GET", this.baseUrl + "/start", request, &response)

	return response, err
}
