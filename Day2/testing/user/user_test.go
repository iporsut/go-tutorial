package user

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserName(t *testing.T) {
	u := User{
		FirstName: "Weerasak",
		LastName:  "Chongnguluam",
	}

	require.Equal(t, "Weerasak Chongnguluam", u.Name())
}
