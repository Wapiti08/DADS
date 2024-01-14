package configer

/*
used to convert the context in config file to variables

*/

import (
	"reflect"
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
	"fmt"
)

type ConfigFields map[string]reflect.Value

func (f ConfigFields) Add(name, value, t string) error {
	switch t {
	case "STRING":
		f[name] = reflect.ValueOf((value))
	case "INTEGER":
		// convert to type int
		i, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		f[name] = reflect.ValueOf(i)
	case "FLOAT":
		fl, err := strconv.ParseFloat(value,64)
		if err != nil {
			return err
		}
		f[name] = reflect.ValueOf(fl)
	case "BOOL":
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		f[name] = reflect.ValueOf(b)
	}
	return nil
}

func MarshalCustomConfig(v reflect.Value, filename string) error {
	/* read config file and build config dictionary
	*/
	// avoid panic to the rest service if problems occur in field parse
	defer func(){
		if r := recover(); r!=nil {
			fmt.Println("Panic occurred", r)
		}
	}()

	if !v.CanSet() {
		fmt.Println("Value passed not settable")
	}
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()
	// make the dict to save parsed result
	fields := make(ConfigFields)
	// read file and process line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// avoid panic
		if strings.Count(line, "|") != 1 || strings.Count(line, ";") != 1 {
			continue
		}
		// split according to '|'
		args := strings.Split(line, "|")
		// further split example and type
		valuestype := strings.Split(args[1],";")
		// trim the space
		name, value, type := strings.TrimSpace(args[0]), strings.TrimSpace(valuetype[0]), strings.ToUpper(strings.TrimSpace(valuetype[1]))
		// save split results to dict
		fields.Add(name, value, type)
	}
	
	if err := scanner.Err(); err!= nil {
		return err
	}
	// dynamically update the values of the struct fields
	vt := v.Type()
	for i:=0 ; i< v.NumField(); i++ {
		fieldType := vt.Field(i)
		fieldValue := v.Field(i)
	
		name := fieldType.Tag.Get("name")

		if name == "" {
			name = fieldType.Name
		}

		if v, ok := fields[name]; ok {
			fieldValue.Set(v)
		}
	}
	return nil
}