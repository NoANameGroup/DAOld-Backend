package elder

import (
	"time"
)

type ElderVO struct {
	Username          string    `json:"username"`
	BlockChainAddress string    `json:"blockChainAddress"`
	Balance           float64   `json:"balance"`
	CreatedAt         time.Time `json:"createdAt"`
}
