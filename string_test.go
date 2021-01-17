package conv_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/gopub/conv"
	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	type String string
	goodCases := []struct {
		Value  interface{}
		Result string
	}{
		{
			"This is a string",
			"This is a string",
		},
		{
			String("Typed string"),
			"Typed string",
		},
		{
			fmt.Errorf("error string"),
			"error string",
		},
		{
			10,
			"10",
		},
		{
			-10,
			"-10",
		},
		{
			10.2,
			"10.2",
		},
		{
			-10.2,
			"-10.2",
		},
		{
			[]byte("This is a byte slice"),
			"This is a byte slice",
		},
		{
			json.Number("10e6"),
			"10e6",
		},
	}

	t.Run("Good", func(t *testing.T) {
		for _, c := range goodCases {
			res, err := conv.ToString(c.Value)
			assert.NoError(t, err, c.Value)
			assert.Equal(t, c.Result, res)
		}
	})

	type Foo struct{}
	badCases := []interface{}{
		Foo{}, &Foo{},
	}
	t.Run("Bad", func(t *testing.T) {
		for _, c := range badCases {
			res, err := conv.ToString(c)
			assert.Error(t, err, c)
			assert.Equal(t, "", res)
		}
	})
}

func TestToStringSlice(t *testing.T) {
	t.Run("SingleString", func(t *testing.T) {
		s := "123"
		l, err := conv.ToStringSlice(s)
		require.NoError(t, err)
		require.Equal(t, []string{s}, l)
	})
	t.Run("SingleInt", func(t *testing.T) {
		s := 123
		l, err := conv.ToStringSlice(s)
		require.NoError(t, err)
		require.Equal(t, []string{fmt.Sprint(s)}, l)
	})
	t.Run("IntSlice", func(t *testing.T) {
		s := []int{123, -1, 9}
		l, err := conv.ToStringSlice(s)
		require.NoError(t, err)
		require.Equal(t, []string{"123", "-1", "9"}, l)
	})
	t.Run("MixSlice", func(t *testing.T) {
		s := []interface{}{123, "hello", "0x123"}
		l, err := conv.ToStringSlice(s)
		require.NoError(t, err)
		require.Equal(t, []string{"123", "hello", "0x123"}, l)
	})
	t.Run("MixArray", func(t *testing.T) {
		s := [3]interface{}{123, "hello", "0x123"}
		l, err := conv.ToStringSlice(s)
		require.NoError(t, err)
		require.Equal(t, []string{"123", "hello", "0x123"}, l)
	})
}
