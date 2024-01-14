package tests

import (
	"bytes"
	"encoding/json"
	"github.com/ShindeSatish/bookstore/internal/dto"
	"github.com/ShindeSatish/bookstore/internal/utils"
	"net/http"
	"testing"
)

const (
	UserRegistrationUrl = "/register"
	UserLoginUrl        = "/login"
)

type registrationTestCase struct {
	name         string
	requestBody  string
	expectedCode int
}

var AuthenticationToken string

func TestUserRegistration(t *testing.T) {
	// get random number
	email := utils.RandomEmail()

	testCases := []registrationTestCase{
		{
			name: "Valid Registration",
			requestBody: `{
            "email": "` + email + `",
            "first_name": "Satish",
            "last_name": "Shinde",
            "password": "StrongPassword",
            "phone": "1234567890"
        	}`,
			expectedCode: http.StatusOK,
		}, {
			name: "Valid Registration",
			requestBody: `{
            "email": "satish@gmail.com",
            "first_name": "Satish",
            "last_name": "Shinde",
            "password": "StrongPassword",
            "phone": "1234567890"
        	}`,
			expectedCode: http.StatusOK,
		},
		{
			name: "Invalid Email",
			requestBody: `{
            "email": "invalidemail",
            "first_name": "Test",
            "last_name": "User",
            "password": "password",
            "phone": "1234567890"
        }`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Email Already Registered",
			requestBody: `{
			"email": "` + email + `",
			"first_name": "Test",
			"last_name": "User",
			"password": "password",
			"phone": "1234567890"
			}`,
			expectedCode: http.StatusConflict,
		},
		// TODO Add more test cases for different scenarios
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requestBody := bytes.NewBuffer([]byte(tc.requestBody))

			resp, err := http.Post(testServer.URL+UserRegistrationUrl, "application/json", requestBody)
			if err != nil {
				t.Fatalf("Failed to send request for %s: %v", tc.name, err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tc.expectedCode {
				t.Errorf("%s: Expected status code %v; got %v", tc.name, tc.expectedCode, resp.StatusCode)
			}
		})
	}
}

type loginTestCase struct {
	name         string
	requestBody  string
	expectedCode int
}
type LoginAPIResponse struct {
	Success bool                         `json:"success"`
	Message string                       `json:"message"`
	Data    dto.AuthenticateUserResponse `json:"data"`
}

func TestUserLogin(t *testing.T) {
	TestUserRegistration(t)

	testCases := []loginTestCase{
		{
			name: "Invalid Password",
			requestBody: `{
            "email": "satish@gmail.com",
            "password": "WrongPassword"
        	}`,
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "Valid Login",
			requestBody: `{
            "email": "satish@gmail.com",
            "password": "StrongPassword"
        	}`,
			expectedCode: http.StatusOK,
		},

		// Add more test cases for different scenarios...
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requestBody := bytes.NewBuffer([]byte(tc.requestBody))

			resp, err := http.Post(testServer.URL+UserLoginUrl, "application/json", requestBody)
			if err != nil {
				t.Fatalf("Failed to send request for %s: %v", tc.name, err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tc.expectedCode {
				t.Errorf("%s: Expected status code %v; got %v", tc.name, tc.expectedCode, resp.StatusCode)
			}

			// Only attempt to read the token if the status code is OK
			if resp.StatusCode == http.StatusOK {
				var response LoginAPIResponse
				err := json.NewDecoder(resp.Body).Decode(&response)
				if err != nil {
					t.Fatalf("Failed to decode response body for %s: %v", tc.name, err)
				}

				AuthenticationToken = response.Data.Token
			}
		})
	}
}
