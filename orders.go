package scoro

import (
	"errors"
)

// OrderLine struct represents order lines data type of Scoro API.
// https://api.scoro.com/api/#orderLinesApiDocs
type OrderLine struct {
	ProductID int     `json:"product_id"`
	UnitPrice Decimal `json:"price"`
	Amount    Decimal `json:"amount"`
	Sum       Decimal `json:"sum"`
	Vat       Decimal `json:"vat"`
	Comment   Strings `json:"comment"`
}

// Order struct represents orders data type of Scoro API.
// https://api.scoro.com/api/#ordersApiDocs
type Order struct {
	Id                       *int              `json:"id,omitempty"`
	QuoteID                  int               `json:"quote_id"`
	No                       string            `json:"no,omitempty"`
	Discount                 float32           `json:"discount,omitempty"`
	Discount2                float32           `json:"discount2,omitempty"`
	Discount3                float32           `json:"discount3,omitempty"`
	Sum                      Decimal           `json:"sum,omitempty"`
	VatSum                   Decimal           `json:"vat_sum,omitempty"`
	Vat                      Decimal           `json:"vat,omitempty"`
	CompanyID                int               `json:"company_id,omitempty"`
	PersonID                 int               `json:"person_id,omitempty"`
	CompanyAddressID         int               `json:"company_address_id,omitempty"`
	InterestedPartyID        int               `json:"interested_party_id,omitempty"`
	InterestedPartyAddressID int               `json:"interested_party_address_id,omitempty"`
	ProjectID                int               `json:"project_id,omitempty"`
	Currency                 string            `json:"currency,omitempty"`
	OwnerID                  int               `json:"owner_id,omitempty"`
	Date                     Date              `json:"date,omitempty"`
	Deadline                 Date              `json:"deadline,omitempty"`
	Status                   string            `json:"status,omitempty"`
	Description              string            `json:"description,omitempty"`
	IsSent                   Bool              `json:"is_sent"`
	Lines                    []OrderLine       `json:"lines,omitempty"`
	ModifiedDate             Time              `json:"modified_date,omitempty"`
	CustomFields             map[string]string `json:"custom_fields,omitempty"`
	IsDeleted                Bool              `json:"is_deleted"`
	DeletedDate              Time              `json:"deleted_date,omitempty"`
}
type OrderList []Order

// OrdersAPI provides type safe wrappers for View/List/Modify/Delete actions
// of orders API
type OrdersAPI struct {
	credentials Credentials
}

func Orders(credentials Credentials) OrdersAPI {
	return OrdersAPI{credentials}
}

func (t OrdersAPI) View(id string) (*Order, error) {
	resp, err := t.Request().SetResponse(orderResponse{}).View(id)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*orderResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Order, nil
}

func (t OrdersAPI) List(filter interface{}, page int, count int) (*OrderList, error) {
	resp, err := t.Request().SetResponse(orderListResponse{}).List(filter, page, count)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*orderListResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Orders, nil
}

func (t OrdersAPI) Modify(product Order) (*Order, error) {
	resp, err := t.Request().SetResponse(orderResponse{}).Modify(product)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*orderResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Order, nil
}

func (t OrdersAPI) Delete(id int) error {
	_, err := t.Request().SetResponse(orderResponse{}).Delete(id, nil)

	return err
}

func (t OrdersAPI) Request() Request {
	return NewRequest(t.credentials, "orders")
}

// Private

type orderResponse struct {
	ResponseHeader `json:",inline"`
	Order          Order `json:"data,omitempty"`
}

type orderListResponse struct {
	ResponseHeader `json:",inline"`
	Orders         OrderList `json:"data,omitempty"`
}

func (t orderResponse) GetResponseHeader() ResponseHeader {
	return t.ResponseHeader
}

func (t orderListResponse) GetResponseHeader() ResponseHeader {
	return t.ResponseHeader
}
