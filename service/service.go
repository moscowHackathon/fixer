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
	Start(GetRequestAll) (GetResponse, error)
	Question(GetRequestAll) (GetResponse, error)
	Answer(GetRequestAnswer) (GetResponse, error)
	Complete(GetRequestAll) (GetResponse, error)
}

type GetRequestAll struct {
	ChanId      string  `json:"id"`
}

type GetRequestAnswer struct {
	ChanId      string  `json:"id"`
	Answer string  `json:"answer"`
}

type GetResponse struct {
	//nextQuestion string `json:"user_message"`
	ID string `json:"id"`
	Message string `json:"message"`
	Error string `json:"error"`
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

func (this fixerService)Start(request GetRequestAll) (GetResponse, error) {
	response := GetResponse{}

	err := this.restCall("GET", this.baseUrl + "/start/" + request.ChanId + "/", nil, &response)

	return response, err
}

func (this fixerService)Question(request GetRequestAll) (GetResponse, error) {
	response := GetResponse{}

	err := this.restCall("GET", this.baseUrl + "/question/" + request.ChanId + "/", nil , &response)

	return response, err
}

func (this fixerService)Complete(request GetRequestAll) (GetResponse, error) {
	response := GetResponse{}

	err := this.restCall("GET", this.baseUrl + "/complete/" + request.ChanId + "/", nil , &response)

	return response, err
}

func (this fixerService)Answer(request GetRequestAnswer) (GetResponse, error) {
	response := GetResponse{}

	err := this.restCall("POST", this.baseUrl + "/answer/" + request.ChanId + "/", request.Answer , &response)

	return response, err
}
