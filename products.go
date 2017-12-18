package scoro

import (
	"errors"
	"strconv"

	"github.com/shopspring/decimal"
)

// Product struct represents products data type of Scoro API.
// https://api.scoro.com/api/#productsApiDocs
type Product struct {
	Id                *int              `json:"product_id,omitempty"`
	Code              string            `json:"code,omitempty"`
	Name              string            `json:"name,omitempty"`
	Names             Strings           `json:"names,omitempty"`
	Price             decimal.Decimal   `json:"price,omitempty"`
	BuyingPrice       decimal.Decimal   `json:"buying_price,omitempty"`
	Description       Strings           `json:"description,omitempty"`
	Description2      Strings           `json:"description2,omitempty"`
	Tag               string            `json:"tag,omitempty"`
	Url               string            `json:"url,omitempty"`
	SupplierID        int               `json:"supplier_id,omitempty"`
	ProductGroupID    int               `json:"productgroup_id,omitempty"`
	IsActive          Bool              `json:"is_active"`
	IsService         Bool              `json:"is_service"`
	DefaultVatCodeID  int               `json:"default_vat_code_id,omitempty"`
	AccountinObjectID int               `json:"accounting_object_id,omitempty"`
	ModifiedDate      Time              `json:"modified_date,omitempty"`
	CustomFields      map[string]string `json:"custom_fields,omitempty"`
	IsDeleted         Bool              `json:"is_deleted"`
	DeletedDate       Time              `json:"deleted_date,omitempty"`
}
type ProductList []Product

// ProductsAPI provides type safe wrappers for View/List/Modify/Delete actions
// of products API
type ProductsAPI struct {
	credentials Credentials
}

func Products(credentials Credentials) ProductsAPI {
	return ProductsAPI{credentials}
}

func (t ProductsAPI) View(id string) (*Product, error) {
	resp, err := t.Request().SetResponse(productResponse{}).View(id)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*productResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Product, nil
}

func (t ProductsAPI) List(filter interface{}, page int, count int) (*ProductList, error) {
	resp, err := t.Request().SetResponse(productListResponse{}).List(filter, page, count)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*productListResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Products, nil
}

func (t ProductsAPI) Modify(product Product) (*Product, error) {
	resp, err := t.Request().SetResponse(productResponse{}).Modify(product)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*productResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Product, nil
}

func (t ProductsAPI) Delete(id int) error {
	_, err := t.Request().SetResponse(productResponse{}).Delete(map[string]string{
		"product_id": strconv.Itoa(id),
	})

	return err
}

func (t ProductsAPI) Request() Request {
	return NewRequest(t.credentials, "products")
}

// Private

type productResponse struct {
	ResponseHeader `json:",inline"`
	Product        Product `json:"data,omitempty"`
}

type productListResponse struct {
	ResponseHeader `json:",inline"`
	Products       ProductList `json:"data,omitempty"`
}

func (t productResponse) GetResponseHeader() ResponseHeader {
	return t.ResponseHeader
}

func (t productListResponse) GetResponseHeader() ResponseHeader {
	return t.ResponseHeader
}
