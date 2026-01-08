package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type File struct {
	Base  string
	Date  string
	Rates map[string]float32
}

func quickCheck(e error) {
	if e != nil {
		panic(e)
	}
}

// ./conversor <amount> <coin>
func main() {
	// checking args amount and coin
	amount, err := strconv.ParseFloat(os.Args[1], 8)
	quickCheck(err)
	coin := os.Args[2]

	path := filepath.Join("./rates.json")
	jsonFile, err := os.ReadFile(path)
	quickCheck(err)

	var encodedJson File
	err = json.Unmarshal(jsonFile, &encodedJson)
	quickCheck(err)

	valueConverted := float32(amount) * encodedJson.Rates[coin]
	fmt.Printf("%.2f BRL = %.2f %s\n", float32(amount), valueConverted, coin)
}
