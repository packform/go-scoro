package scoro

import (
	"errors"
)

// Relation struct represents relations data type of Scoro API.
// https://api.scoro.com/api/#relationsApiDocs
type Relation struct {
	ObjectID int `json:"object_id,omitempty"`

	// RelatedObjects is map containing related objects ID-s for each type on
	// list request and list of ID-s on modify/delete request.
	RelatedObjects interface{} `json:"related_objects,omitempty"`
	Type           string      `json:"type	String,omitempty"`
}
type RelationList []Relation

// RelationsAPI provides type safe wrappers for View/List/Modify/Delete actions
// of relations API
type RelationsAPI struct {
	credentials Credentials
}

func Relations(credentials Credentials) RelationsAPI {
	return RelationsAPI{credentials}
}

func (t RelationsAPI) List(filter interface{}, page int, count int) (*RelationList, error) {
	resp, err := t.Request().SetResponse(relationListResponse{}).List(filter, page, count)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*relationListResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Relations, nil
}

func (t RelationsAPI) Modify(relation Relation) (*Relation, error) {
	resp, err := t.Request().SetResponse(relationResponse{}).Modify(relation)
	if err != nil {
		return nil, err
	}

	result, ok := resp.(*relationResponse)
	if !ok {
		return nil, errors.New("Invalid response format")
	}

	return &result.Relation, nil
}

func (t RelationsAPI) Delete(id int) error {
	_, err := t.Request().SetResponse(relationResponse{}).Delete(id, nil)

	return err
}

func (t RelationsAPI) Request() Request {
	return NewRequest(t.credentials, "relations")
}

// Private

type relationResponse struct {
	ResponseHeader `json:",inline"`
	Relation       Relation `json:"data,omitempty"`
}

type relationListResponse struct {
	ResponseHeader `json:",inline"`
	Relations      RelationList `json:"data,omitempty"`
}

func (t relationResponse) GetResponseHeader() ResponseHeader {
	return t.ResponseHeader
}

func (t relationListResponse) GetResponseHeader() ResponseHeader {
	return t.ResponseHeader
}
