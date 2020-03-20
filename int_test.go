package conv_test

import (
	"math"
	"testing"
	"time"

	"github.com/gopub/conv"
	"github.com/stretchr/testify/assert"
)

func TestToInt64(t *testing.T) {
	type Bool bool
	type Int int
	goodCases := []struct {
		Value  interface{}
		Result int64
	}{
		{
			true,
			1,
		}, {
			false,
			0,
		}, {
			123456,
			123456,
		}, {
			Int(123456),
			123456,
		}, {
			int8(123),
			123,
		}, {
			int16(123),
			123,
		}, {
			int32(123),
			123,
		}, {
			int64(123),
			123,
		}, {
			uint8(123),
			123,
		}, {
			uint16(123),
			123,
		}, {
			uint32(123),
			123,
		}, {
			uint64(123),
			123,
		}, {
			math.MaxInt64,
			math.MaxInt64,
		}, {
			math.MinInt64,
			math.MinInt64,
		}, {
			0.0,
			0,
		},
		{
			-123.1,
			-123,
		},
		{
			Bool(false),
			0,
		},
	}

	t.Run("Good", func(t *testing.T) {
		for _, c := range goodCases {
			res, err := conv.ToInt64(c.Value)
			assert.NoError(t, err, c.Value)
			assert.Equal(t, c.Result, res)
		}
	})

	type Foo struct{}
	badCases := []interface{}{
		"hello", time.Now(), Foo{}, uint64(math.MaxInt64 + 1), uint64(math.MaxUint64),
	}
	t.Run("Bad", func(t *testing.T) {
		for _, c := range badCases {
			res, err := conv.ToInt64(c)
			assert.Error(t, err, c)
			assert.Empty(t, res)
		}
	})
}

func TestToInt(t *testing.T) {
	type Bool bool
	type Int int
	goodCases := []struct {
		Value  interface{}
		Result int
	}{
		{
			true,
			1,
		}, {
			false,
			0,
		}, {
			123456,
			123456,
		}, {
			Int(123456),
			123456,
		}, {
			int8(123),
			123,
		}, {
			int16(123),
			123,
		}, {
			int32(123),
			123,
		}, {
			int64(123),
			123,
		}, {
			uint8(123),
			123,
		}, {
			uint16(123),
			123,
		}, {
			uint32(123),
			123,
		}, {
			uint64(123),
			123,
		}, {
			0.0,
			0,
		},
		{
			-123.1,
			-123,
		},
		{
			Bool(false),
			0,
		},
	}

	t.Run("Good", func(t *testing.T) {
		for _, c := range goodCases {
			res, err := conv.ToInt(c.Value)
			assert.NoError(t, err, c.Value)
			assert.Equal(t, c.Result, res)
		}
	})

	type Foo struct{}
	badCases := []interface{}{
		"hello", time.Now(), Foo{}, uint64(math.MaxInt64 + 1), uint64(math.MaxUint64),
	}
	t.Run("Bad", func(t *testing.T) {
		for _, c := range badCases {
			res, err := conv.ToInt(c)
			assert.Error(t, err, c)
			assert.Empty(t, res)
		}
	})
}

func TestToUint64(t *testing.T) {
	type Bool bool
	type Int int
	goodCases := []struct {
		Value  interface{}
		Result uint64
	}{
		{
			true,
			1,
		}, {
			false,
			0,
		}, {
			123456,
			123456,
		}, {
			Int(123456),
			123456,
		}, {
			int8(123),
			123,
		}, {
			int16(123),
			123,
		}, {
			int32(123),
			123,
		}, {
			int64(123),
			123,
		}, {
			uint8(123),
			123,
		}, {
			uint16(123),
			123,
		}, {
			uint32(123),
			123,
		}, {
			uint64(123),
			123,
		}, {
			math.MaxInt64,
			math.MaxInt64,
		}, {
			0.0,
			0,
		},
		{
			Bool(false),
			0,
		},
	}

	t.Run("Good", func(t *testing.T) {
		for _, c := range goodCases {
			res, err := conv.ToUint64(c.Value)
			assert.NoError(t, err, c.Value)
			assert.Equal(t, c.Result, res)
		}
	})

	type Foo struct{}
	badCases := []interface{}{
		"hello", time.Now(), Foo{}, "18446744073709551616", -1,
	}
	t.Run("Bad", func(t *testing.T) {
		for _, c := range badCases {
			res, err := conv.ToUint64(c)
			assert.Error(t, err, c)
			assert.Empty(t, res)
		}
	})
}
