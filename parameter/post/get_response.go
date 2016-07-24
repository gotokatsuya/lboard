package post

import (
	usermodel "github.com/gotokatsuya/lboard/model/user"
)

// GetResponse ...
type GetResponse struct {
	Instance instance
}

// NewGetResponse ...
func NewGetResponse() *GetResponse {
	return &GetResponse{}
}

// Convert ...
func (r *GetResponse) Convert(user *usermodel.Entity) error {
	ins := instance{}
	if err := ins.Convert(user); err != nil {
		return err
	}
	r.Instance = ins
	return nil
}
