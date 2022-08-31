package utils

import json "github.com/bytedance/sonic"

// TypeConverter
//
// Example :
// user := user.User {
// 		Name:     "John",
// 		IsActive: 1,
// }
//
// data, err := utils.TypeConverter[models.User](&user)
//
// log.Println(reflect.TypeOf(data)) // will output *models.User
func TypeConverter[R any](data any) (*R, error) {
	var result R
	b, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func ListConverter[R any](data any) ([]R, error) {
	var result []R
	b, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
