package configer

import (
	"encoding/json"
	"fmt"
	"os"
)

func decodeJSONConfig(v interface{}, filename string) error {
	fmt.Println("Decoding JSON")
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	return json.NewDecoder(f).Decode(v)
}