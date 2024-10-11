package tools

import (
	"net/url"
    "strings"
	"fmt"
)

func SanitizeInput(input string) string {
    return input[:len(input)-1] // Remove the newline character
}

func CheckForValidUrl(inputUrl string) bool {
    _, err := url.ParseRequestURI(inputUrl)
    return err == nil
}

func ToSSHFormat(url string) string {
	if strings.HasPrefix(url, "git@") {
		return url
	}

	if strings.HasPrefix(url, "https://") {
		parts := strings.Split(url, "/")
		if len(parts) >= 3 {
			domain := parts[2]
			sshURL := strings.Replace(url, fmt.Sprintf("https://%s/", domain), fmt.Sprintf("git@%s:", domain), 1)
			if !strings.HasSuffix(sshURL, ".git") {
				sshURL += ".git"
			}
			return sshURL
		}
	}
	return url
}