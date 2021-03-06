package conv_test

import (
	"testing"

	"github.com/gopub/conv"
	"github.com/stretchr/testify/assert"
)

func TestSnakeToCamel(t *testing.T) {
	tests := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "hello",
			Output: "hello",
		},
		{
			Input:  "hello_world",
			Output: "helloWorld",
		},
		{
			Input:  "hello_world_",
			Output: "helloWorld",
		},
		{
			Input:  "hello_world_id",
			Output: "helloWorldId",
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.Output, conv.ToCamel(tc.Input))
	}
}

func TestCamelToSnake(t *testing.T) {
	tests := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "hello",
			Output: "hello",
		},
		{
			Input:  "helloWorld",
			Output: "hello_world",
		},
		{
			Input:  "helloWorldID",
			Output: "hello_world_id",
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.Output, conv.ToSnake(tc.Input))
	}
}
