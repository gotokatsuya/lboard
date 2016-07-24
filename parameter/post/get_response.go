package post

import (
	postmodel "github.com/gotokatsuya/lboard/model/post"
)

// GetResponse ...
type GetResponse struct {
	Instances []instance
}

// NewGetResponse ...
func NewGetResponse() *GetResponse {
	return &GetResponse{}
}

// Convert ...
func (r *GetResponse) Convert(posts []postmodel.Entity) error {
	instances := make([]instance, len(posts))
	for i, post := range posts {
		var ins instance
		if err := ins.Convert(&post); err != nil {
			return err
		}
		instances[i] = ins
	}
	r.Instances = instances
	return nil
}
