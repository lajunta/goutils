package html

import (
	"html/template"
)

// Safe return safe html
func Safe(s string) template.HTML {
	return template.HTML(s)
}

// URLSafe return safe url in href string
func URLSafe(s string) template.URL {
	return template.URL(s)
}
