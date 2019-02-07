package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	specialChars := struct {
		Ampersand         string `json:"ampersand"`
		LeftAngleBracket  string `json:"left_angle_bracket"`
		RightAngleBracket string `json:"right_angle_bracket"`
	}{
		Ampersand:         "&",
		LeftAngleBracket:  "<",
		RightAngleBracket: ">",
	}

	bytes, _ := json.Marshal(specialChars)
	fmt.Println(string(bytes))

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(specialChars)

	encoder.SetEscapeHTML(false)
	encoder.Encode(specialChars)
}
