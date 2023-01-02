package expense

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func CreateExpensesHandler(c echo.Context) error {
	var e Expense
	err := c.Bind(&e)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	row := db.QueryRow("INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id", e.Title, e.Amount, e.Note, pq.Array(e.Tags))
	err = row.Scan(&e.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, e)
}
