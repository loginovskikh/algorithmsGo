package generator

import (
	"errors"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateInt(length int) ([]int, error) {
	if length < 1 {
		return nil, errors.New("data length in generator < 1")
	}

	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = rand.Int()
	}

	return res, nil
}
