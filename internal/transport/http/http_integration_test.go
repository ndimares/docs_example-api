package http_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/speakeasy-api/speakeasy-go-sdk"
	httptransport "github.com/speakeasy-api/speakeasy-test-rest-service/internal/transport/http"
	"github.com/speakeasy-api/speakeasy-test-rest-service/internal/transport/http/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	os.Setenv("SPEAKEASY_SERVER_URL", "localhost:8080")

	speakeasy.Configure(speakeasy.Config{
		APIKey: "test",
	})

	os.Exit(m.Run())
}

func TestServer_Health_Success(t *testing.T) {
	tests := []struct {
		name     string
		wantCode int
	}{
		{
			name:     "success",
			wantCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			u := mocks.NewMockUsers(ctrl)
			d := mocks.NewMockDB(ctrl)

			ht, err := httptransport.New(u, d)
			require.NoError(t, err)
			require.NotNil(t, ht)

			r := mux.NewRouter()

			hr, err := ht.AddRoutes(r)
			require.NoError(t, err)

			w := httptest.NewRecorder()

			d.EXPECT().PingContext(gomock.Any()).Return(nil).Times(1)

			req, err := http.NewRequest(http.MethodGet, "/health", nil)
			require.NoError(t, err)

			hr.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Code)
		})
	}
}

func TestServer_Health_Error(t *testing.T) {
	tests := []struct {
		name     string
		wantErr  string
		wantCode int
	}{
		{
			name:     "fails",
			wantErr:  "test fail",
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			u := mocks.NewMockUsers(ctrl)
			d := mocks.NewMockDB(ctrl)

			ht, err := httptransport.New(u, d)
			require.NoError(t, err)
			require.NotNil(t, ht)

			r := mux.NewRouter()

			hr, err := ht.AddRoutes(r)
			require.NoError(t, err)

			w := httptest.NewRecorder()

			d.EXPECT().PingContext(gomock.Any()).Return(errors.New(tt.wantErr)).Times(1)

			req, err := http.NewRequest(http.MethodGet, "/health", nil)
			require.NoError(t, err)

			hr.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Code)

			var res struct {
				Error string `json:"error"`
			}

			err = json.Unmarshal(w.Body.Bytes(), &res)
			require.NoError(t, err)
			assert.Equal(t, tt.wantErr, res.Error)
		})
	}
}
