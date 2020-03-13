// Code generated by jsonschema. DO NOT EDIT.
package main

// ID: https://example.com/person.schema.json

type Person struct {
	// Age in years which must be equal to or greater than zero.
	Age *int32 `json:"Age,omitempty" validate:"gte=0"`
	// The person's first name.
	FirstName string `json:"FirstName,omitempty"`
	// The person's last name.
	LastName string `json:"LastName,omitempty"`
}
