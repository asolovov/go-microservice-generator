package models

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type User struct {
	UUID      uuid.UUID
	Name      string
	Address   common.Address
	Balance   *big.Int
	CreatedAt time.Time
}
