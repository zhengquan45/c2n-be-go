package response

import (
	"c2n/model"
	"encoding/json"
)

type ProductContractResponse struct {
	ID                        uint    `json:"id"`
	Name                      string  `json:"name"`
	Description               string  `json:"description"`
	Img                       string  `json:"img"`
	Status                    uint    `json:"status"`
	Amount                    string  `json:"amount"`
	SaleContractAddress       string  `json:"sale_contract_address"`
	TokenAddress              string  `json:"token_address"`
	PaymentToken              string  `json:"payment_token"`
	Follower                  uint    `json:"follower"`
	Tge                       int64   `json:"tge"`
	ProjectWebsite            string  `json:"project_website"`
	AboutHtml                 string  `json:"about_html"`
	RegistrationTimeStarts    int64   `json:"registration_time_starts"`
	RegistrationTimeEnds      int64   `json:"registration_time_ends"`
	SaleStart                 int64   `json:"sale_start"`
	SaleEnd                   int64   `json:"sale_end"`
	MaxParticipation          string  `json:"max_participation"`
	TokenPriceInPT            string  `json:"token_price_in_PT"`
	TotalTokensSold           string  `json:"total_tokens_sold"`
	AmountOfTokensToSell      string  `json:"amount_of_tokens_to_sell"`
	TotalRaised               string  `json:"total_raised"`
	Symbol                    string  `json:"symbol"`
	Decimals                  uint    `json:"decimals"`
	UnlockTime                int64   `json:"unlock_time"`
	Medias                    string  `json:"medias"`
	NumberOfRegistrations     uint    `json:"number_of_registrations"`
	Vesting                   string  `json:"vesting"`
	Tricker                   string  `json:"tricker"`
	TokenName                 string  `json:"token_name"`
	Roi                       string  `json:"roi"`
	VestingPortionsUnlockTime []int64 `json:"vesting_portions_unlock_time"`
	VestingPercentPerPortion  []int64 `json:"vesting_percent_per_portion"`
	CreateTime                int64   `json:"create_time"`
	UpdateTime                int64   `json:"update_time"`
	Type                      uint    `json:"type"`
	CardLink                  string  `json:"card_link"`
	TweetId                   string  `json:"tweet_id"`
	ChainId                   uint    `json:"chain_id"`
	PaymentTokenDecimals      uint    `json:"payment_token_decimals"`
	CurrentPrice              uint    `json:"current_price"`
}

// ProductContractToProductContractResponse 转换 ProductContract 为 ProductContractResponse
func ProductContractToProductContractResponse(productContract model.ProductContract) ProductContractResponse {
	var vestingPortionsUnlockTime []int64
	_ = json.Unmarshal([]byte(productContract.VestingPortionsUnlockTime), &vestingPortionsUnlockTime)
	var vestingPercentPerPortion []int64
	_ = json.Unmarshal([]byte(productContract.VestingPercentPerPortion), &vestingPercentPerPortion)
	return ProductContractResponse{
		ID:                        productContract.ID,
		Name:                      productContract.Name,
		Description:               productContract.Description,
		Img:                       productContract.Img,
		Status:                    productContract.Status,
		Amount:                    productContract.Amount,
		SaleContractAddress:       productContract.SaleContractAddress,
		TokenAddress:              productContract.TokenAddress,
		PaymentToken:              productContract.PaymentToken,
		Follower:                  productContract.Follower,
		Tge:                       productContract.Tge.UnixMilli(),
		ProjectWebsite:            productContract.ProjectWebsite,
		AboutHtml:                 productContract.AboutHtml,
		RegistrationTimeStarts:    productContract.RegistrationTimeStarts.UnixMilli(),
		RegistrationTimeEnds:      productContract.RegistrationTimeEnds.UnixMilli(),
		SaleStart:                 productContract.SaleStart.UnixMilli(),
		SaleEnd:                   productContract.SaleEnd.UnixMilli(),
		MaxParticipation:          productContract.MaxParticipation,
		TokenPriceInPT:            productContract.TokenPriceInPT,
		TotalTokensSold:           productContract.TotalTokensSold,
		AmountOfTokensToSell:      productContract.AmountOfTokensToSell,
		TotalRaised:               productContract.TotalRaised,
		Symbol:                    productContract.Symbol,
		Decimals:                  productContract.Decimals,
		UnlockTime:                productContract.UnlockTime.UnixMilli(),
		Medias:                    productContract.Medias,
		NumberOfRegistrations:     productContract.NumberOfRegistrations,
		Vesting:                   productContract.Vesting,
		Tricker:                   productContract.Tricker,
		TokenName:                 productContract.TokenName,
		Roi:                       productContract.Roi,
		VestingPortionsUnlockTime: vestingPortionsUnlockTime,
		VestingPercentPerPortion:  vestingPercentPerPortion,
		CreateTime:                productContract.CreateTime.UnixMilli(),
		UpdateTime:                productContract.UpdateTime.UnixMilli(),
		Type:                      productContract.Type,
		CardLink:                  productContract.CardLink,
		TweetId:                   productContract.TweetId,
		ChainId:                   productContract.ChainId,
		PaymentTokenDecimals:      productContract.PaymentTokenDecimals,
		CurrentPrice:              productContract.CurrentPrice,
	}
}
