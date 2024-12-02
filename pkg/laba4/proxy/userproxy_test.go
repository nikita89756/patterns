package userproxy

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"misis/pkg/laba4/user"
)


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
	
	mockLogger := new(MockMlogger)
	mockDatabase := new(MockDatabase)

	
	userProxy := NewUserProxy(mockLogger, mockDatabase)

	
	userID := 1
	userData := user.User{Name: "John Doe"}

	
	mockDatabase.On("GetUser", userID).Return(&userData)

	
	mockLogger.On("LogInfo", "Данные взяты из базы").Return("Logged: Данные взяты из базы")
	mockLogger.On("LogInfo", "Данные взяты из кэша").Return("Logged: Данные взяты из кэша")

	
	userFromDB := userProxy.GetData(userID)
	assert.Equal(t, &userData, userFromDB, "Expected user data from database")
	mockDatabase.AssertExpectations(t)
	
	userFromCache := userProxy.GetData(userID)
	assert.Equal(t, &userData, userFromCache, "Expected user data from cache")

	
	mockDatabase.AssertExpectations(t)
	mockLogger.AssertExpectations(t)
}
