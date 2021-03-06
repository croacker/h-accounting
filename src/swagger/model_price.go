/*
 * OFD Checks API
 *
 * This is a OFD Checks API
 *
 * API version: 1.0.0
 * Contact: croacker@homeworld.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type Price struct {

	Id string `json:"id,omitempty"`

	Price float32 `json:"price"`

	DateTime time.Time `json:"dateTime,omitempty"`

	Product *Product `json:"product,omitempty"`

	Sailer *Sailer `json:"sailer,omitempty"`
}
