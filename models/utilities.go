package models

import (
	"bytes"
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	u "github.com/titouanfreville/popcubeapi/utils"

	"github.com/pborman/uuid"
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
// type StringInterface map[string]interface{}

// // StringMap Redefine type map[string]string
// type StringMap map[string]string

// // StringArray Reddefine type []string
// type StringArray []string

// //EncryptStringMap define type map[string]string for encryption usage
// type EncryptStringMap map[string]string

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")

// NewID is a globally unique identifier.  It is a [A-Z0-9] string 26
// characters long.  It is a UUID version 4 Guid that is zbased32 encoded
// with the padding stripped off.
func NewID() string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(uuid.NewRandom())
	encoder.Close()
	b.Truncate(26) // removes the '==' padding
	return b.String()
}

//NewRandomString Generate a randow string length by provided int.
func NewRandomString(length int) string {
	var b bytes.Buffer
	str := make([]byte, length+8)
	rand.Read(str)
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(str)
	encoder.Close()
	b.Truncate(length) // removes the '==' padding
	return b.String()
}

// GetMillis is a convience method to get milliseconds since epoch.
func GetMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// MapToJSON converts a map to a json string
// func MapToJSON(objmap map[string]string) string {
// 	b, err := json.Marshal(objmap)
// 	if err != nil {
// 		return ""
// 	}
// 	return string(b)
// }

// // MapFromJSON will decode the key/value pair map
// func MapFromJSON(data io.Reader) map[string]string {
// 	decoder := json.NewDecoder(data)

// 	var objmap map[string]string
// 	if err := decoder.Decode(&objmap); err != nil {
// 		return make(map[string]string)
// 	}
// 	return objmap
// }

// // ArrayToJSON transfor an array into a json array
// func ArrayToJSON(objmap []string) string {
// 	b, err := json.Marshal(objmap)
// 	if err != nil {
// 		return ""
// 	}
// 	return string(b)
// }

// // ArrayFromJSON Try to parse a json array into a go string array
// func ArrayFromJSON(data io.Reader) []string {
// 	decoder := json.NewDecoder(data)

// 	var objmap []string
// 	if err := decoder.Decode(&objmap); err != nil {
// 		return make([]string, 0)
// 	}
// 	return objmap
// }

// // ArrayFromInterface transfor an interface into a Json object
// func ArrayFromInterface(data interface{}) []string {
// 	stringArray := []string{}

// 	dataArray, ok := data.([]interface{})
// 	if !ok {
// 		return stringArray
// 	}

// 	for _, v := range dataArray {
// 		if str, ok := v.(string); ok {
// 			stringArray = append(stringArray, str)
// 		}
// 	}

// 	return stringArray
// }

// // StringInArray Search if provided string exist in provided array
// func StringInArray(a string, array []string) bool {
// 	for _, b := range array {
// 		if b == a {
// 			return true
// 		}
// 	}
// 	return false
// }

// // StringInterfaceToJSON convert String interface into Json object
// func StringInterfaceToJSON(objmap map[string]interface{}) string {
// 	b, err := json.Marshal(objmap)
// 	if err != nil {
// 		return ""
// 	}
// 	return string(b)
// }

// // StringInterfaceFromJSON Try to parse a json into map[string]interace{}
// func StringInterfaceFromJSON(data io.Reader) map[string]interface{} {
// 	decoder := json.NewDecoder(data)

// 	var objmap map[string]interface{}
// 	if err := decoder.Decode(&objmap); err != nil {
// 		return make(map[string]interface{})
// 	}
// 	return objmap
// }

// // StringToJSON convert provided string into Json object
// func StringToJSON(s string) string {
// 	b, err := json.Marshal(s)
// 	if err != nil {
// 		return ""
// 	}
// 	return string(b)
// }

// // StringFromJSON Convert providing json into string
// func StringFromJSON(data io.Reader) string {
// 	decoder := json.NewDecoder(data)

// 	var s string
// 	if err := decoder.Decode(&s); err != nil {
// 		return ""
// 	}
// 	return s
// }

// u.IsLower check if a string contain only lower cas characters
// func u.IsLower(s string) bool {
// 	return strings.ToLower(s) == s
// }

// IsValidEmail check email validity
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil && u.IsLower(email)
}

// IsValidDomain check if provided value is a possible domain name
func IsValidDomain(domain string) bool {
	return u.IsLower(domain) && IsValidAlphaNum(domain, true)
}

var reservedName = []string{
	"signup",
	"login",
	"admin",
	"channel",
	"post",
	"api",
	"oauth",
}

//IsValidIDChannelentifier check if string provided is a correct channel identifier
func IsValidIDChannelentifier(s string) bool {
	return IsValidAlphaNum(s, true)
}

//IsValidOrganisationIdentifier check if string provided is a correct organisation identifier
func IsValidOrganisationIdentifier(s string) bool {

	return IsValidAlphaNum(s, true)
}

var validAlphaNumUnderscore = regexp.MustCompile(`^[a-z0-9]+([a-z\-\_0-9]+|(__)?)[a-z0-9]+$`)
var validAlphaNum = regexp.MustCompile(`^[a-z0-9]+([a-z\-0-9]+|(__)?)[a-z0-9]+$`)

//IsValidAlphaNum Check that string is correct lower case alpha numeric chain
func IsValidAlphaNum(s string, allowUnderscores bool) bool {
	var match bool
	if allowUnderscores {
		match = validAlphaNumUnderscore.MatchString(s)
	} else {
		match = validAlphaNum.MatchString(s)
	}

	return match
}

