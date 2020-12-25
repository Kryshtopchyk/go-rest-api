package store_test

import (
	"github.com/Kryshtopchyk/go-rest-api/internal/app/model"
	"github.com/Kryshtopchyk/go-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databeseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email:             "user@test.com",
		EncryptedPassword: "1234qwer",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databeseURL)
	defer teardown("users")

	email := "user@test.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email:             "user@test.com",
		EncryptedPassword: "1234qwer",
	})

	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
