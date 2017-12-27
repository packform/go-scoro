package scoro

import (
	"errors"
)

// InvoiceLine struct represents invoice lines data type of Scoro API.
// https://api.scoro.com/api/#invoiceLinesApiDocs
type InvoiceLine struct {
	Id              int               `json:"id"`
	ProductID       int               `json:"product_id"`
	Comment         Strings           `json:"comment"`
	Comment2        Strings           `json:"comment2"`
	UnitPrice       Decimal           `json:"price"`
	Amount          Decimal           `json:"amount"`
	Amount2         Decimal           `json:"amount"`
	Discount        Decimal           `json:"discount"`
	Sum             Decimal           `json:"sum"`
	Vat             Decimal           `json:"vat"`
	Unit            string            `json:"unit"`
	FinanceObjectID int               `json:"finance_object_id"`
	Cost            Decimal           `json:"cost"`
	ProjectID       int               `json:"project_id"`
	CustomFields    map[string]string `json:"custom_fields,omitempty"`
}

// Invoice struct represents invoices data type of Scoro API.
// https://api.scoro.com/api/#invoicesApiDocs
type Invoice struct {
	Id                       *int              `json:"id,omitempty"`
	PaymentType              string            `json:"payment_type,omitempty"`
	Fine                     string            `json:"fine,omitempty"`
	QuoteID                  int               `json:"quote_id"`
	OrderID                  int               `json:"order_id"`
	PrepaymentPercent        float32           `json:"prepayment_percent,omitempty"`
	PrepaymentSum            Decimal           `json:"prepayment_sum,omitempty"`
	ReferenceNo              string            `json:"reference_no,omitempty"`
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
	module      string
}

// InvoicesAPI provides type safe wrappers for View/List/Modify/Delete actions
// of invoices API
func Invoices(credentials Credentials) InvoicesAPI {
	return InvoicesAPI{
		credentials: credentials,
		module:      "invoices",
	}
}

// InvoicesAPI provides type safe wrappers for View/List/Modify/Delete actions
// of prepayments API
// https://api.scoro.com/api/#prepaymentsApiDocs
func PrepaymentInvoices(credentials Credentials) InvoicesAPI {
	return InvoicesAPI{
		credentials: credentials,
		module:      "invoices/prepayments",
	}
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
	return NewRequest(t.credentials, t.module)
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
