package model

import (
	"time"
)

type ProductContract struct {
	ID                        uint      `json:"id" gorm:"primarykey"`
	Name                      string    `json:"name" gorm:"type:varchar(80);not null"`
	Description               string    `json:"description" gorm:"type:longtext"`
	Img                       string    `json:"img" gorm:"type:varchar(500)"`
	TwitterName               string    `json:"twitter_name" gorm:"type:varchar(40)"`
	Status                    uint      `json:"status" gorm:"type:int(4);default:0;not null"`
	Amount                    string    `json:"amount" gorm:"type:varchar(40)"`
	SaleContractAddress       string    `json:"sale_contract_address" gorm:"type:varchar(42)"`
	TokenAddress              string    `json:"token_address" gorm:"type:varchar(42)"`
	PaymentToken              string    `json:"payment_token" gorm:"type:varchar(42)"`
	Follower                  uint      `json:"follower" gorm:"type:int(8);default:0;not null"`
	Tge                       time.Time `json:"tge" gorm:"type:datetime"`
	ProjectWebsite            string    `json:"project_website" gorm:"type:varchar(500)"`
	AboutHtml                 string    `json:"about_html" gorm:"type:varchar(500)"`
	RegistrationTimeStarts    time.Time `json:"registration_time_starts" gorm:"type:datetime"`
	RegistrationTimeEnds      time.Time `json:"registration_time_ends" gorm:"type:datetime"`
	SaleStart                 time.Time `json:"sale_start" gorm:"type:datetime"`
	SaleEnd                   time.Time `json:"sale_end" gorm:"type:datetime"`
	MaxParticipation          string    `json:"max_participation" gorm:"type:varchar(40)"`
	TokenPriceInPT            string    `json:"token_price_in_PT" gorm:"type:varchar(40)"`
	TotalTokensSold           string    `json:"total_tokens_sold" gorm:"type:varchar(40)"`
	AmountOfTokensToSell      string    `json:"amount_of_tokens_to_sell" gorm:"type:varchar(60)"`
	TotalRaised               string    `json:"total_raised" gorm:"type:varchar(60)"`
	Symbol                    string    `json:"symbol" gorm:"type:varchar(60)"`
	Decimals                  uint      `json:"decimals" gorm:"type:int(8);default:8"`
	UnlockTime                time.Time `json:"unlock_time" gorm:"type:datetime"`
	Medias                    string    `json:"medias" gorm:"type:varchar(200)"`
	NumberOfRegistrations     uint      `json:"number_of_registrations" gorm:"type:int(8)"`
	Vesting                   string    `json:"vesting" gorm:"type:varchar(40)"`
	Tricker                   string    `json:"tricker" gorm:"type:varchar(40)"`
	TokenName                 string    `json:"token_name" gorm:"type:varchar(20)"`
	Roi                       string    `json:"roi" gorm:"type:varchar(20)"`
	VestingPortionsUnlockTime string    `json:"vesting_portions_unlock_time" gorm:"type:varchar(60)"`
	VestingPercentPerPortion  string    `json:"vesting_percent_per_portion" gorm:"type:varchar(60)"`
	CreateTime                time.Time `json:"create_time" gorm:"type:datetime; not null"`
	UpdateTime                time.Time `json:"update_time" gorm:"type:datetime; not null"`
	Type                      uint      `json:"type" gorm:"type:int(8)"`
	CardLink                  string    `json:"card_link" gorm:"type:varchar(200)"`
	TweetId                   string    `json:"tweet_id" gorm:"type:varchar(40)"`
	ChainId                   uint      `json:"chain_id" gorm:"type:int(8)"`
	PaymentTokenDecimals      uint      `json:"payment_token_decimals" gorm:"type:int(8)"`
	CurrentPrice              uint      `json:"current_price" gorm:"type:bigint(12)"`
}
