package storage

import (
	"fmt"
	"testing"
)

type mock struct {
}

func (m *mock) Get(fieldID string) (interface{}, error) {
	switch fieldID {
	case "string":
		return "My String", nil
	case "int":
		return 123, nil
	}
	return nil, nil
}

func TestGetterInterface(t *testing.T) {
	mockStorage := &mock{}
	testCases := []struct {
		input  string
		expect string
	}{
		{
			input:  "string",
			expect: "string",
		},
		{
			input:  "int",
			expect: "int",
		},
	}
	for _, tcase := range testCases {
		got, _ := mockStorage.Get(tcase.input)
		gotType := fmt.Sprintf("%T", got)
		if gotType != tcase.expect {
			t.Errorf("Expect: %s; Got: %s", tcase.expect, gotType)
		}
		fmt.Printf("%T\n", got)
	}

	i, _ := mockStorage.Get("int")
	fmt.Println(i)
}
