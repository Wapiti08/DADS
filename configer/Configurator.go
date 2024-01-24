package configer

import (
	"errors"
	"reflect"
)

// Configuration file types
const (
	// untyped interger ordinal number
	CUSTOM uint8 = iota
	JSON
	XML
)

// define the error type --- return value
var errorType = errors.New("Type must be a pointer or a struct")

// read the configure file and fills the provided struct with configuration parameters
func GetConfiguration(confType uint8, obj interface{}, filename string) error {

	myValue := reflect.ValueOf(obj)
	// check if this is pointer type
	if myValue.Kind() != reflect.Ptr || myValue.IsNil() {
		return errorType
	}

	myType := reflect.TypeOf(obj)
	// check whether the value is struct type
	if myType.Kind() != reflect.Struct {
		return errorType
	}

	// check the conf type with switch
	switch confType {
	case CUSTOM:
		err = MarshalCustomConfig(obj, filename)
	case JSON:
		err = decodeJSONConfig(obj, filename)
	case XML:
		err = decodeXMLConfig(obj, filename)
	}
	return err
}
