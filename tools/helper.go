package tools

import (
	"net/url"
)

func SanitizeInput(input string) string {
    return input[:len(input)-1] // Remove the newline character
}

func CheckForValidUrl(inputUrl string) bool {
    _, err := url.ParseRequestURI(inputUrl)
    return err == nil
}