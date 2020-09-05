/*
 * Wager service APIs
 *
 * APIs for a wager system
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type CreateWagerRequest struct {

	TotalWagerValue int32 `json:"total_wager_value" binding:"required"`

	Odds int32 `json:"odds" binding:"required"`

	SellingPercentage int32 `json:"selling_percentage" binding:"required"`

	SellingPrice float32 `json:"selling_price" binding:"required"`
}