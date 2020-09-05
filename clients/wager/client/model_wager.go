/*
 * Wager service APIs
 *
 * APIs for a wager system
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wager
// Wager struct for Wager
type Wager struct {
	Id int64 `json:"id"`
	TotalWagerValue int32 `json:"total_wager_value"`
	Odds int32 `json:"odds"`
	SellingPercentage int32 `json:"selling_percentage"`
	SellingPrice float32 `json:"selling_price"`
	CurrentSellingPrice float32 `json:"current_selling_price"`
	PercentageSold int32 `json:"percentage_sold"`
	AmountSold float64 `json:"amount_sold"`
	PlacedAt int64 `json:"placed_at"`
}
