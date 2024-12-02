package userproxy

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"misis/pkg/laba4/user"
)

// MockDatabase is a mock implementation of database.Database
type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) Connect() {
	m.Called()
}

func (m *MockDatabase) GetUser(id int) *user.User {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*user.User)
}

// MockLogger is a mock implementation of Mlogger.Mlogger
type MockMlogger struct {
	mock.Mock
}

func (m *MockMlogger) LevelDebug() {
	m.Called()
}

func (m *MockMlogger) LevelInfo() {
	m.Called()
}

func (m *MockMlogger) LogInfo(str string) string {
	args := m.Called(str)
	return args.String(0)
}

func (m *MockMlogger) LogError(str string) string {
	args := m.Called(str)
	return args.String(0)
}

func TestUserProxy_GetData(t *testing.T) {
	// Create mock objects
	mockLogger := new(MockMlogger)
	mockDatabase := new(MockDatabase)

	// Create UserProxy instance
	userProxy := NewUserProxy(mockLogger, mockDatabase)

	// Test data
	userID := 1
	userData := user.User{Name: "John Doe"}

	// Set expectations for mockDatabase
	mockDatabase.On("GetUser", userID).Return(&userData)

	// Set expectations for mockLogger
	mockLogger.On("LogInfo", "Данные взяты из базы").Return("Logged: Данные взяты из базы")
	mockLogger.On("LogInfo", "Данные взяты из кэша").Return("Logged: Данные взяты из кэша")

	// Test getting data from database
	userFromDB := userProxy.GetData(userID)
	assert.Equal(t, &userData, userFromDB, "Expected user data from database")
	mockDatabase.AssertExpectations(t)
	// Test getting data from cache
	userFromCache := userProxy.GetData(userID)
	assert.Equal(t, &userData, userFromCache, "Expected user data from cache")

	// Assert that the expectations were met
	mockDatabase.AssertExpectations(t)
	mockLogger.AssertExpectations(t)
}
