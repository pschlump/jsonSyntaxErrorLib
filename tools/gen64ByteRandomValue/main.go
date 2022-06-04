package main

import (
	"fmt"
	"math/rand"
	"time"
)

// export QR_JWT_KEY="23032b350f2bfca57e8363067b718fa6b35a862854f020c90fa9b28150d2bf2efbd3925a06d2ef478a96c3b4627d223d2e3dfd212e05984fadcbf59ad067a380"
//                    49ccce7aa7d103455b156eee00105f71bf560178068e15ae60d98bb4e36d3cfc34b72ef263d099bc2f1e230fd3e7717453ab2fbe574c62b81d708875333a266d

func main() {
	s := RandomSecret(128)
	fmt.Printf("%s\n", s)
}

// RandomSecret will generate a random secret of given length
func RandomSecret(length int) string {
	bytes := make([]rune, length)

	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("0123456789abcdef")
	for i := range bytes {
		// TODO - replace with cryptograpic random generation!
		bytes[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(bytes)
}
