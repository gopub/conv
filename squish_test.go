package conv_test

import (
	"testing"
	"time"

	"github.com/gopub/conv"
	"github.com/stretchr/testify/require"
)

func TestSquishString(t *testing.T) {
	s := "\n\n\t\t1 \t2\n3  4\n\t \n5   "
	require.Equal(t, "1 2 3 4 5", conv.SquishString(s))
}

func TestSquishStringFields(t *testing.T) {
	type FullName struct {
		FirstName string
		LastName  string
	}

	type User struct {
		ID        int64
		Name      FullName
		Address   string
		CreatedAt time.Time
	}

	u := &User{
		ID: 1,
		Name: FullName{
			FirstName: "Tom  ",
			LastName:  "  Jim ",
		},
		Address: "\n\n \t Toronto   Canada   ",
	}
	conv.SquishStringFields(u)
	require.Equal(t, "Tom", u.Name.FirstName)
	require.Equal(t, "Jim", u.Name.LastName)
	require.Equal(t, "Toronto Canada", u.Address)
}
