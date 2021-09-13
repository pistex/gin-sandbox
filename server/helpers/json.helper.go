package helpers

import "encoding/json"

func JsonMapper(src interface{}, dest interface{}) error {
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dest)
}
