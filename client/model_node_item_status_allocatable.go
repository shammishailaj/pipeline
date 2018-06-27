/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: 0.3.0
 * Contact: info@banzaicloud.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

type NodeItemStatusAllocatable struct {
	Cpu string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
	Pods string `json:"pods,omitempty"`
}