package configer

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

func decodeXMLConfig(v interface{}, filename string) error {
	fmt.Println("Decoding XML")
	file, err := os.Open(filename)

	if err != nil{
		log.Fatal("open error")
		return err
	}

	err = xml.NewDecoder(file).Decode(v)
	return err

}

