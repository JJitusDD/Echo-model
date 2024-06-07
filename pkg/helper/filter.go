package helpers

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"echo-model/internal/domain/model/request"
)

// List operator allow in filter
var mapOperators = map[string]string{
	"_gte":     ">=",
	"_lte":     "<=",
	"_gt":      ">",
	"_lt":      "<",
	"_contain": "ilike",
	"_is_null": "is_nil",
	"_in":      "IN",
	"_ne":      "<>",
}

var mapType = map[string]string{
	"created_date":     "datetime",
	"approved_date":    "datetime",
	"last_update_time": "datetime",
	"report_date":      "datetime",
	"apply_to_date":    "datetime",
	"apply_from_date":  "datetime",
	"create_date":      "datetime",
	"expired_date":     "datetime",
	"start_date":       "datetime",
	"trans_date":       "datetime",
}

// custom JSON to map to complex conditions
var mapField = map[string]string{
	"text_box": "merchant_code_ilike_OR_merchant_name_ilike_OR_phone_no_ilike",
}

func ExtractFiltersNew(in interface{}) map[string]interface{} {
	filterDoc := make(map[string]interface{})
	val := reflect.ValueOf(in)
	for i := 0; i < val.Type().NumField(); i++ {
		f, op, val2 := MapFieldOps(val.Field(i), val.Type().Field(i))
		if f == "" && op == "" {
			continue
		}

		if strings.Contains(f, "_OR_") {
			s := strings.Split(f, "_OR_")
			var keyOr []string
			var val2Change string
			//"name1 = @name OR name2 = @name"
			for j := 0; j < len(s); j++ {
				var field string
				if strings.Contains(s[j], "_ilike") {
					field = strings.Split(s[j], "_ilike")[0] + " ilike @" + f
					val2Change = "_ilike"
				} else {
					field = s[j] + " = @" + f
				}
				keyOr = append(keyOr, field)
			}

			if val2Change == "_ilike" {
				val2 = "%" + fmt.Sprint(val2) + "%"
			}

			keyStr := strings.Join(keyOr, " OR ")
			keyStr = "(" + keyStr + ")"
			filterDoc[keyStr] = map[string]interface{}{f: val2}
			continue
		}

		if strings.Contains(f, "_INOR_") {
			s := strings.Split(f, "_INOR_")
			var keyOr []string
			//"name1 = @name OR name2 = @name"
			for j := 0; j < len(s); j++ {
				field := s[j] + " IN @" + f
				keyOr = append(keyOr, field)
			}
			keyStr := strings.Join(keyOr, " OR ")
			keyStr = "(" + keyStr + ")"
			filterDoc[keyStr] = map[string]interface{}{f: val2}
			continue
		}

		if ctype, ok := mapType[f]; ok {
			val2 = parseType(val2, ctype)
		}
		var key string

		if op == "is_nil" {
			if reflect.DeepEqual(val2, false) {
				key = f + " IS NOT NULL AND " + f + " <> ''"
			} else {
				key = f + " IS NULL OR " + f + " = '' "
			}
		} else if op == "IN" {
			key = f + " " + op + " (?)"
		} else {
			key = f + " " + op + " ?"
		}
		if op == "ilike" {
			if fmt.Sprint(val2) == "" {
				continue
			}
			filterDoc[key] = "%" + fmt.Sprint(val2) + "%"
		} else if op == "is_nil" {
			filterDoc[key] = nil
		} else {
			filterDoc[key] = val2
		}

	}
	return filterDoc
}

