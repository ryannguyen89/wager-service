package models

type ListWagerRequest struct {
	Page int32 `form:"page"`
	Limit int32 `form:"limit"`
}
