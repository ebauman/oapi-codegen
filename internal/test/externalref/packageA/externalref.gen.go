// Package packageA provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/ebauman/oapi-codegen DO NOT EDIT.
package packageA

import (
	externalRef0 "github.com/ebauman/oapi-codegen/internal/test/externalref/packageB"
)

// ObjectA defines model for ObjectA.
type ObjectA struct {
	Name    *string               `json:"name,omitempty"`
	ObjectB *externalRef0.ObjectB `json:"object_b,omitempty"`
}
