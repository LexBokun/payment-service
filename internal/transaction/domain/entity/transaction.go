package entity

import (
	"github.com/LexBokun/transaction-service/internal/pkg/lists/status"
	"github.com/LexBokun/transaction-service/internal/pkg/lists/types"
	"time"
)

type Transaction struct {
	ID        string
	Type      types.Type
	Status    status.Status
	From      Operation
	To        Operation
	CreatedAt time.Time
}
