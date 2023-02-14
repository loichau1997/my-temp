package common

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/sendgrid/rest"
	"math/rand"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

func RandStringBytes(n int, upper bool) string {

	var LETTER_BYTES = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
	if upper {
		LETTER_BYTES = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	}

	b := make([]byte, n)
	for i := range b {
		b[i] = LETTER_BYTES[rand.Intn(len(LETTER_BYTES))]
	}
	return string(b)
}

func SendRestAPI(url string, method rest.Method, header map[string]string, queryParam map[string]string, bodyInput interface{}) (body string, headers map[string][]string, err error) {
	request := rest.Request{
		Method:      method,
		BaseURL:     url,
		Headers:     header,
		QueryParams: queryParam,
	}
	if bodyInput != nil {
		bodyData, err := json.Marshal(bodyInput)
		if err != nil {
			return "", nil, err
		}
		request.Body = bodyData
	}
	response, err := rest.Send(request)
	if err != nil {
		return "", nil, err
	} else {
		if response.StatusCode != http.StatusOK &&
			response.StatusCode != http.StatusCreated &&
			response.StatusCode != http.StatusNoContent &&
			response.StatusCode != http.StatusAccepted {
			return response.Body,
				response.Headers,
				fmt.Errorf("Error when call Rest API, Status %s , Response Body %s", response.StatusCode, response.Body)
		} else {
			return response.Body, response.Headers, nil
		}
	}
}

// SliceContains return true if value exist in slice, otherwise false
func SliceContains(value, slice interface{}) bool {
	valueInterface := reflect.ValueOf(slice)
	if valueInterface.Kind() == reflect.Slice {
		for i := 0; i < valueInterface.Len(); i++ {
			if value == valueInterface.Index(i).Interface() {
				return true
			}
		}
	} else {
		return false
	}
	return false
}

func GenerateObjectKey(microservice, objectName, userKey string) (objectKey string, err error) {
	// generate nano id
	id, err := gonanoid.New()
	if err != nil {
		return "", err
	}
	// result
	result := append(make([]string, 0), microservice, objectName, id)
	if userKey != "" {
		// check valid userKey
		if strings.Contains(userKey, KeySeparator) {
			return "", fmt.Errorf(`user key cannot contain "%s" characters`, KeySeparator)
		}
		result = append(result, userKey)
	}
	objectKey = strings.Join(result, KeySeparator)
	return objectKey, nil
}

func GetObjectKey(str string) string {
	splitString := strings.Split(str, KeySeparator)
	return splitString[len(splitString)-1]
}

func RouterGroupWithObject(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	return routerGroup.Group(fmt.Sprintf("/:%v", GinParamObject))
}

func (t Time) Value() (driver.Value, error) {
	if !t.IsSet() {
		return "null", nil
	}
	return t.Time, nil
}
func (t *Time) IsSet() bool {
	return t.UnixNano() != (time.Time{}).UnixNano()
}
func Sync(from interface{}, to interface{}) interface{} {
	_from := reflect.ValueOf(from)
	_fromType := _from.Type()
	_to := reflect.ValueOf(to)

	for i := 0; i < _from.NumField(); i++ {
		fromName := _fromType.Field(i).Name
		field := _to.Elem().FieldByName(fromName)
		if !_from.Field(i).IsNil() && field.IsValid() && field.CanSet() {
			fromValue := _from.Field(i).Elem()
			fromType := reflect.TypeOf(fromValue.Interface())
			if fromType.String() == "uuid.UUID" {
				if fromValue.Interface() != uuid.Nil {
					field.Set(fromValue)
				}
			} else if fromType.String() == "string" {
				if field.Kind() == reflect.Ptr {
					tmp := fromValue.String()
					field.Set(reflect.ValueOf(&tmp))
				} else {
					field.Set(fromValue)
				}
			} else if fromType.String() == "service.Time" {
				tmp := fromValue.Interface().(Time)
				if tmp.IsSet() {
					if field.Kind() == reflect.Ptr {
						field.Set(reflect.ValueOf(&tmp))
					} else {
						field.Set(fromValue)
					}
				}
			} else {
				field.Set(fromValue)
			}
		}
	}
	return to
}
