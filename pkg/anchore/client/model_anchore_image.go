/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.12
 * Contact: nurmi@anchore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client
import (
	"time"
)
// AnchoreImage A unique image in the engine. May have multiple tags or references. Unique to an image content across registries or repositories.
type AnchoreImage struct {
	// A metadata content record for a specific image, containing different content type entries
	ImageContent map[string]interface{} `json:"image_content,omitempty"`
	// Details specific to an image reference and type such as tag and image source
	ImageDetail []ImageDetail `json:"image_detail,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	ImageDigest string `json:"imageDigest,omitempty"`
	UserId string `json:"userId,omitempty"`
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// State of the image
	ImageStatus string `json:"image_status,omitempty"`
	// A state value for the current status of the analysis progress of the image
	AnalysisStatus string `json:"analysis_status,omitempty"`
}
