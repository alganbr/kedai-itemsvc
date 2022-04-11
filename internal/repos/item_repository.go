package repos

import (
	"context"
	"github.com/alganbr/kedai-itemsvc/internal/databases"
	"github.com/alganbr/kedai-itemsvc/internal/models"
	"github.com/alganbr/kedai-utils/errors"
	"github.com/jackc/pgx/v4"
	"net/http"
)

type IItemRepository interface {
	Get(int64) (*models.Item, *errors.Error)
	GetByEmail(string) (*models.Item, *errors.Error)
	Create(*models.Item) *errors.Error
	Update(*models.Item) *errors.Error
}

type ItemRepository struct {
	db *databases.DB
}

func NewItemRepository(db *databases.DB) IItemRepository {
	return &ItemRepository{
		db: db,
	}
}

func (repo *ItemRepository) Get(id int64) (*models.Item, *errors.Error) {
	var item models.Item
	err := repo.db.Pool.QueryRow(context.Background(), getItemQuery, id).Scan(
		&item.ItemId,
		&item.Name,
		&item.Description,
		&item.CreatedAt,
		&item.CreatedBy,
		&item.UpdatedAt,
		&item.UpdatedBy,
	)

	if err == pgx.ErrNoRows {
		return nil, &errors.Error{
			Code:    http.StatusNoContent,
			Message: err.Error(),
		}
	} else if err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &item, nil
}

func (repo *ItemRepository) GetByEmail(email string) (*models.Item, *errors.Error) {
	var item models.Item
	err := repo.db.Pool.QueryRow(context.Background(), getItemByEmailQuery, email).Scan(
		&item.ItemId,
		&item.Name,
		&item.Description,
		&item.CreatedAt,
		&item.CreatedBy,
		&item.UpdatedAt,
		&item.UpdatedBy,
	)

	if err == pgx.ErrNoRows {
		return nil, &errors.Error{
			Code:    http.StatusNoContent,
			Message: err.Error(),
		}
	} else if err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &item, nil
}

func (repo *ItemRepository) Create(item *models.Item) *errors.Error {
	err := repo.db.Pool.QueryRow(context.Background(), createItemQuery,
		item.Name,
		item.Description,
		item.CreatedAt,
		item.CreatedBy,
		item.UpdatedAt,
		item.UpdatedBy,
	).Scan(&item.ItemId)

	if err != nil {
		return &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (repo *ItemRepository) Update(item *models.Item) *errors.Error {
	_, err := repo.db.Pool.Exec(context.Background(), updateItemQuery,
		item.Name,
		item.Description,
		item.UpdatedAt,
		item.UpdatedBy,
		item.ItemId,
	)

	if err != nil {
		return &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

const (
	getItemQuery = `
		SELECT 
			id,
			name,
			description,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM item
		WHERE item_id = $1
	`

	getItemByEmailQuery = `
		SELECT 
			id,
			name,
			description,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM item
		WHERE email = $1
	`

	createItemQuery = `
		INSERT INTO item (
			name,
			description,
			created_at,
			created_by,
			updated_at,
			updated_by
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		) RETURNING id
	`

	updateItemQuery = `
		UPDATE item
		SET
			name = $1,
			description = $2,
			updated_at = $3,
			updated_by = $4
		WHERE item_id = $5
	`
)
