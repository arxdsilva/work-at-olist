package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arxdsilva/olist/storage"
	"github.com/labstack/echo"
)

func TestServer_saveRecord(t *testing.T) {
	recordJSON := `{
		"id": "123",
		"type": "start",
		"timestamp":"2016-02-29T12:00:00Z",
		"call_id": "qualquercoisa",
		"source": "1515151515",
		"destination": "1515151515"
	}`
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/records", strings.NewReader(recordJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	s := &Server{Storage: storage.FakeStorage{}}
	if assert.NoError(t, s.saveRecord(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
