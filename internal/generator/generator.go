package generator

import (
	"errors"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateInt(length, min, max int) ([]int, error) {
	if length < 1 {
		return nil, errors.New("data length in generator < 1")
	}
	if min > max {
		return nil, errors.New("invalid arguments order in generator")
	}

	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = rand.Intn(max-min) + min
	}

	return res, nil
}
