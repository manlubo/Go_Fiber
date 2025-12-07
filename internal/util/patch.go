package util

import (
	"reflect"
	"time"
)

// BuildPatch: struct pointer를 받아서
// nil 또는 ""이 아닌 필드만 json태그 기준으로 map 생성
func BuildPatch(input interface{}, skipFields ...string) map[string]interface{} {
	patch := make(map[string]interface{})

	v := reflect.ValueOf(input).Elem()
	t := reflect.TypeOf(input).Elem()

	// skipFields 를 JSON 태그 기준으로 Set으로 변환
	skip := map[string]bool{}
	for _, s := range skipFields {
		skip[s] = true
	}

	for i := 0; i < v.NumField(); i++ {
		fieldVal := v.Field(i)
		fieldType := t.Field(i)

		// JSON 태그 가져오기
		jsonTag := fieldType.Tag.Get("json")

		// json:"-" 이거나 태그가 비어있으면 skip
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		// json tag 기준 skip
		if skip[jsonTag] {
			continue
		}

		// 포인터 필드만 처리 (*string 등)
		if fieldVal.Kind() != reflect.Ptr {
			continue
		}

		if fieldVal.IsNil() {
			continue
		}

		if fieldVal.Elem().Kind() == reflect.String && fieldVal.Elem().String() == "" {
			continue
		}

		patch[jsonTag] = fieldVal.Elem().Interface()
	}

	// 업데이트 시간 자동 반영
	now := time.Now().Unix()
	patch["updatedAt"] = now

	return patch
}
