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

type Shop struct {

	Id string `json:"id,omitempty"`

	Name string `json:"name"`

	Address string `json:"address,omitempty"`

	Sailer *Sailer `json:"sailer,omitempty"`
}
