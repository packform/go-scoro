package scoro

import (
	"errors"

	"github.com/shopspring/decimal"
)

// PrepaymentLine struct represents prepayment lines data type of Scoro API.
// https://api.scoro.com/api/#prepaymentLinesApiDocs
type PrepaymentLine struct {
	ProductID int             `json:"product_id"`
	UnitPrice decimal.Decimal `json:"price"`
	Amount    decimal.Decimal `json:"amount"`
	Sum       decimal.Decimal `json:"sum"`
	Vat       decimal.Decimal `json:"vat"`
	Comment   Strings         `json:"comment"`
}

// Prepayment struct represents prepayments data type of Scoro API.
// https://api.scoro.com/api/#prepaymentsApiDocs
type Prepayment struct {
	Id                       *int              `json:"id,omitempty"`
	PaymentType              string            `json:"payment_type,omitempty"`
	Fine                     string            `json:"fine,omitempty"`
	QuoteID                  int               `json:"quote_id"`
	OrderID                  int               `json:"order_id"`
	PrepaymentSum            float32           `json:"prepayment_sum,omitempty"`
	ReferenceNo              string            `json:"reference_no,omitempty"`
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
	Date                     Time              `json:"date,omitempty"`
	Deadline                 Time              `json:"deadline,omitempty"`
	Status                   string            `json:"status,omitempty"`
	Description              string            `json:"description,omitempty"`
	IsSent                   Bool              `json:"is_sent"`
	Lines                    []PrepaymentLine  `json:"lines,omitempty"`
	ModifiedDate             Time              `json:"modified_date,omitempty"`
	CustomFields             map[string]string `json:"custom_fields,omitempty"`
	IsDeleted                Bool              `json:"is_deleted"`
	DeletedDate              Time              `json:"deleted_date,omitempty"`
}
type PrepaymentList []Prepayment

// PrepaymentsAPI provides type safe wrappers for View/List/Modify/Delete actions
// of prepayments API
type PrepaymentsAPI struct {
	credentials Credentials
}

func Prepayments(credentials Credentials) PrepaymentsAPI {
	return PrepaymentsAPI{credentials}
}

func (t PrepaymentsAPI) View(id string) (*Prepayment, error) {
	resp, err := t.Request().SetResponse(prepaymentResponse{}).View(id)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*prepaymentResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Prepayment, nil
}

func (t PrepaymentsAPI) List(filter interface{}, page int, count int) (*PrepaymentList, error) {
	resp, err := t.Request().SetResponse(prepaymentListResponse{}).List(filter, page, count)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*prepaymentListResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Prepayments, nil
}

func (t PrepaymentsAPI) Modify(product Prepayment) (*Prepayment, error) {
	resp, err := t.Request().SetResponse(prepaymentResponse{}).Modify(product)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*prepaymentResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Prepayment, nil
}

func (t PrepaymentsAPI) Delete(id int) error {
	_, err := t.Request().SetResponse(prepaymentResponse{}).Delete(id, nil)
	return err
}

func (t PrepaymentsAPI) Request() Request {
	return NewRequest(t.credentials, "prepayments")
}

// Private

type prepaymentResponse struct {
	ResponseHeader `json:",inline"`
	Prepayment     Prepayment `json:"data,omitempty"`
}

type prepaymentListResponse struct {
	ResponseHeader `json:",inline"`
	Prepayments    PrepaymentList `json:"data,omitempty"`
}

func (t prepaymentResponse) GetResponseHeader() ResponseHeader {
	return t.ResponseHeader
}

func (t prepaymentListResponse) GetResponseHeader() ResponseHeader {
	return t.ResponseHeader
}
