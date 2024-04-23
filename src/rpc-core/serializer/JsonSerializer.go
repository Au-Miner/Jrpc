package serializer

import "encoding/json"

type JsonSerializer struct{}

func (j *JsonSerializer) Serialize(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func (j *JsonSerializer) Deserialize(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (j *JsonSerializer) GetCode() int {
	return JSONSerializer
}
