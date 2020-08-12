package domain

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "No user should be found with id equals zero")
	assert.NotNil(t, err, "Error expected with id equals zero")
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "User 0 not found", err.Msg)

	//	if user != nil {
	//		t.Error("No user should be found with id equals zero")
	//	}
	//	if err == nil {
	//		t.Error("Error expected with id equals zero")
	//	}
	//
	//	if err.Status != http.StatusNotFound {
	//		t.Error("404 Not found expected")
	//	}
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err, "No error should be returned")
	assert.EqualValues(t, user.Id, 123, fmt.Sprintf("Wrong user id: %v, shoud be 123", user.Id))
	assert.EqualValues(t, user.FName, "Zion", fmt.Sprintf("Wrong FName: %v, shoud be Zion", user.FName))
	assert.EqualValues(t, user.LName, "Phelps", fmt.Sprintf("Wrong LName: %v, shoud be Phelps", user.LName))
	assert.EqualValues(t, user.Email, "zion@phelps.com", fmt.Sprintf("Wrong LName: %v, shoud be zion@phelps.com", user.Email))
}
