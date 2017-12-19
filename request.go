package scoro

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	resty "gopkg.in/resty.v1"
)

// Request helps to build and send custom request to Scoro API. It supports
// automatic mappings of data structures into request body and from response.
//
// It is unlikely that you will need to use this type directly. Most of the time
// service wrappers should be used instead. But it could be helpful
// in some case, for instance, if some new request isn't implemented yet.
//
// Example:
//
// response := NewRequest(credentials).SetResponse(ProductResponse{}).View(id)
type Request struct {
	credentials Credentials
	lang        string
	respType    ResponseType
	entityType  string
}

// ResponseHeader represents base part of API response that is common for
// all responses.
//
// Example of defining new response type:
//
//		type SomeNewResponse{
//			RequestHeader `json:",inline"`
//			Data 					SomeDataType `json:"data,omitempty"`
//	 	}
type ResponseHeader struct {
	Status     string `json:"status"`
	StatusCode string `json:"statusCode"`
	Messages   *struct {
		Error []string `json:"error"`
	} `json:"messages,omitempty"`
}

// ResponseType interface provides methods for retrieving information common for all
// responses. Each response type must implement this interface in order to be
// compatible with Request type.
type ResponseType interface {
	GetResponseHeader() ResponseHeader
}

// NewRequest creates request object configured with specified credentials and
// pointed to the endpoint corresponding to the specified entityType.
//
// entityType can be "products", "orders", "invoices" or any other type supported
// by Scoro API
func NewRequest(credentials Credentials, entityType string) Request {
	return Request{
		credentials: credentials,
		lang:        DefaultLang,
		entityType:  entityType,
	}
}

// SetResponse method is to register the response object for automatic unmarshalling
// of JSON responses. Response type shoul conforms to the ResponseType interface.
//
//		request.SetResult(ProductResponse{})
//
// Accessing a result value
//		if resp, err := request.View(id); err != nil {
//			product := resp.(*productResponse).Product
//		}
func (t Request) SetResponse(response ResponseType) Request {
	t.respType = response
	return t
}

// View method sends "view" action request
func (t Request) View(id string) (interface{}, error) {
	url := makeUrl(t.credentials.CompanyID, t.entityType, "view", id)
	body := requestBody{Credentials: t.credentials, Lang: t.lang}

	return sendRequest(url, body, t.respType)
}

// List method sends "list" action request
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

// Modify method sends "modify" action request
func (t Request) Modify(obj interface{}) (interface{}, error) {
	url := makeUrl(t.credentials.CompanyID, t.entityType, "modify")
	body := requestBody{Credentials: t.credentials, Lang: t.lang, Request: obj}

	return sendRequest(url, body, t.respType)
}

// Delete method sends "delete" action request
func (t Request) Delete(id int, filter interface{}) (interface{}, error) {
	url := makeUrl(t.credentials.CompanyID, t.entityType, "delete", strconv.Itoa(id))
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

	response, validFormat := restyResp.Result().(ResponseType)
	if !validFormat {
		return nil, errors.New("Invalid format")
	}

	header := response.GetResponseHeader()

	if header.Status == "OK" {
		return response, nil
	}

	if header.Messages != nil {
		errStr := strings.Join(header.Messages.Error, "; ")
		return nil, errors.New(errStr)
	}

	return nil, errors.New("Error: " + header.StatusCode)
}
