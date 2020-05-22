package pagination

import "time"

type Payment struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	Name        string    `json:"name"`
	CreatedTime time.Time `json:"createdTime"`
}
