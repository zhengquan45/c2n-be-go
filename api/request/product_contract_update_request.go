package request

import (
	"time"
)

type ProductContractUpdateRequest struct {
	ID                uint      `json:"id" binding:"required"`
	SaleAddress       string    `json:"sale_address"`
	SaleToken         string    `json:"sale_token"`
	SaleOwner         string    `json:"sale_owner"`
	TokenPriceInPT    string    `json:"token_price_in_PT"`
	PaymentToken      string    `json:"payment_token"`
	TotalToken        string    `json:"total_token"`
	SaleEndTime       time.Time `json:"sale_end_time"`
	TokensUnlockTime  time.Time `json:"tokens_unlock_time"`
	RegistrationStart time.Time `json:"registration_start"`
	RegistrationEnd   time.Time `json:"registration_end"`
	SaleStartTime     time.Time `json:"sale_start_time"`
}
