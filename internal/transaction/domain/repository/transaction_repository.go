package repository

import (
	"context"

	"github.com/LexBokun/transaction-service/internal/pkg/lists/status"
	"github.com/LexBokun/transaction-service/internal/transaction/domain/entity"
)

type TransactionRepository interface {
	Create(ctx context.Context, tx *entity.Transaction) (string, error)
	Get(ctx context.Context, id string) (*entity.Transaction, error)
	List(ctx context.Context, limit int64, cursor string) (*[]entity.Transaction, string, error)
	Update(ctx context.Context, id string, status status.Status) error
}
