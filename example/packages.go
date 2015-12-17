package example

import (
	"fmt"
	"math/rand"
	"time"
)

func PakcagesMain() {
	rand.Seed(int64(time.Now().Nanosecond()))
	fmt.Println("My favorite number is", rand.Intn(10))
}
