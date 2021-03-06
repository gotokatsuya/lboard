package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/gotokatsuya/lboard/model"
)

// Repository ...
type Repository struct {
	*model.RootRepository
}

// NewRepository ...
func NewRepository(ctx *gin.Context) *Repository {
	return &Repository{
		RootRepository: model.NewRootRepository(ctx, new(Entity)),
	}
}

// GetByID ...
func (r *Repository) GetByID(id int64) (*Entity, error) {
	ent := new(Entity)
	ent.UserID = id
	if err := r.RootRepository.GetByPrimary(ent); err != nil {
		return nil, err
	}
	return ent, nil
}

// Create inserts the entity
func (r *Repository) Create(ent *Entity) (*Entity, error) {
	if err := r.RootRepository.Create(ent); err != nil {
		return nil, err
	}
	return ent, nil
}

// Update ...
func (r *Repository) Update(ent *Entity) (*Entity, error) {
	if err := r.RootRepository.UpdateByPrimary(ent); err != nil {
		return nil, err
	}
	return ent, nil
}

// FindByIDs ...
func (r *Repository) FindByIDs(ids []int64) ([]Entity, error) {
	var ents []Entity
	if err := r.DB.Where("user_id IN (?)", ids).Find(&ents).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ents, nil
}
