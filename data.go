package sensitive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

// secretValue is the masked value used by lib
var secretValue = func() string {
	secret := os.Getenv("SET_VALUE")
	if secret == "" {
		return "dummy-value"
	}
	return secret
}

// Mask used to hide/mask sensitive fields in input. This currently
// supports only `struct` and `string` fields when an `int` field is provided
// for masking an error is emitted.When no fields are provided for masking
// it emits an error with informative message
func Mask(d any, fields ...string) (string, error) {
	if len(fields) < 1 {
		fmt.Printf("%s\n", "required atleast one field")
	}
	inpBytes, err := json.Marshal(&d)
	if err != nil {
		return "", err
	}
	// extract list of fields
	inputType := reflect.TypeOf(d)
	inpFields := reflect.VisibleFields(inputType)
	typ := reflect.StructOf(inpFields)
	v := reflect.New(typ).Elem()
	if err := json.Unmarshal(inpBytes, v.Addr().Interface()); err != nil {
		return "", err
	}
	// replace sensitive fields
	for _, field := range fields {
		switch {
		case v.FieldByName(field).Kind().String() == "string":
			v.FieldByName(field).SetString(secretValue())
		default:
			return "",
				fmt.Errorf("invalid column data type,supports only string data type")
		}
	}

	s := v.Addr().Interface()
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(s); err != nil {
		return "", err
	}
	return b.String(), nil

}
