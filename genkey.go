package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func genkey() string {
	rand.Seed(time.Now().UnixNano())

	lcnkey := ""
	for anu := 0; anu < 4; anu++ {
		block := make([]byte, 4)
		for j := 0; j < 4; j++ {
			charSet := []byte(strings.ToUpper("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"))
			block[j] = charSet[rand.Intn(len(charSet))]
		}
		lcnkey += string(block) + "-"
	}
	lcnkey = lcnkey[:len(lcnkey)-1]
	return lcnkey
}

func main() {
	fmt.Println(genkey())
}
