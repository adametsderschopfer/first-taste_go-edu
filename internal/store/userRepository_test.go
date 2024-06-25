package store_test

import (
	"github.com/adametsderschopfer/http-rest-api_go-edu/internal/store"
	"github.com/adametsderschopfer/http-rest-api_go-edu/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "test@example.org",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
