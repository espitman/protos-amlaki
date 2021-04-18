package neighborhood

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

func (t *RequestCheckExistByProvinceIdAndCityId) Unmarshal(entry interface{}) RequestCheckExistByProvinceIdAndCityId {
	obj, _ := json.Marshal(entry)
	response := RequestCheckExistByProvinceIdAndCityId{}
	_ = json.Unmarshal(obj, &response)
	return response
}

func (t *RequestCheckValidNeighborhood) Unmarshal(entry interface{}) RequestCheckValidNeighborhood {
	obj, _ := json.Marshal(entry)
	response := RequestCheckValidNeighborhood{}
	_ = json.Unmarshal(obj, &response)
	return response
}

func (t *ResponseList) Unmarshal(entry interface{}) ResponseList {
	obj, _ := json.Marshal(entry)
	response := ResponseList{}
	_ = json.Unmarshal(obj, &response)
	return response
}