// Etag function create a string used for cache and coockies storage
func Etag(parts ...interface{}) string {

	Etag := CurrentVersion

	for _, part := range parts {
		Etag += fmt.Sprintf(".%v", part)
	}

	return Etag
}

var validHashtag = regexp.MustCompile(`^(#\pL[\pL\d\-_.]*[\pL\d])$`)
var puncStart = regexp.MustCompile(`^[^\pL\d\s#]+`)
var hashtagStart = regexp.MustCompile(`^#{2,}`)
var puncEnd = regexp.MustCompile(`[^\pL\d\s]+$`)

// ParseHashtags parse #xxxxxxxxx declaration in messages
func ParseHashtags(text string) (string, string) {
	words := strings.Fields(text)

	hashtagString := ""
	plainString := ""
	for _, word := range words {
		// trim off surrounding punctuation
		word = puncStart.ReplaceAllString(word, "")
		word = puncEnd.ReplaceAllString(word, "")

		// and remove extra pound #s
		word = hashtagStart.ReplaceAllString(word, "#")

		if validHashtag.MatchString(word) {
			hashtagString += " " + word
		} else {
			plainString += " " + word
		}
	}

	if len(hashtagString) > 1000 {
		hashtagString = hashtagString[:999]
		lastSpace := strings.LastIndex(hashtagString, " ")
		if lastSpace > -1 {
			hashtagString = hashtagString[:lastSpace]
		} else {
			hashtagString = ""
		}
	}

	return strings.TrimSpace(hashtagString), strings.TrimSpace(plainString)
}

// func IsFileExtImage(ext string) bool {
// 	ext = strings.ToLower(ext)
// 	for _, imgExt := range IMAGE_EXTENSIONS {
// 		if ext == imgExt {
// 			return true
// 		}
// 	}
// 	return false
// }

// func GetImageMimeType(ext string) string {
// 	ext = strings.ToLower(ext)
// 	if len(IMAGE_MIME_TYPES[ext]) == 0 {
// 		return "image"
// 	} else {
// 		return IMAGE_MIME_TYPES[ext]
// 	}
// }

// ClearMentionTags remove mention tags from messages
func ClearMentionTags(post string) string {
	post = strings.Replace(post, "<mention>", "", -1)
	post = strings.Replace(post, "</mention>", "", -1)
	return post
}

// URLRegex is a small variable to expose regexp matching URL
var URLRegex = regexp.MustCompile(`^((?:[a-z]+:\/\/)?(?:(?:[a-z0-9\-]+\.)+(?:[a-z]{2}|aero|arpa|biz|com|coop|edu|gov|info|int|jobs|mil|museum|name|nato|net|org|pro|travel|local|internal))(:[0-9]{1,5})?(?:\/[a-z0-9_\-\.~]+)*(\/([a-z0-9_\-\.]*)(?:\?[a-z0-9+_~\-\.%=&amp;]*)?)?(?:#[a-zA-Z0-9!$&'()*+.=-_~:@/?]*)?)(?:\s+|$)$`)

// PartialURLRegex is a small variable to expose regexp matching URL (from parial parst)
var PartialURLRegex = regexp.MustCompile(`/([A-Za-z0-9]{26})/([A-Za-z0-9]{26})/((?:[A-Za-z0-9]{26})?.+(?:\.[A-Za-z0-9]{3,})?)`)

//SplitRunes Split runes define a table of runes saying if they are spliters or not
var SplitRunes = map[rune]bool{',': true, ' ': true, '.': true, '!': true, '?': true, ':': true, ';': true, '\n': true, '<': true, '>': true, '(': true, ')': true, '{': true, '}': true, '[': true, ']': true, '+': true, '/': true, '\\': true}

// IsValidHTTPURL check validity of provided string as http url
func IsValidHTTPURL(rawURL string) bool {
	if strings.Index(rawURL, "http://") != 0 && strings.Index(rawURL, "https://") != 0 {
		return false
	}

	if _, err := url.ParseRequestURI(rawURL); err != nil {
		return false
	}

	return true
}

// IsValidHTTPSURL check validity of provided string as https url
func IsValidHTTPSURL(rawURL string) bool {
	if strings.Index(rawURL, "https://") != 0 {
		return false
	}

	if _, err := url.ParseRequestURI(rawURL); err != nil {
		return false
	}

	return true
}

// func IsValidTurnOrStunServer(rawURI string) bool {
// 	if strings.Index(rawURI, "turn:") != 0 && strings.Index(rawURI, "stun:") != 0 {
// 		return false
// 	}

// 	if _, err := url.ParseRequestURI(rawURI); err != nil {
// 		return false
// 	}

// 	return true
// }

// IsSafeLink check if provided link can be considered as Safe
func IsSafeLink(link *string) bool {
	if link != nil {
		if IsValidHTTPURL(*link) {
			return true
		}
		if strings.HasPrefix(*link, "/") {
			return true
		}
		return false

	}

	return true
}

// IsValidWebsocketURL Check if provided string can be used as websocked url.
func IsValidWebsocketURL(rawURL string) bool {
	if strings.Index(rawURL, "ws://") != 0 && strings.Index(rawURL, "wss://") != 0 {
		return false
	}

	if _, err := url.ParseRequestURI(rawURL); err != nil {
		return false
	}

	return true
}
