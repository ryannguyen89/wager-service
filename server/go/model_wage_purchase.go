/*
 * Wager service APIs
 *
 * APIs for a wager system
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package wager

type WagePurchase struct {

	Id int64 `json:"id,omitempty"`

	WagerId int64 `json:"wager_id,omitempty"`

	BuyingPrice float32 `json:"buying_price,omitempty"`

	BoughtAt int64 `json:"bought_at,omitempty"`
}
