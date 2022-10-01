package utils

import (
	"fmt"
	"time"
)

func ParseTimeUntil(date time.Time) (string, error) {
	now := time.Now().Unix()
	seconds := date.Unix() - now
	hours := int(seconds / 3600)
	days := int(hours / 24)

	var timeuntil string
	if days > 0 {
		timeuntil = fmt.Sprintf("%v jours", days)
	} else if hours > 0 {
		timeuntil = fmt.Sprintf("%v heures", hours)
	} else if seconds > 0 {
		timeuntil = "<1 heure"
	} else {
		timeuntil = "[terminé]"
	}
	return timeuntil, nil
}

func ParseTimeAgo(layout string, value string) (string, error) {
	now := time.Now().Unix()
	date, err := time.Parse(layout, value)
	if err != nil {
		return "", fmt.Errorf("failed to parse date: %w", err)
	}

	seconds := now - date.Unix()
	hours := int(seconds / 3600)
	days := int(hours / 24)

	var timeago string
	if days > 0 {
		timeago = fmt.Sprintf("%v jours", days)
	} else if hours > 0 {
		timeago = fmt.Sprintf("%v heures", hours)
	} else {
		timeago = "à l'instant"
	}
	return timeago, nil
}
