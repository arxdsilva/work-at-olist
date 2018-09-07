package api

import (
	"net/http"
	"net/http/httptest"
	"strings"

	check "gopkg.in/check.v1"

	"github.com/arxdsilva/olist/storage"
	"github.com/labstack/echo"
)

func (s *S) TestServer_saveRecord(c *check.C) {
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
	ctx := e.NewContext(req, rec)
	sv := &Server{Storage: storage.FakeStorage{}}
	c.Assert(sv.saveRecord(ctx), check.IsNil)
	c.Assert(http.StatusCreated, check.Equals, rec.Code)
}
