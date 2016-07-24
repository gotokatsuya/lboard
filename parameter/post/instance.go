package post

import (
	"encoding/json"

	postmodel "github.com/gotokatsuya/lboard/model/post"
)

type instance struct {
	PostID int64  `json:"post_id"`
	Text   string `json:"text"`
}

// Convert ...
func (i *instance) Convert(post *postmodel.Entity) error {
	b, err := json.Marshal(post)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, i); err != nil {
		return err
	}
	return nil
}
