package printindent

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

func JSON(data interface{}) error {
	encoded, err := jsoniter.ConfigCompatibleWithStandardLibrary.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	fmt.Println(string(encoded))

	return nil
}
