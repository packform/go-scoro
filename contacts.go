package scoro

import (
	"errors"
)

// Contact struct represents contacts data type of Scoro API.
// https://api.scoro.com/api/#contactsApiDocs
type Contact struct {
	ContactID      *int              `json:"contact_id,omitempty"`
	Name           string            `json:"name,omitempty"`
	Lastname       string            `json:"lastname,omitempty"`
	ContactType    string            `json:"contact_type,omitempty"`
	IdCode         string            `json:"id_code,omitempty"`
	BankAccount    string            `json:"bankaccount,omitempty"`
	Birthday       Date              `json:"birthday,omitempty"`
	Position       string            `json:"position,omitempty"`
	Comments       string            `json:"comments,omitempty"`
	Sex            string            `json:"sex,omitempty"`
	VatNo          string            `json:"vatno,omitempty"`
	Timezone       string            `json:"timezone,omitempty"`
	ManagerID      int               `json:"manager_id,omitempty"`
	IsSupplier     Bool              `json:"is_supplier,omitempty"`
	IsClient       Bool              `json:"is_client,omitempty"`
	ModifiedDate   Time              `json:"modified_date,omitempty"`
	Addresses      []Address         `json:"addresses,omitempty"`
	MeansOfContact MeansOfContact    `json:"means_of_contact,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	ReferenceNo    string            `json:"reference_no,omitempty"`
	CustomFields   map[string]string `json:"custom_fields,omitempty"`
	IsDeleted      Bool              `json:"is_deleted"`
}
type ContactList []Contact

type Address struct {
	Country      string `json:"country,omitempty"`
	County       string `json:"county,omitempty"`
	Municipality string `json:"municipality,omitempty"`
	City         string `json:"city,omitempty"`
	Street       string `json:"street,omitempty"`
	ZipCode      string `json:"zipcode,omitempty"`
}

type MeansOfContact struct {
	Mobiles  []string `json:"mobile,omitempty"`
	Phones   []string `json:"phone,omitempty"`
	Emails   []string `json:"email,omitempty"`
	Websites []string `json:"website,omitempty"`
	Skypes   []string `json:"skype,omitempty"`
	Faxes    []string `json:"fax,omitempty"`
}

// ContactsAPI provides type safe wrappers for View/List/Modify/Delete actions
// of contacts API
type ContactsAPI struct {
	credentials Credentials
}

func Contacts(credentials Credentials) ContactsAPI {
	return ContactsAPI{credentials}
}

func (t ContactsAPI) View(id string) (*Contact, error) {
	resp, err := t.Request().SetResponse(contactResponse{}).View(id)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*contactResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Contact, nil
}

func (t ContactsAPI) List(filter interface{}, page int, count int) (*ContactList, error) {
	resp, err := t.Request().SetResponse(contactListResponse{}).List(filter, page, count)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*contactListResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Contacts, nil
}

func (t ContactsAPI) Modify(product Contact) (*Contact, error) {
	resp, err := t.Request().SetResponse(contactResponse{}).Modify(product)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*contactResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Contact, nil
}

func (t ContactsAPI) Delete(id int) error {
	_, err := t.Request().SetResponse(contactResponse{}).Delete(id, nil)

	return err
}

func (t ContactsAPI) Request() Request {
	return NewRequest(t.credentials, "contacts")
}

// Private

type contactResponse struct {
	ResponseHeader `json:",inline"`
	Contact        Contact `json:"data,omitempty"`
}

type contactListResponse struct {
	ResponseHeader `json:",inline"`
	Contacts       ContactList `json:"data,omitempty"`
}

func (t contactResponse) GetResponseHeader() ResponseHeader {
	return t.ResponseHeader
}

func (t contactListResponse) GetResponseHeader() ResponseHeader {
	return t.ResponseHeader
}
