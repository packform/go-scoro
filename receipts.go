package scoro

import (
	"errors"
)

// Receipt struct represents receipts data type of Scoro API.
// https://api.scoro.com/api/#receiptsApiDocs
type Receipt struct {
	Id           *int    `json:"receipt_id,omitempty"`
	Date         Date    `json:"date,omitempty"`
	InvoiceID    *int    `json:"invoice_id,omitempty"`
	PrepaymentID *int    `json:"prepayment_id,omitempty"`
	Sum          Decimal `json:"sum,omitempty"`
	SalesDocType string  `json:"sales_doc_type,omitempty"`
	ContactID    int     `json:"contact_id,omitempty"`
	ContactName  string  `json:"contact_name,omitempty"`
}
type ReceiptList []Receipt

// ReceiptsAPI provides type safe wrappers for View/List/Modify/Delete actions
// of receipts API
type ReceiptsAPI struct {
	credentials Credentials
}

func Receipts(credentials Credentials) ReceiptsAPI {
	return ReceiptsAPI{credentials}
}

func (t ReceiptsAPI) View(id string) (*Receipt, error) {
	resp, err := t.Request().SetResponse(receiptResponse{}).View(id)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*receiptResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Receipt, nil
}

func (t ReceiptsAPI) List(filter interface{}, page int, count int) (*ReceiptList, error) {
	resp, err := t.Request().SetResponse(receiptListResponse{}).List(filter, page, count)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*receiptListResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Receipts, nil
}

func (t ReceiptsAPI) Modify(receipt Receipt) (*Receipt, error) {
	resp, err := t.Request().SetResponse(receiptResponse{}).Modify(receipt)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*receiptResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Receipt, nil
}

func (t ReceiptsAPI) Delete(id int) error {
	_, err := t.Request().SetResponse(receiptResponse{}).Delete(id, nil)

	return err
}

func (t ReceiptsAPI) Request() Request {
	return NewRequest(t.credentials, "receipts")
}

// Private

type receiptResponse struct {
	ResponseHeader `json:",inline"`
	Receipt        Receipt `json:"data,omitempty"`
}

type receiptListResponse struct {
	ResponseHeader `json:",inline"`
	Receipts       ReceiptList `json:"data,omitempty"`
}

func (t receiptResponse) GetResponseHeader() ResponseHeader {
	return t.ResponseHeader
}

func (t receiptListResponse) GetResponseHeader() ResponseHeader {
	return t.ResponseHeader
}
