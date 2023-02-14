package common

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"sort"
	"strings"
)

func ConvertLanguageSTJ(lstLanguage []string, valueStr *string) (interface{}, error) {
	if valueStr == nil || *valueStr == "" {
		return nil, nil
	}

	convertJson := make(map[string]interface{})
	strEnd := "[:]"
	for _, language := range lstLanguage {
		strStart := fmt.Sprintf("[:%s]", language)
		if !strings.Contains(*valueStr, strStart) {
			continue
		}
		rx := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(strStart) + `(.*?)` + regexp.QuoteMeta(strEnd))
		matches := rx.FindAllStringSubmatch(*valueStr, -1)

		if len(matches) > 0 && len(matches[0]) > 1 {
			convertJson[language] = matches[0][1]
		}
	}

	var jsonData interface{}
	bConvert, err := json.Marshal(convertJson)
	if err != nil {
		return "", err
	}

	if err = json.Unmarshal(bConvert, &jsonData); err != nil {
		return "", err
	}
	*valueStr = ""
	return jsonData, nil
}

func ConvertLanguageJTS(valueJson interface{}) (string, error) {

	value := ""
	var returnJson map[string]interface{}
	bValue, err := json.Marshal(valueJson)
	err = json.Unmarshal(bValue, &returnJson)
	if err != nil {
		return "", err
	}
	for v, ok := range returnJson {
		value += "[:" + v + "]" + ok.(string) + "[:]"
	}

	return value, nil
}

func ConvertInterfaceToSliceOfInterface(value interface{}) []interface{} {
	var result []interface{}
	valueInterface := reflect.ValueOf(value)
	if valueInterface.Kind() == reflect.Slice {
		for i := 0; i < valueInterface.Len(); i++ {
			result = append(result, valueInterface.Index(i).Interface())
		}
	}
	return result
}

func ConvertToMapInterface(value interface{}) (valueMap map[string]interface{}, err error) {

	var bValue []byte

	switch value.(type) {
	case string:
		bValue = []byte(value.(string))
	default:
		bValue, err = json.Marshal(value)
		if err != nil {
			return nil, err
		}
	}

	if err := json.Unmarshal(bValue, &valueMap); err != nil {
		return nil, err
	}

	return valueMap, nil
}

func ConvertToMapInterface2(value string) (map[string]interface{}, error) {

	var esBody map[string]interface{}
	b := []byte(value)

	if err := json.Unmarshal(b, &esBody); err != nil {
		return nil, err
	}

	var valueMap map[string]interface{}
	valueMap, _ = ConvertToMapInterface(esBody)

	return valueMap, nil
}

func ConvertToMapString(value interface{}) (map[string]string, error) {

	bValue, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	var valueMap map[string]string
	if err := json.Unmarshal(bValue, &valueMap); err != nil {
		return nil, err
	}

	return valueMap, nil
}

func ConvertMapInterfaceToString(mapValue map[string]interface{}) (map[string]string, error) {

	valueMap := make(map[string]string)
	for key, value := range mapValue {
		valueMap[key] = fmt.Sprintf("%v", value)
	}

	return valueMap, nil
}

func ConvertMapInterfaceToQueryString(m map[string]interface{}) (string, error) {
	var params []string
	for k, v := range m {
		params = append(params, fmt.Sprintf("%s=%v", k, v))
	}
	// sort slice alphabetically
	sort.Sort(sort.StringSlice(params))

	// convert slice to query string
	query := strings.Join(params, "&")

	// remove escape characters
	result, err := url.QueryUnescape(query)
	if err != nil {
		return "", err
	}

	return result, nil
}

func ConvertToArray(paramIn interface{}, mapIn string) (interface{}, error) {
	var arrayResult []interface{}

	valueGet, err := GetMappingValue(strings.TrimPrefix(mapIn, mappingArrayPrefix), paramIn)
	if err != nil {
		return nil, err
	}
	arrayResult = append(arrayResult, valueGet)
	return arrayResult, nil
}
