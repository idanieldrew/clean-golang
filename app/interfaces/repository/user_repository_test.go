package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	cache, cacheMock := redismock.NewClientMock()
	cacheMock.ExpectGet("users").RedisNil()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"Id", "Name", "Email", "Phone", "CreatedAt", "UpdatedAt",
	}).
		AddRow(1, "daniel", "dan@dan.com", "0912538201", time.Time{}, time.Time{}).
		AddRow(2, "sahar", "sahar@sahar.com", "09125598550", time.Time{}, time.Time{})

	mock.ExpectQuery(All).WillReturnRows(rows)
	cacheMock.ExpectSet("users", rows, 15*time.Minute)

	r := &UserRepository{
		Connection: db,
		Cache:      cache,
	}

	users, AllErr := r.All()

	assert.NoError(t, AllErr)
	assert.NotNil(t, users)
}
