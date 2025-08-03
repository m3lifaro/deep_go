package homework9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerialization(t *testing.T) {
	tests := map[string]struct {
		data   Person
		result string
	}{
		"test case with empty fields": {
			result: "name=\nage=0\nmarried=false",
		},
		"test case with fields": {
			data: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
			},
			result: "name=John Doe\nage=30\nmarried=true",
		},
		"test case with omitempty field": {
			data: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
				Address: "Paris",
			},
			result: "name=John Doe\naddress=Paris\nage=30\nmarried=true",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Serialize(test.data)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestSerialization2(t *testing.T) {
	var defP = Person{
		Name:    "John Doe",
		Age:     30,
		Married: true,
	}
	tests := map[string]struct {
		data   PersonEnh
		result string
	}{
		"test case with empty fields": {
			result: "person={name=\nage=0\nmarried=false}\nfamily=false",
		},
		"test case with fields": {
			data:   PersonEnh{Person: defP, HasFamily: false},
			result: "person={name=John Doe\nage=30\nmarried=true}\nfamily=false",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Serialize(test.data)
			assert.Equal(t, test.result, result)
		})
	}
}
