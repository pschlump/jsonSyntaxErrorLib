package main

import (
	"fmt"
	"os"

	"github.com/pschlump/dbgo"
	"github.com/pschlump/mapstructure"
)

func main() {
	fmt.Printf("----------------------- t1 ----------------------- \n")
	t1()
	fmt.Printf("----------------------- t2 ----------------------- \n")
	t2()
	fmt.Printf("----------------------- t3 ----------------------- \n")
	t2()
}

func t1() {
	type Person struct {
		Name   string
		Age    int
		Emails []string
		Extra  map[string]string
	}

	// This input can come from anywhere, but typically comes from
	// something like decoding JSON where we're not quite sure of the
	// struct initially.
	input := map[string]interface{}{
		"name":   "Mitchell",
		"age":    91,
		"emails": []string{"one", "two", "three"},
		"extra": map[string]string{
			"twitter": "mitchellh",
		},
	}

	var result Person
	err := mapstructure.Decode(input, &result)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", dbgo.SVarI(result))

	// Output:
	/*
		{
			"Name": "Mitchell",
			"Age": 91,
			"Emails": [
				"one",
				"two",
				"three"
			],
			"Extra": {
				"twitter": "mitchellh"
			}
		}
	*/
}

func t2() {
	type Person struct {
		Name string
		Age  int
	}

	// This input can come from anywhere, but typically comes from
	// something like decoding JSON where we're not quite sure of the
	// struct initially.
	input := map[string]interface{}{
		"name":  "Mitchell",
		"age":   91,
		"email": "foo@bar.com",
	}

	// For metadata, we make a more advanced DecoderConfig so we can
	// more finely configure the decoder that is used. In this case, we
	// just tell the decoder we want to track metadata.
	var md mapstructure.Metadata
	var result Person
	config := &mapstructure.DecoderConfig{
		Metadata: &md,
		Result:   &result,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	if err := decoder.Decode(input); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Unused keys: %s\n", dbgo.SVarI(md.Unused))
	// Output:
	// Unused keys: []string{"email"}
	fmt.Printf("Matched Data: %s\n", dbgo.SVarI(result))
	// Output:
	//
}

func t3() {
	from := map[string]string{
		"name":  "Mitchell",
		"ages":  "91",
		"email": "foo@bar.com",
	}
	type Person struct {
		Name string
		Age  string
	}
	var result Person

	err := copy_to_struct(from, &result)
	if err != nil {
		fmt.Printf("Test3: err %s\n", err)
	} else {

	}
}

func copy_to_struct(from map[string]string, result interface{}) (err error) {

	// This input can come from anywhere, but typically comes from
	// something like decoding JSON where we're not quite sure of the
	// struct initially.
	input := make(map[string]interface{})
	for k, v := range from {
		input[k] = v
	}

	// For metadata, we make a more advanced DecoderConfig so we can
	// more finely configure the decoder that is used. In this case, we
	// just tell the decoder we want to track metadata.
	var md mapstructure.Metadata
	// var result Person
	config := &mapstructure.DecoderConfig{
		Metadata: &md,
		Result:   result,
	}

	decoder, e0 := mapstructure.NewDecoder(config)
	if e0 != nil {
		err = e0
		dbgo.Printf("Error: %s at:%(LF)\n", err)
		return
	}

	if err = decoder.Decode(input); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("Unused keys: %s\n", dbgo.SVarI(md.Unused))
	// Output:
	// Unused keys: []string{"email"}
	fmt.Printf("Matched Data: %s\n", dbgo.SVarI(result))
	// Output:
	//

	return
}
