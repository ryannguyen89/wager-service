/*
 * Wager service APIs
 *
 * APIs for a wager system
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wager

type CreateWageRequest struct {

	TotalWagerValue int32 `json:"total_wager_value,omitempty"`

	Odds int32 `json:"odds,omitempty"`

	SellingPercentage int32 `json:"selling_percentage,omitempty"`

	SellingPrice float32 `json:"selling_price,omitempty"`
}
