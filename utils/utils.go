package utils

import (
	"encoding/json"
	"io"
	"strings"
)

const (
	lowerCaseLetters = "abcdefghijklmnopqrstuvwxyz"
	upperCaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers          = "0123456789"
	symbols          = " !\"\\#$%&'()*+,-./:;<=>?@[]^_`|~"
	//CurrentVersion  exprt the current application version (Used for Etags)
	CurrentVersion = "0.0.0"
)

// StringInterface Interface for map[string]
type StringInterface map[string]interface{}

// StringMap Redefine type map[string]string
type StringMap map[string]string

// StringArray Reddefine type []string
type StringArray []string

//StringArrayIntersection interesection between string arrays
func StringArrayIntersection(arr1, arr2 []string) []string {
	arrMap := map[string]bool{}
	result := []string{}

	for _, value := range arr1 {
		arrMap[value] = true
	}

	for _, value := range arr2 {
		if arrMap[value] {
			result = append(result, value)
		}
	}

	return result
}

// func fileExistsInConfigFolder(filename string) bool {
//   if len(filename) == 0 {
//     return false
//   }

//   if _, err := os.Stat(FindConfigFile(filename)); err == nil {
//     return true
//   }
//   return false
// }

//RemoveDuplicatesFromStringArray remove duplicate string from array ...
func RemoveDuplicatesFromStringArray(arr []string) []string {
	result := make([]string, 0, len(arr))
	seen := make(map[string]bool)

	for _, item := range arr {
		if !seen[item] {
			result = append(result, item)
			seen[item] = true
		}
	}

	return result
}

// MapToJSON converts a map to a json string
func MapToJSON(objmap map[string]string) string {
	b, err := json.Marshal(objmap)
	if err != nil {
		return ""
	}
	return string(b)
}

// MapFromJSON will decode the key/value pair map
func MapFromJSON(data io.Reader) map[string]string {
	decoder := json.NewDecoder(data)

	var objmap map[string]string
	if err := decoder.Decode(&objmap); err != nil {
		return make(map[string]string)
	}
	return objmap
}

// ArrayToJSON transfor an array into a json array
func ArrayToJSON(objmap []string) string {
	b, err := json.Marshal(objmap)
	if err != nil {
		return ""
	}
	return string(b)
}

// ArrayFromJSON Try to parse a json array into a go string array
func ArrayFromJSON(data io.Reader) []string {
	decoder := json.NewDecoder(data)

	var objmap []string
	if err := decoder.Decode(&objmap); err != nil {
		return make([]string, 0)
	}
	return objmap
}

// ArrayFromInterface transfor an interface into a Json object
func ArrayFromInterface(data interface{}) []string {
	stringArray := []string{}

	dataArray, ok := data.([]interface{})
	if !ok {
		return stringArray
	}

	for _, v := range dataArray {
		if str, ok := v.(string); ok {
			stringArray = append(stringArray, str)
		}
	}

	return stringArray
}

// StringInArray Search if provided string exist in provided array
func StringInArray(a string, array []string) bool {
	for _, b := range array {
		if b == a {
			return true
		}
	}
	return false
}

// StringInterfaceToJSON convert String interface into Json object
func StringInterfaceToJSON(objmap map[string]interface{}) string {
	b, err := json.Marshal(objmap)
	if err != nil {
		return ""
	}
	return string(b)
}

// StringInterfaceFromJSON Try to parse a json into map[string]interace{}
func StringInterfaceFromJSON(data io.Reader) map[string]interface{} {
	decoder := json.NewDecoder(data)

	var objmap map[string]interface{}
	if err := decoder.Decode(&objmap); err != nil {
		return make(map[string]interface{})
	}
	return objmap
}

// StringToJSON convert provided string into Json object
func StringToJSON(s string) string {
	b, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(b)
}

// StringFromJSON Convert providing json into string
func StringFromJSON(data io.Reader) string {
	decoder := json.NewDecoder(data)

	var s string
	if err := decoder.Decode(&s); err != nil {
		return ""
	}
	return s
}

// IsLower check if a string contain only lower cas characters
func IsLower(s string) bool {
	return strings.ToLower(s) == s
}
