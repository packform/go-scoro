package scoro

import (
	"errors"

	"github.com/shopspring/decimal"
)

// InvoiceLine struct represents invoice lines data type of Scoro API.
// https://api.scoro.com/api/#invoiceLinesApiDocs
type InvoiceLine struct {
	ProductID int             `json:"product_id"`
	UnitPrice decimal.Decimal `json:"price"`
	Amount    decimal.Decimal `json:"amount"`
	Sum       decimal.Decimal `json:"sum"`
	Vat       decimal.Decimal `json:"vat"`
	Comment   Strings         `json:"comment"`
}

// Invoice struct represents invoices data type of Scoro API.
// https://api.scoro.com/api/#invoicesApiDocs
type Invoice struct {
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
	Lines                    []InvoiceLine     `json:"lines,omitempty"`
	ModifiedDate             Time              `json:"modified_date,omitempty"`
	CustomFields             map[string]string `json:"custom_fields,omitempty"`
	IsDeleted                Bool              `json:"is_deleted"`
	DeletedDate              Time              `json:"deleted_date,omitempty"`
}
type InvoiceList []Invoice

// InvoicesAPI provides type safe wrappers for View/List/Modify/Delete actions
// of invoices API
type InvoicesAPI struct {
	credentials Credentials
}

func Invoices(credentials Credentials) InvoicesAPI {
	return InvoicesAPI{credentials}
}

func (t InvoicesAPI) View(id string) (*Invoice, error) {
	resp, err := t.Request().SetResponse(invoiceResponse{}).View(id)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*invoiceResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Invoice, nil
}

func (t InvoicesAPI) List(filter interface{}, page int, count int) (*InvoiceList, error) {
	resp, err := t.Request().SetResponse(invoiceListResponse{}).List(filter, page, count)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*invoiceListResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Invoices, nil
}

func (t InvoicesAPI) Modify(product Invoice) (*Invoice, error) {
	resp, err := t.Request().SetResponse(invoiceResponse{}).Modify(product)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*invoiceResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Invoice, nil
}

func (t InvoicesAPI) Delete(id int) error {
	_, err := t.Request().SetResponse(invoiceResponse{}).Delete(id, nil)

	return err
}

func (t InvoicesAPI) Request() Request {
	return NewRequest(t.credentials, "invoices")
}

// Private

type invoiceResponse struct {
	ResponseHeader `json:",inline"`
	Invoice        Invoice `json:"data,omitempty"`
}

type invoiceListResponse struct {
	ResponseHeader `json:",inline"`
	Invoices       InvoiceList `json:"data,omitempty"`
}

func (t invoiceResponse) GetResponseHeader() ResponseHeader {
	return t.ResponseHeader
}

func (t invoiceListResponse) GetResponseHeader() ResponseHeader {
	return t.ResponseHeader
}
