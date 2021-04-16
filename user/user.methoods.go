package user

import "encoding/json"

func (t *ResponseDetails) Unmarshal(entry interface{}) ResponseDetails {
	obj, _ := json.Marshal(entry)
	response := ResponseDetails{}
	_ = json.Unmarshal(obj, &response)
	return response
}

func (t *RequestRegisger) Unmarshal(entry interface{}) RequestRegisger {
	obj, _ := json.Marshal(entry)
	response := RequestRegisger{}
	_ = json.Unmarshal(obj, &response)
	return response
}

func (t *RequestLogin) Unmarshal(entry interface{}) RequestLogin {
	obj, _ := json.Marshal(entry)
	response := RequestLogin{}
	_ = json.Unmarshal(obj, &response)
	return response
}
