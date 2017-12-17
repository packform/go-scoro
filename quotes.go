package scoro

import (
	"errors"
	"strconv"

	"github.com/shopspring/decimal"
)

type QuoteLine struct {
	ProductID int             `json:"product_id"`
	UnitPrice decimal.Decimal `json:"price"`
	Amount    decimal.Decimal `json:"amount"`
	Sum       decimal.Decimal `json:"sum"`
	Vat       decimal.Decimal `json:"vat"`
}

type Quote struct {
	Id                       *int              `json:"id,omitempty"`
	No                       string            `json:"no,omitempty"`
	Discount                 float32           `json:"discount,omitempty"`
	Discount2                float32           `json:"discount2,omitempty"`
	Discount3                float32           `json:"discount3,omitempty"`
	Sum                      decimal.Decimal   `json:"sum,omitempty"`
	VatSum                   decimal.Decimal   `json:"vat_sum,omitempty"`
	Vat                      decimal.Decimal   `json:"vat,omitempty"`
	CompanyID                int               `json:"company_id,omitempty"`
	PersonID                 int               `json:"person_id,omitempty"`
	CompanyAddressID         int               `json:"company_address_id,omitempty"`
	InterestedPartyID        int               `json:"interested_party_id,omitempty"`
	InterestedPartyAddressID int               `json:"interested_party_address_id,omitempty"`
	ProjectID                int               `json:"project_id,omitempty"`
	Currency                 string            `json:"currency,omitempty"`
	OwnerID                  int               `json:"owner_id,omitempty"`
	Status                   string            `json:"status,omitempty"`
	Description              string            `json:"description,omitempty"`
	IsSent                   Bool              `json:"is_sent"`
	AccountID                string            `json:"account_id,omitempty"`
	Lines                    []QuoteLine       `json:"lines,omitempty"`
	ModifiedDate             Time              `json:"modified_date,omitempty"`
	CustomFields             map[string]string `json:"custom_fields,omitempty"`
	IsDeleted                Bool              `json:"is_deleted"`
	DeletedDate              Time              `json:"deleted_date,omitempty"`
}

type QuoteList []Quote

type QuotesAPI struct {
	credentials Credentials
}

func Quotes(credentials Credentials) QuotesAPI {
	return QuotesAPI{credentials}
}

func (t QuotesAPI) View(id string) (*Quote, error) {
	resp, err := t.Request().SetResponse(quoteResponse{}).View(id)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*quoteResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Quote, nil
}

func (t QuotesAPI) List(filter interface{}, page int, count int) (*QuoteList, error) {
	resp, err := t.Request().SetResponse(quoteListResponse{}).List(filter, page, count)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*quoteListResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Quotes, nil
}

func (t QuotesAPI) Modify(product Quote) (*Quote, error) {
	resp, err := t.Request().SetResponse(quoteResponse{}).Modify(product)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*quoteResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Quote, nil
}

func (t QuotesAPI) Delete(id int) error {
	_, err := t.Request().SetResponse(quoteResponse{}).Delete(map[string]string{
		"id": strconv.Itoa(id),
	})

	return err
}

func (t QuotesAPI) Request() Request {
	return NewRequest(t.credentials, "quotes")
}

// Private

type quoteResponse struct {
	Response `json:",inline"`
	Quote    Quote `json:"data,omitempty"`
}

type quoteListResponse struct {
	Response `json:",inline"`
	Quotes   QuoteList `json:"data,omitempty"`
}

func (t quoteResponse) GetResponse() Response {
	return t.Response
}

func (t quoteListResponse) GetResponse() Response {
	return t.Response
}