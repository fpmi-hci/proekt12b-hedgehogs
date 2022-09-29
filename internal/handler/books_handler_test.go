package handler

import (
	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/service"
	mock_service "github.com/fpmi-hci/proekt12b-hedgehogs/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_GetAllBooks(t *testing.T) {
	type mockBehavior func(s *mock_service.MockBooks)

	testTable := []struct {
		name                string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_service.MockBooks) {

				s.EXPECT().GetAllBooks()
			},
			expectedStatusCode:  200,
			expectedRequestBody: `null`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			books := mock_service.NewMockBooks(c)
			testCase.mockBehavior(books)

			services := &service.Service{Books: books}
			handler := NewHandler(services)

			r := gin.New()
			r.GET("/books", handler.getAllBooks)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/books", nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())

		})
	}
}
