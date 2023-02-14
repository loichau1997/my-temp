package common

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func MappingApiParam(paramIn interface{}, paramMapping interface{}, mappingType string) (map[string]interface{}, error) {
	mapParamMapping, err := ConvertToMapString(paramMapping)
	if err != nil {
		return nil, fmt.Errorf("Error when ConvertToMapString %s", err.Error())
	}

	if len(mapParamMapping) == 0 {
		return nil, nil
	}

	countMapping := 0
	mapParamOut, err := MappingLogic(paramIn, mapParamMapping, mappingType, &countMapping)
	if err != nil {
		return nil, err
	} else if countMapping == 0 {
		return nil, fmt.Errorf("No matching fields exist")
	}

	return mapParamOut, nil
}

func MappingLogic(paramIn interface{}, mapParamMapping map[string]string, mappingType string, countMapping *int) (map[string]interface{}, error) {
	mapParamOut := make(map[string]interface{})
	for key, value := range mapParamMapping {
		var mapIn, mapOut string
		switch mappingType {
		case MAPPING_IN:
			mapIn = key
			mapOut = value
		case MAPPING_OUT:
			mapIn = value
			mapOut = key
		}

		if strings.HasPrefix(mapIn, mappingJsonPrefix) {
			keyValue, err := MappingArrayJson(paramIn, mapIn, mappingType)
			if err != nil {
				return nil, err
			}
			mapParamOut, err = SetMappingValue(mapOut, keyValue, mapParamOut, countMapping)
			if err != nil {
				return nil, err
			}
			continue
		} else if strings.HasPrefix(mapIn, mappingArrayPrefix) {
			keyValue, err := ConvertToArray(paramIn, mapIn)
			if err != nil {
				return nil, err
			}
			mapParamOut, err = SetMappingValue(mapOut, keyValue, mapParamOut, countMapping)
			if err != nil {
				return nil, err
			}
			continue
		}

		valueGet, err := GetMappingValue(mapIn, paramIn)
		if err != nil {
			return nil, err
		}
		if valueGet != nil {
			mapParamOut, err = SetMappingValue(mapOut, valueGet, mapParamOut, countMapping)
			if err != nil {
				return nil, err
			}
		} else {
			continue
		}
	}
	return mapParamOut, nil
}

func MappingArrayJson(paramIn interface{}, mappingKey string, mappingType string) (interface{}, error) {
	var sliceOfMapInterface []map[string]interface{}
	// trim json string
	str := strings.TrimPrefix(mappingKey, mappingJsonPrefix)
	r, err := regexp.Compile("{.*}$")
	if err != nil {
		return nil, err
	}
	// convert json string to map
	var mapParamMapping map[string]string
	err = json.Unmarshal([]byte(r.FindString(str)), &mapParamMapping)
	if err != nil {
		return nil, err
	}

	paramInMap, err := ConvertToMapInterface(paramIn)
	if err != nil {
		return nil, err
	}

	// convert interface to slice of interface
	sliceParamIn := ConvertInterfaceToSliceOfInterface(paramInMap[strings.TrimSuffix(str, r.FindString(str))])
	countMappingArray := 0

	// for each json data value to map
	for _, paramIn := range sliceParamIn {
		mapParamOut, err := MappingLogic(paramIn, mapParamMapping, mappingType, &countMappingArray)
		if err != nil {
			return nil, err
		}
		// FIXME hardcoded
		mapParamOut["weight"] = 1 + countMappingArray
		sliceOfMapInterface = append(sliceOfMapInterface, mapParamOut)
	}

	if countMappingArray == 0 {
		return nil, fmt.Errorf("No matching fields exist")
	}

	return sliceOfMapInterface, nil
}

func GetMappingValue(fieldMap string, data interface{}) (interface{}, error) {
	splitMap := strings.Split(fieldMap, "->")
	var dataMap map[string]interface{}
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &dataMap)
	if err != nil {
		return nil, err
	}

	var returnValue interface{}
	for i := 0; i < len(splitMap); i++ {
		if value, ok := dataMap[splitMap[i]]; ok {
			if i != len(splitMap)-1 {
				var childValue map[string]interface{}
				bChild, err := json.Marshal(value)
				if err != nil {
					return nil, err
				}
				err = json.Unmarshal(bChild, &childValue)
				if err != nil {
					return nil, err
				}
				dataMap = childValue
			} else {
				returnValue = value
			}
		}
	}

	return returnValue, nil
}

func SetMappingValue(fieldMap string, data interface{}, mapSetValue map[string]interface{}, countMapping *int) (map[string]interface{}, error) {

	fieldMapSplit := strings.Split(fieldMap, "->")
	lenFieldMap := len(fieldMapSplit) - 1
	childMap := make(map[string]interface{})

	for i := lenFieldMap; i >= 0; i-- {
		if i == lenFieldMap {
			childMap[fieldMapSplit[i]] = data
		} else if i != 0 && i != lenFieldMap {
			childMap[fieldMapSplit[i]] = childMap
		}

		if mapVal, ok := mapSetValue[fieldMapSplit[i]]; ok {
			bMapVal, err := json.Marshal(mapVal)
			if err != nil {
				return nil, err
			}
			var currentVal map[string]interface{}
			if err := json.Unmarshal(bMapVal, &currentVal); err != nil {
				return nil, err
			}
			appendMap := make(map[string]map[string]interface{})
			for key, val := range currentVal {
				tempCurrent := make(map[string]interface{})
				tempCurrent[key] = val
				appendMap[fieldMapSplit[i]] = tempCurrent
			}

			for key, val := range childMap {
				appendMap[fieldMapSplit[i]][key] = val
			}

			childMap[fieldMapSplit[i]] = appendMap[fieldMapSplit[i]]
			mapSetValue[fieldMapSplit[i]] = childMap[fieldMapSplit[i]]
			continue
		}

		if i == 0 && lenFieldMap == 0 {
			mapSetValue[fieldMapSplit[i]] = data
		} else if i == 0 {
			mapSetValue[fieldMapSplit[i]] = childMap
		}
	}
	*countMapping++
	return mapSetValue, nil
}
