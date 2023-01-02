package expense

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func GetExpenseHandler(c echo.Context) error {
	id := c.Param("id")
	stmt, err := db.Prepare("SELECT id, title, amount, note, tags FROM expenses WHERE id = $1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query expense statment:" + err.Error()})
	}
	row := stmt.QueryRow(id)
	e := Expense{}
	err = row.Scan(&e.ID, &e.Title, &e.Amount, &e.Note, pq.Array(&e.Tags))
	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, Err{Message: "expense not found"})
	case nil:
		return c.JSON(http.StatusOK, e)
	default:
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan expense:" + err.Error()})
	}
}
