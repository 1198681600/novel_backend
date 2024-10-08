package utils

import "encoding/json"

func ConvertPbTypeToModelTypes[F comparable, T comparable](fs []F) (ts []T, err error) {
	marshal, err := json.Marshal(fs)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(marshal, &ts)
	return
}

func ConvertPbTypeToModelType[F any, T any](fs F) (ts T, err error) {
	marshal, err := json.Marshal(fs)
	if err != nil {
		return ts, err
	}
	err = json.Unmarshal(marshal, &ts)
	return
}
