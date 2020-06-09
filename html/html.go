package html

import (
	"html/template"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Safe return safe html
func Safe(s string) template.HTML {
	return template.HTML(s)
}

// URLSafe return safe url in href string
func URLSafe(s string) template.URL {
	return template.URL(s)
}

// Hex convert objectID to string
func Hex(id primitive.ObjectID) string {
	return id.Hex()
}
