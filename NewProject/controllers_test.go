package controllers_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"NewProject/controllers"
	"NewProject/models"
)

type MockDB struct{}

func (m *MockDB) Query(query string, args ...interface{}) (models.Rows, error) {
	return nil, nil
}

func (m *MockDB) Exec(query string, args ...interface{}) (models.Result, error) {
	return nil, nil
}

type MockRows struct{}

func (m *MockRows) Next() bool {
	return false
}

func (m *MockRows) Scan(dest ...interface{}) error {
	return nil
}

func (m *MockRows) Close() error {
	return nil
}

func TestGetUser(t *testing.T) {
	userController := &controllers.UserController{}

	req, err := http.NewRequest("GET", "/user", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userController.GetUser)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	expected := `{"ID":1,"Name":"John","Email":"john@example.com"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestCreateUser_Positive(t *testing.T) {
	userController := &controllers.UserController{}
	jsonStr := []byte(`{"ID":2,"Name":"Alice","Email":"alice@example.com"}`)
	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonStr))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userController.CreateUser)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), `"ID":2`)
	assert.Contains(t, rr.Body.String(), `"Name":"Alice"`)
	assert.Contains(t, rr.Body.String(), `"Email":"alice@example.com"`)
}

func TestCreateUser_Negative_BadRequest(t *testing.T) {
	userController := &controllers.UserController{}
	jsonStr := []byte(`{"ID":3}`)
	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonStr))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userController.CreateUser)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetAllUsersFromDB_Positive(t *testing.T) {
	mockDB := &MockDB{}
	userController := &controllers.UserController{DB: mockDB}
	expectedUsers := []models.User{
		{ID: 1, Name: "John", Email: "john@example.com"},
		{ID: 2, Name: "Alice", Email: "alice@example.com"},
	}

	userController.DB = mockDB
	userController.GetAllUsersFromDBFunc = func() ([]models.User, error) {
		return expectedUsers, nil
	}
	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userController.GetAllUsersHandler)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	for _, user := range expectedUsers {
		assert.Contains(t, rr.Body.String(), user.Name)
		assert.Contains(t, rr.Body.String(), user.Email)
	}
}

func TestGetAllUsersFromDB_Negative_Error(t *testing.T) {
	mockDB := &MockDB{}
	userController := &controllers.UserController{DB: mockDB}
	expectedErr := errors.New("error fetching users")
	userController.DB = mockDB
	userController.GetAllUsersFromDBFunc = func() ([]models.User, error) {
		return nil, expectedErr
	}

	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userController.GetAllUsersHandler)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
