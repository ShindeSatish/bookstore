package tests

import (
	"net/http"
	"strings"
	"testing"
)

type createOrderTestCase struct {
	name         string
	request      string
	expectedCode int
	token        string
}

func TestCreateOrder(t *testing.T) {
	TestUserLogin(t)
	createOrderTestCases := []createOrderTestCase{
		{
			name:         "Create Order Success",
			request:      `{"additional_charges":0,"discount_amount":0,"items": [{"book_id": 1, "quantity": 2}], "total_price": 25.0}`,
			expectedCode: http.StatusOK,
			token:        AuthenticationToken, // Replace with a valid token for testing
		},
		{
			name:         "Create Order Unauthorized",
			request:      `{"items": [{"book_id": 1, "quantity": 2}], "total_price": 25.0}`,
			expectedCode: http.StatusUnauthorized,
			token:        "",
		},
	}

	for _, tc := range createOrderTestCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup the HTTP request
			req, err := http.NewRequest("POST", testServer.URL+"/order", strings.NewReader(tc.request))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Set the authorization header if a token is provided
			if tc.token != "" {
				req.Header.Set("Authorization", tc.token)
			}

			// Perform the request
			resp, err := testServer.Client().Do(req)
			if err != nil {
				t.Fatalf("Failed to send request: %v", err)
			}
			defer resp.Body.Close()

			// Check the response status code
			if resp.StatusCode != tc.expectedCode {
				t.Errorf("Expected status code %v; got %v", tc.expectedCode, resp.StatusCode)
			}

			// Add additional checks if needed (e.g., response body content)
		})
	}
}
