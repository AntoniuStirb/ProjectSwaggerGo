/*
 * DevNest Portal API (OpenAPI 3.0)
 *
 * The DevNest Portal API endpoints definitions based on the OpenAPI 3.0 specification.
 *
 * API version: 1.0.0
 * Contact: contact@devnest.ro
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

import (
	"time"
)

type NewProject struct {
	Name string `json:"name"`

	ClientName string `json:"clientName"`

	Description string `json:"description,omitempty"`

	StartingDate time.Time `json:"startingDate"`

	EndingDate time.Time `json:"endingDate,omitempty"`

	Currency string `json:"currency"`

	Status bool `json:"status,omitempty"`
}