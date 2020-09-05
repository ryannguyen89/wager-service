/*
 * Wager service APIs
 *
 * APIs for a wager system
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package controllers

import (
	"github.com/ryannguyen89/wager-service/business"
	"github.com/ryannguyen89/wager-service/models"
	"github.com/ryannguyen89/wager-service/models/errors"
	"github.com/ryannguyen89/wager-service/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListWagersGet - 
func ListWagersGet(c *gin.Context) {
	var (
		req models.ListWagerRequest
	)
	if err := c.ShouldBind(&req); err != nil {
		log.Printf("Bind request error: %v", err)
		c.JSON(http.StatusBadRequest, models.Error{Error:err.Error()})
		return
	}
	log.Printf("Request: %s", utils.JsonSerialize(req))

	// Validate request
	code, msg := validateListWagerRequest(&req)
	if code != http.StatusOK {
		log.Printf("Validate request error: %s", msg)
		c.JSON(code, models.Error{Error:msg})
		return
	}

	lWagers, code, msg := business.ListWagers(c.Request.Context(), &req)
	if code != http.StatusOK {
		c.JSON(code, models.Error{Error:msg})
		return
	}
	c.JSON(http.StatusOK, lWagers)
}

func validateListWagerRequest(r *models.ListWagerRequest) (code int, message string) {
	if r.Page <= 0 {
		code = http.StatusBadRequest
		message = errors.InvalidPage
		return
	}

	if r.Limit <= 0 {
		code = http.StatusBadRequest
		message = errors.InvalidLimit
		return
	}

	code = http.StatusOK
	return
}
