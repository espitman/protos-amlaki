package province

import "encoding/json"

func (t *ResponseDetails) Unmarshal(entry interface{}) ResponseDetails {
	obj, _ := json.Marshal(entry)
	response := ResponseDetails{}
	_ = json.Unmarshal(obj, &response)
	return response
}

func (t *RequestDetails) Unmarshal(entry interface{}) RequestDetails {
	obj, _ := json.Marshal(entry)
	response := RequestDetails{}
	_ = json.Unmarshal(obj, &response)
	return response
}

func (t *RequestCreate) Unmarshal(entry interface{}) RequestCreate {
	obj, _ := json.Marshal(entry)
	response := RequestCreate{}
	_ = json.Unmarshal(obj, &response)
	return response
}

func (t *RequestCheckDuplicatedNameAndTitle) Unmarshal(entry interface{}) RequestCheckDuplicatedNameAndTitle {
	obj, _ := json.Marshal(entry)
	response := RequestCheckDuplicatedNameAndTitle{}
	_ = json.Unmarshal(obj, &response)
	return response
}

func (t *ResponseList) Unmarshal(entry interface{}) ResponseList {
	obj, _ := json.Marshal(entry)
	response := ResponseList{}
	_ = json.Unmarshal(obj, &response)
	return response
}
