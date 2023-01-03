package handle

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetExpenses(t *testing.T) {
	// Arrange
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/expenses", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	newsMockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
		AddRow("1", "Expense 1", 100, "Note for expense 1", `{tag1,tag2}`).
		AddRow("2", "Expense 2", 50, "Note for expense 2", `{tag3,tag4}`)

	db, mock, err := sqlmock.New()
	mock.ExpectPrepare(regexp.QuoteMeta(`SELECT id, title, amount, note, tags FROM expenses`)).
		ExpectQuery().
		WillReturnRows(newsMockRows)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	expected := "[{\"id\":1,\"title\":\"Expense 1\",\"amount\":100,\"note\":\"Note for expense 1\",\"tags\":[\"tag1\",\"tag2\"]},{\"id\":2,\"title\":\"Expense 2\",\"amount\":50,\"note\":\"Note for expense 2\",\"tags\":[\"tag3\",\"tag4\"]}]"
	h := handler{db}
	c := e.NewContext(req, rec)

	// Act
	err = h.GetExpenses(c)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
	}
}
