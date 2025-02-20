package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"echo-store-api/internal/domain"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Register(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserUsecase) Login(email, password string) (string, error) {
	args := m.Called(email, password)
	return args.String(0), args.Error(1)
}

func (m *MockUserUsecase) GetProfile(id uint) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserUsecase) UpdateProfile(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestUserHandler_Register(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    RegisterRequest
		mockError      error
		expectedStatus int
	}{
		{
			name: "Success",
			requestBody: RegisterRequest{
				Email:    "test@example.com",
				Password: "password123",
				Name:     "Test User",
			},
			mockError:      nil,
			expectedStatus: http.StatusCreated,
		},
		{
			name: "User Already Exists",
			requestBody: RegisterRequest{
				Email:    "existing@example.com",
				Password: "password123",
				Name:     "Existing User",
			},
			mockError:      errors.New("user already exists"),
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			e := echo.New()
			mockUsecase := new(MockUserUsecase)
			h := NewUserHandler(mockUsecase)

			// Mock expectations
			mockUsecase.On("Register", mock.AnythingOfType("*domain.User")).Return(tt.mockError)

			// Create request
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Test
			err := h.Register(c)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)

			mockUsecase.AssertExpectations(t)
		})
	}
}
