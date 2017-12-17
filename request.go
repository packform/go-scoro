package scoro

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	resty "gopkg.in/resty.v1"
)

type Request struct {
	credentials Credentials
	lang        string
	respType    ResponseType
	entityType  string
}

type Response struct {
	Status     string `json:"status"`
	StatusCode string `json:"statusCode"`
	Messages   *struct {
		Error []string `json:"error"`
	} `json:"messages,omitempty"`
}

type ResponseType interface {
	GetResponse() Response
}

func NewRequest(credentials Credentials, entityType string) Request {
	return Request{
		credentials: credentials,
		lang:        DefaultLang,
		entityType:  entityType,
	}
}

func (t Request) SetResponse(response ResponseType) Request {
	t.respType = response
	return t
}

func (t Request) View(id string) (interface{}, error) {
	url := makeUrl(t.credentials.CompanyID, t.entityType, "view", id)
	body := requestBody{Credentials: t.credentials, Lang: t.lang}

	return sendRequest(url, body, t.respType)
}

func (t Request) List(filter interface{}, page int, count int) (interface{}, error) {
	url := makeUrl(t.credentials.CompanyID, t.entityType, "list")
	body := requestBody{
		Credentials: t.credentials,
		Lang:        t.lang,
		Filter:      filter,
		Page:        page,
		PerPage:     count,
	}

	return sendRequest(url, body, t.respType)
}

func (t Request) Modify(obj interface{}) (interface{}, error) {
	url := makeUrl(t.credentials.CompanyID, t.entityType, "modify")
	body := requestBody{Credentials: t.credentials, Lang: t.lang, Request: obj}

	return sendRequest(url, body, t.respType)
}

func (t Request) Delete(filter interface{}) (interface{}, error) {
	url := makeUrl(t.credentials.CompanyID, t.entityType, "delete")
	body := requestBody{Credentials: t.credentials, Lang: t.lang, Request: filter}

	return sendRequest(url, body, t.respType)
}

// Private

type requestBody struct {
	Credentials `json:",inline"`
	Lang        string      `json:"lang"`
	Page        int         `json:"page"`
	PerPage     int         `json:"per_page"`
	Request     interface{} `json:"request,omitempty"`
	Filter      interface{} `json:"filter,omitempty"`
}

func makeUrl(companyId string, entityType string, action string, params ...string) string {
	baseURL := fmt.Sprintf("https://%v.scoro.com/api/v1", companyId)

	urlParts := []string{baseURL, entityType, action}
	urlParts = append(urlParts, params...)

	return strings.Join(urlParts, "/")
}

func sendRequest(url string, body interface{}, respType ResponseType) (interface{}, error) {
	req := resty.R()
	req = req.ExpectContentType("application/json")
	req = req.SetResult(respType)

	if body != nil {
		req = req.SetBody(body)
	}

	resp, err := req.Post(url)
	if err != nil {
		return nil, err
	}

	return unmarshalResponse(resp)
}

func unmarshalResponse(restyResp *resty.Response) (interface{}, error) {
	if restyResp.StatusCode() != http.StatusOK {
		return nil, errors.New("Error status: " + restyResp.Status())
	}

	if _, ok := restyResp.Result().(ResponseType); !ok {
		return nil, errors.New("Invalid format")
	}

	response := restyResp.Result().(ResponseType).GetResponse()

	if response.Status == "OK" {
		return restyResp.Result(), nil
	}

	if response.Messages != nil {
		errStr := strings.Join(response.Messages.Error, "; ")
		return nil, errors.New(errStr)
	}

	return nil, errors.New("Error: " + response.StatusCode)
}
