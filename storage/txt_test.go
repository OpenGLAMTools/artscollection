package storage

import (
	"reflect"
	"testing"
)

var testFields = []Field{
	Field{
		Key:   "4",
		Name:  "Feld1Int",
		Type:  "int",
		Group: "",
		Order: 4,
	},
	Field{
		Key:   "2",
		Name:  "Feld2",
		Type:  "string",
		Group: "g1",
		Order: 2,
	},
	Field{
		Key:   "3",
		Name:  "Feld3",
		Type:  "string",
		Group: "g1",
		Order: 3,
	},
	Field{
		Key:   "1",
		Name:  "Feld1",
		Type:  "bool",
		Group: "",
		Order: 1,
	},
}

func TestSet(t *testing.T) {
	testCases := []struct {
		Key   string
		Input interface{}
		Error error
	}{
		{"3", int64(3), ErrTypeNotSupported},
		{"12345", "string value", ErrFieldNotSupported},
		{"4", "string value", ErrWrongType},
		{"4", 4, nil},
	}

	for _, tc := range testCases {
		txt := NewTxtStorage()
		txt.Fields = testFields
		err := txt.Set(tc.Key, tc.Input)
		if err != tc.Error {
			t.Errorf("Wrong Error!\nExp:%v\nGot:%v\n", tc.Error, err)
		}
	}
}

func TestGet(t *testing.T) {
	testCases := []struct {
		Key   string
		Input interface{}
	}{
		{"1", true},
		{"2", "String for field2"},
		{"4", 2704},
	}
	txt := NewTxtStorage()
	txt.Fields = testFields
	for _, tc := range testCases {
		err := txt.Set(tc.Key, tc.Input)
		if err != nil {
			t.Error(err)
		}
		got, ok := txt.Get(tc.Key)
		if !ok {
			t.Error("Get should return true for ok!")
		}
		if !reflect.DeepEqual(got, tc.Input) {
			t.Errorf("Error in Get method\nGot: %#v\nExp:%#v\n", got, tc.Input)
		}
	}
	val, ok := txt.Get("NotExist")
	if ok || val != nil {
		t.Error(
			"Get should return false for ok!\n",
			"Got a not expected value: ",
			val,
		)
	}
}

func TestFieldMethods(t *testing.T) {
	txt := NewTxtStorage()
	// Add the fields via AddField
	for _, f := range testFields {
		txt.AddField(f)
	}
	if !reflect.DeepEqual(txt.Fields, testFields) {
		t.Errorf("Error when adding fields!\nGot: %#v\nExp: %#v", txt.Fields, testFields)
	}
}

func TestMarshal(t *testing.T) {
	testCases := []struct {
		Key   string
		Input interface{}
	}{
		{"1", true},
		{"2", "String for field2"},
		{"4", 2704},
	}
	txt := NewTxtStorage()
	txt.Fields = testFields
	for _, tc := range testCases {
		err := txt.Set(tc.Key, tc.Input)
		if err != nil {
			t.Error(err)
		}
	}
	b, _ := txt.Marshal()

	txt2 := NewTxtStorage()
	txt2.Unmarshal(b)

	for _, tc := range testCases {
		got, ok := txt2.Get(tc.Key)
		if !ok {
			t.Error("Value not found!")
		}
		if got != tc.Input {
			t.Errorf("Got:%#v\nExp:%#v", got, tc.Input)
		}
	}
}
