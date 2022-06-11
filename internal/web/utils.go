package web

import (
	"fmt"
	"strconv"
	"strings"
)

func parseIntList(query string) ([]int, error) {
	ids := make([]int, 0)
	split := strings.Split(query, ",")
	for _, idStr := range split {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse id: %w", err)
		}
		ids = append(ids, id)
	}
	return ids, nil
}
