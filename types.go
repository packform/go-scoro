package scoro

// Implementation for marshaling and unmarshaling of common data types used in the API

import (
	"encoding/json"
	"strings"
	"time"
)

// TimePattern represents time format used in Scoro API requests and responses
const TimePattern = `"2006-01-02 15:04:05"`

// Time type provides the implementation of JSON time serialization into Scoro API format.
//
// Notes
//
// 	- format is YYYY-MM-DD hh:mm:ss
// 	- null value is supported
// 	- "0000-00-00 00:00:00" is considered as null
type Time struct {
	time.Time `json:",inline"`
}

func (t Time) MarshalJSON() ([]byte, error) {
	timeStr := t.Time.Format(TimePattern)
	return []byte(timeStr), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	timeStr := string(data)

	// Ignore null, like in the main JSON package.
	if timeStr == "null" {
		return nil
	}

	// We consider "0000-00-00 00:00:00" as nil.
	if timeStr == `"0000-00-00 00:00:00"` {
		return nil
	}

	var err error
	t.Time, err = time.Parse(TimePattern, timeStr)
	return err
}

// Bool type provides the implementation of serialization boolean values into Scoro API format.
// True -> "1" and False -> "0" mappings are used to serialize boolean values to into request JSON.
// Both boolean and "0"/"1" values are supported in response.
type Bool struct {
	Value bool `json:",inlline"`
}

func (t Bool) MarshalJSON() ([]byte, error) {
	if t.Value {
		return []byte(`"1"`), nil
	}

	return []byte(`"0"`), nil
}

func (t *Bool) UnmarshalJSON(data []byte) error {
	boolStr := string(data)

	// Ignore null, like in the main JSON package.
	if boolStr == "null" {
		return nil
	}

	var boolValue bool
	if err := json.Unmarshal(data, &boolValue); err == nil {
		t.Value = boolValue
		return nil
	}

	intStr := strings.Replace(boolStr, `"`, "", -1)
	t.Value = intStr != "0"
	return nil
}

// Strings provides support for string fields with localization support. Scoro
// expects strings as localied dictionaries for some fields in request. However,
// it can return them as localized dictionary as single string (in requested lang) for
// the same fields in response. Marshal/Unmarshal implementations for this
// type handle both cases appropriately.
//
// Examples:
//
// 		field := scoro.MakeStrings("Some string", scoro.DefaultLang)
// 		field := scoro.MakeStrings("Привет", "rus")
type Strings struct {
	Values map[string]string `json:",inline"`
}

// MakeStrings is helper method that creates strings for single language, it
// can be convinient if you don't need localization for multiple languages.
func MakeStrings(str string, lang string) Strings {
	values := make(map[string]string)
	values[lang] = str

	return Strings{values}
}

func (t Strings) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Values)
}

func (t *Strings) UnmarshalJSON(data []byte) error {
	str := string(data)

	t.Values = make(map[string]string)

	// Handle single string
	if str[0] == '"' {
		var defString string
		err := json.Unmarshal(data, &defString)

		if err == nil {
			t.Values["eng"] = defString
		}
		return err
	}

	return json.Unmarshal(data, t.Values)
}
