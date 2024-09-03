package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Victor Coelho", "v@v.com", "4343554")
	assert.Nil(t, err)
	assert.NotEmpty(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Victor Coelho", user.Name)
	assert.Equal(t, "v@v.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Victor Coelho", "v@v.com", "4343554")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("4343554"))
	assert.False(t, user.ValidatePassword("4343554534"))
	assert.NotEqual(t, "4343554", user.Password)
}
