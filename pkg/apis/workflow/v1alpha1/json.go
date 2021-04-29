package v1alpha1

import "encoding/json"

func MustJSON(v interface{}) string {
	x, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(x)
}
