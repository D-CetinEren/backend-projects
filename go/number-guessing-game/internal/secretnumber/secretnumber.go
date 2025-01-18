package secretnumber

import (
	"fmt"
	"math/rand"
)

func SecretNumber() (result int) {
	result = rand.Intn(100) + 1
	fmt.Println("Secret number has been chosen!")
	return result
}
