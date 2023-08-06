package webbasics

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Address struct {
	Firstline string
	Lastline  string
	State     string
	Pincode   uint32
}

type User struct {
	// Since this struct will be read by JSON(another package),
	// made the variables exported (Starts with an uppercase)
	Name        string
	Email       string
	Age         int
	UserAddress Address
	Timestamp   time.Time
	Options     []string
}

func jsonEncodeExample() {
	u := User{
		Name:  "John Doe",
		Email: "j@j.com",
		Age:   123,
		UserAddress: Address{
			Firstline: "45 Braodway St",
			Lastline:  "Test lastline",
			State:     "NY",
			Pincode:   4783578,
		},
		Timestamp: time.Now(),
		Options:   []string{"option1", "option3"},
	}

	// Marshalling: converting memory represented data to a form
	// that could be used for communication and storage
	// If valid, types are converted to equivalent JSON types
	// Note: Default zeros are ignored IG
	uJson, err := json.Marshal(u) // returns byte arr and an error
	check(err)

	fmt.Printf("nJSON representation of %#vu is: %#vn", u, string(uJson))
}

func jsonDecodeExample() {
	// Since here we have backquote string, no need to explicitly escape special chars
	encodedStr := `{"Name":"John Doe",
					"Email":"j@j.com",
					"Age":123,
					"UserAddress":{
						"Firstline":"45 Braodway St",
						"Lastline":"Test lastline",
						"State":"NY",
						"Pincode":4783578},
					"Timestamp":"2023-08-06T07:02:11.876958336+05:30",
					"Options":["option1","option3"]}`

	var u User

	// Decode using type of destination var
	err := json.Unmarshal([]byte(encodedStr), &u)
	check(err)
	fmt.Printf("\nDecoded JSON object: %#v\n", u)

	// Unmarshall will only match those fields for which
	// there's a type definition  in dest struct
	type GuestUser struct {
		Name string
	}

	var guest GuestUser

	err = json.Unmarshal([]byte(encodedStr), &guest)
	check(err)
	fmt.Printf("\nDecoded JSON object: %#v\n", guest)

	// Generic JSOn blob
	// json Unmarshalls Generic JSON blobs to
	// map[string]interface{} when object and []interface when string
	var genericUser interface{}
	err = json.Unmarshal([]byte(encodedStr), &genericUser)
	check(err)
	fmt.Printf("\nDecoded JSON object: %#v\n", genericUser)

	// Assert whether the decoded object is arr or obj
	obj := genericUser.(map[string]interface{})

	// iterate over the map
	for k, v := range obj {
		fmt.Printf("\nKey: %#v, val: %#v\n", k, v)

		// Type assertion
		switch vv := v.(type) {
		case map[string]interface{}:
			fmt.Printf("\n%v is an object with value: %v\n", k, vv)
		case []interface{}:
			fmt.Printf("\n%v is an arr with value: %v\n", k, vv)

		case string:
			fmt.Printf("\n%v is a stirng with value: %v\n", k, vv)

		case int:
			fmt.Printf("\n%v is an integer with value: %v\n", k, vv)

		default:
			fmt.Println("Unknown Type")

		}

	}

	// Unmarshal with references
	//  If there's an object/value corresponding to the type in dest struct
	// unmarshall will create it in memory and add reference else it will be nil
	type Admin struct {
		Name        string
		Credentials *string
	}
	var admin Admin
	err = json.Unmarshal([]byte(encodedStr), &admin)
	check(err)
	fmt.Printf("\nDecoded JSON object: %#v\n", admin)

}

func jsonEncodeMap() {

	mp := map[string]interface{}{
		"Name":  "John",
		"Age":   12,
		"Links": []string{"github", "twitter", "facebook"},
	}

	mapToJSON, err := json.Marshal(mp)
	check(err)

	fmt.Printf("\nJSON from go map: %#v is: \n %#v\n", mp, string(mapToJSON))

}

func getWriter() {
	_, err := os.OpenFile("data/test1.json", os.O_RDONLY, 0776)
	check(err)
}

func jsonEncoderDecoder() {
	fmt.Println("JSON encoder decoder")

	// Decode: read from io.Reader stream
	fileReader, err := os.OpenFile("data/test1.json", os.O_RDONLY, 0776)
	check(err)
	defer fileReader.Close()

	dec := json.NewDecoder(fileReader) // Decoder requires io.Reader

	var r any
	err = dec.Decode(&r)
	check(err)

	fmt.Printf("\nRead from file: %#v\n", r)

	// Encoder: write to a stream
	fileWriter, err := os.OpenFile("data/test2.json", os.O_CREATE|os.O_WRONLY, 0776)
	check(err)
	enc := json.NewEncoder(fileWriter)
	enc.Encode(&r)

}

func JSONImpl() {
	// jsonEncodeExample()
	// jsonDecodeExample()
	// jsonEncodeMap()
	jsonEncoderDecoder()

}
