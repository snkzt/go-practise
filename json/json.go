package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Built-in json library support custom data type
// Only exported fields will be encoded/decoded in JSON.
// Fields must start with Capital letters to be exported.
// To be exported Both struct name and field name must be sart with capital letter.
type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// Encoding Boolean, int, float and string to JSON strings
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))
	// Slices and maps will be encoded to JSON arrays and objects
	slcB, _ := json.Marshal([]string{"apple", "peach", "kiwi"}) // Output will be an array
	fmt.Println(string(slcB))

	mapB, _ := json.Marshal(map[string]int{"apple": 5, "Tomato": 7}) // Output will be an object
	fmt.Println(string(mapB))
	// Custom data type will automatically be encoded by JSON package.
	// It will only include exportyed fields in the encoded output and
	// Use those names as the JSON keys.
	res1B, _ := json.Marshal(&response1{Page: 1, Fruits: []string{"apple", "peach", "kiwi"}})
	fmt.Println(string(res1B))
	// You can use tags on struct field declarations to customise
	// the encoded JSON key names.
	// Response2 struct has tags for each field as `json:"page"` and `json:"fruits"` respectively.
	// JSON key names for the output will be "page" and "fruits" which start with non-capital letters.
	res2B, _ := json.Marshal(&response2{Page: 1, Fruits: []string{"grapes", "peach", "pasteque"}})
	fmt.Println(string(res2B))

	// Decoding JSON
	byt := []byte(`{"num":6.13,"strings":["a","b"]}`)
	// The JSON package can put the decided data in dat which will hold a map of strings
	// to arbitrary datatypes.
	var dat map[string]interface{}
	// Actual decoding and a check for associated errors.
	// The data of byt will ve stored in pointed dat map.
	if err := json.Unmarshal(byt, &dat); err != nil {
		fmt.Errorf("Error occured when decoding JSON file: %w", err)
	}
	fmt.Println(dat)
	// Type assertion: This is to use the values in the decoded map
	// Retrieving num of dat and assign in a variable num
	num := dat["num"].(float64)
	fmt.Println(num, dat)
	// Accessing neested (a value in an array in a map value) data requires
	// a series of conversions.
	strings := dat["strings"].([]interface{})
	str1 := strings[0].(string)
	fmt.Println(strings, str1)
	// Decode JSON into custom data types.
	str := `{"page": 1, "fruits": ["peach", "grapes"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res) // Store decoded data str into the custom type res
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
	// JSON encodings can be streamed directly to os.Writers (os.Stdout or HTTP response body)
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"peach": 5, "tomato": 7} // ??? tomato will be printed first when make it Tomato
	enc.Encode(d)
}
