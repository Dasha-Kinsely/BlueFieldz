package processes

import (
	"encoding/json"
	//"log"
)

// Usage: an example to create []byte that is to be encoded into json is: fetchCondition := []byte("{\""+option+"\":\""+identifier+"\"}")
func CustomTypeStruct(inputs []byte) *map[string]interface{} {
	//log.Println(inputs, string(inputs))
	outputs := make(map[string]interface{})
	_ = json.Unmarshal(inputs, &outputs)
	//log.Println(e)
	return &outputs
}