package conv_test

import (
	"github.com/gopub/conv"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func TestSquishString(t *testing.T) {
	s := "\n\n\t\t1 \t2\n3  4\n\t \n5   "
	require.Equal(t, "1 2 3 4 5", conv.SquishString(s))
}

func TestSquishStringFields(t *testing.T) {
	var foo struct {
		ID        int64
		Name      string
		CreatedAt time.Time
	}

	foo.ID = rand.Int63()
	foo.Name = "  hello  "
	foo.CreatedAt = time.Now()

	conv.SquishStringFields(&foo)
	require.Equal(t, "hello", foo.Name)
}
