// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Endpoint struct {
	EndpointID   string `json:"endpoint_id"`
	EndpointName string `json:"endpoint_name"`
	Entries      int    `json:"entries"`
}

type UpdateEndpointinput struct {
	EndpointName *string `json:"endpoint_name,omitempty"`
}
