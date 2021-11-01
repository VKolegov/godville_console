package utils

import (
	"errors"
	"fmt"
)

func ParseSavings(savingsString string) (int, error) {
	var (
		savings = 0
		a       string // a = тысяч
	)
	_, err := fmt.Sscanf(savingsString, "%d %s", &savings, &a)

	if err != nil {

		return -1, errors.New(
			fmt.Sprintf(
				"Ошибка при парсинге сбережений \"%s\": %s\n", savingsString, err.Error(),
			),
		)
	}

	return savings, nil
}