func ExtractFiltersWithPrefix(in interface{}, prefix string) map[string]interface{} {
	filterDoc := make(map[string]interface{})
	val := reflect.ValueOf(in)
	for i := 0; i < val.Type().NumField(); i++ {
		f, op, val2 := MapFieldOps(val.Field(i), val.Type().Field(i))
		if f == "" && op == "" {
			continue
		}

		startWith := ""
		if len(prefix) > 0 {
			startWith = prefix + "."
		}

		if strings.Contains(f, "_OR_") {
			s := strings.Split(f, "_OR_")
			var keyOr []string
			var val2Change string
			//"name1 = @name OR name2 = @name"
			for j := 0; j < len(s); j++ {
				var field string
				if strings.Contains(s[j], "_ilike") {
					field = startWith + strings.Split(s[j], "_ilike")[0] + " like @" + f
					val2Change = "_ilike"
				} else {
					field = startWith + s[j] + " = @" + f
				}
				keyOr = append(keyOr, field)
			}

			if val2Change == "_ilike" {
				val2 = "%" + fmt.Sprint(val2) + "%"
			}

			keyStr := strings.Join(keyOr, " OR ")
			keyStr = "(" + keyStr + ")"
			filterDoc[keyStr] = map[string]interface{}{f: val2}
			continue
		}

		if strings.Contains(f, "_INOR_") {
			s := strings.Split(f, "_INOR_")
			var keyOr []string
			//"name1 = @name OR name2 = @name"
			for j := 0; j < len(s); j++ {
				field := startWith + s[j] + " IN @" + f
				keyOr = append(keyOr, field)
			}
			keyStr := strings.Join(keyOr, " OR ")
			keyStr = "(" + keyStr + ")"
			filterDoc[keyStr] = map[string]interface{}{f: val2}
			continue
		}

		if ctype, ok := mapType[f]; ok {
			val2 = parseType(val2, ctype)
		}
		var key string

		if op == "is_nil" {
			if reflect.DeepEqual(val2, false) {
				key = startWith + f + " IS NOT NULL AND " + f + " <> ''"
			} else {
				key = startWith + f + " IS NULL OR " + f + " = '' "
			}
		} else if op == "IN" {
			key = startWith + f + " " + op + " (?)"
		} else {
			key = startWith + f + " " + op + " ?"
		}
		if op == "ilike" {
			if fmt.Sprint(val2) == "" {
				continue
			}
			filterDoc[key] = "%" + fmt.Sprint(val2) + "%"
		} else if op == "is_nil" {
			filterDoc[key] = nil
		} else {
			filterDoc[key] = val2
		}

	}
	return filterDoc
}

func parseType(val interface{}, ctype string) interface{} {
	switch ctype {
	case "datetime":
		return time.Unix(val.(int64), 0)
	default:
		return val
	}
}

func MapFieldOps(in reflect.Value, strField reflect.StructField) (string, string, interface{}) {
	if in.IsZero() {
		return "", "", nil
	}
	if in.Type().Name() != "" {
		return "", "", nil
	}

	fieldName := strField.Tag.Get("json")
	fieldName = strings.Replace(fieldName, ",omitempty", "", -1)
	if mapField[fieldName] != "" {
		fieldName = mapField[fieldName]
	}
	if fieldName == "from_date" {
		fieldName = "created_date_gte"
	}
	if fieldName == "to_date" {
		fieldName = "created_date_lte"
	}

	if fieldName == "start_date" {
		fieldName = "start_date_gte"
	}

	if fieldName == "expired_date" {
		fieldName = "expired_date_lte"
	}
	if fieldName == "pagination" {
		return "", "", nil
	}
	f, op := GetFilterOperator(fieldName)

	val := reflect.Indirect(in).Interface()

	if _, ok := mapType[f]; ok {
		if reflect.DeepEqual(val.(int64), int64(0)) {
			return "", "", nil
		}
	}

	if f == "" || reflect.DeepEqual(val, "") && op == "=" {
		return "", "", nil
	}

	return f, op, val
}

// GetFilterOperator ...
func GetFilterOperator(fieldName string) (string, string) {
	op := "="
	field := ""
	for k, v := range mapOperators {
		if strings.HasSuffix(fieldName, k) {
			op = v
			field = fieldName[:len(fieldName)-len(k)]
			return field, op
		}
	}
	return fieldName, op
}

func RootDirectory() string {
	rootDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting root directory:", err)
		return ""
	}
	fmt.Println("Root directory:", rootDirectory)

	return rootDirectory
}

func GetPagingRequest(info *request.Pagination) *request.Pagination {
	result := &request.Pagination{
		Page:    1,
		Order:   "",
		Offset:  0,
		PerPage: 10,
	}
	if info != nil {
		result.PerPage = info.PerPage
		result.Page = info.Page
		result.Offset = info.Offset
		if info.Page == 0 {
			result.Page = 1
		}

		if info.Page == 0 {
			result.Page = 1
		}

		if info.Order != "" {
			result.Order = info.Order
		}
		if info.PerPage <= 0 {
			result.PerPage = 10
		}
		if info.PerPage >= 1000 {
			result.PerPage = 1000
		}
	}

	return result
}
