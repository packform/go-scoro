package scoro

import (
	"encoding/json"
	"strings"
	"time"
)

const TimePattern = `"2006-01-02 15:04:05"`

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

// Localized strings, or single string in requested language

type Strings struct {
	Values map[string]string `json:",inline"`
}

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
