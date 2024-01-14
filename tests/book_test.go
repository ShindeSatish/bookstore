package tests

import (
	"encoding/json"
	"github.com/ShindeSatish/bookstore/internal/dto"
	"net/http"
	"testing"
)

const (
	BooksUrl = "/books"
)

type getBooksTestCase struct {
	name         string
	expectedCode int
}

type APIResponse struct {
	Success bool
	Message string
	Data    []dto.BookResponse
}

func TestGetBooks(t *testing.T) {
	testCases := []getBooksTestCase{
		{
			name:         "Get Books Success",
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a GET request
			req, err := http.NewRequest("GET", testServer.URL+BooksUrl, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
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

			// Parse response body
			var apiResponse APIResponse
			err = json.NewDecoder(resp.Body).Decode(&apiResponse)
			if err != nil {
				t.Fatalf("Failed to decode response body: %v", err)
			}

			if apiResponse.Success == true {
				books := apiResponse.Data
				if len(books) == 0 {
					t.Errorf("%s: No books found in response", tc.name)
				}
			}
		})
	}
}
