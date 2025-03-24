package utils

import (
	"errors"
	"fmt"
	"os"
)

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprintf("%v", value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func ReadFloatFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 1000, errors.New("error reading balance file")
	}
	valueText := string(data)
	var value float64
	fmt.Sscan(valueText, &value)
	return value, nil
}
