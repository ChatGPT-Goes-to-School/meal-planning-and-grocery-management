package utils

import (
	"strconv"
)

func ConvertParamToInt(param string) (int, error) {
    id, err := strconv.Atoi(param)
    if err != nil {
        return 0, err
    }
    return id, nil
}