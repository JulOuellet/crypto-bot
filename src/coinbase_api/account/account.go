package account

import (
    "time"
)

type Balance struct {
   Value    string `json:"value"`
   Currency string  `json:"currency"`
}

type Account struct {
   UUID		    string     `json:"uuid"`
   Name             string     `json:"name"`
   Currency         string     `json:"currency"`
   AvailableBalance Balance    `json:"available_balance"`
   Default          bool       `json:"default"`
   Active           bool       `json:"active"`
   CreatedAt        time.Time  `json:"created_at"`
   UpdatedAt        time.Time  `json:"updated_at"`
   DeletedAt        *time.Time `json:"deleted_at"`
   Type             string     `json:"type"`
   Ready            bool       `json:"ready"`
   Hold             Balance    `json:"hold"`
}

func (a *Account) HasBalance() bool {
	return a.AvailableBalance.Value != "0"
}

