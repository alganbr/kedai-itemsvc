package managers

import (
	"github.com/alganbr/kedai-itemsvc/internal/models"
	"github.com/alganbr/kedai-itemsvc/internal/repos"
	"github.com/alganbr/kedai-utils/datetime"
	"github.com/alganbr/kedai-utils/errors"
	"net/http"
)

type IItemManager interface {
	Get(int64) (*models.Item, *errors.Error)
	Create(*models.ItemRq) (*models.Item, *errors.Error)
	Update(int64, *models.ItemRq) (*models.Item, *errors.Error)
	Patch(int64, *models.ItemRq) (*models.Item, *errors.Error)
}

type ItemManager struct {
	itemRepo repos.IItemRepository
}

func NewItemManager(itemRepo repos.IItemRepository) IItemManager {
	return &ItemManager{
		itemRepo: itemRepo,
	}
}

func (mgr *ItemManager) Get(id int64) (*models.Item, *errors.Error) {
	item, err := mgr.itemRepo.Get(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (mgr *ItemManager) Create(rq *models.ItemRq) (*models.Item, *errors.Error) {
	item := &models.Item{
		Name:        rq.Name,
		Description: rq.Description,
		CreatedAt:   datetime.GetUtcNow(),
		CreatedBy:   rq.RequestedBy,
		UpdatedAt:   datetime.GetUtcNow(),
		UpdatedBy:   rq.RequestedBy,
	}

	err := mgr.itemRepo.Create(item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (mgr *ItemManager) Update(id int64, rq *models.ItemRq) (*models.Item, *errors.Error) {
	item, err := mgr.Get(id)
	if err != nil {
		return nil, err
	}

	updated := false
	if rq.Name != item.Name {
		item.Name = rq.Name
		updated = true
	}
	if rq.Description != item.Description {
		item.Description = rq.Description
		updated = true
	}
	if !updated {
		return nil, &errors.Error{
			Code:    http.StatusNotModified,
			Message: "Not modified",
		}
	}

	item.UpdatedAt = datetime.GetUtcNow()
	item.UpdatedBy = rq.RequestedBy

	err = mgr.itemRepo.Update(item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (mgr *ItemManager) Patch(id int64, rq *models.ItemRq) (*models.Item, *errors.Error) {
	item, err := mgr.Get(id)
	if err != nil {
		return nil, err
	}

	if rq.Name == "" {
		rq.Name = item.Name
	}
	if rq.Description == "" {
		rq.Description = item.Description
	}

	item, err = mgr.Update(id, rq)
	if err != nil {
		return nil, err
	}

	return item, nil
}
