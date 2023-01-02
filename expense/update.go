package expense

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func UpdateExpenseHandler(c echo.Context) error {
	var e Expense
	id := c.Param("id")
	err := c.Bind(&e)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	row := db.QueryRow("SELECT id, title, amount, note, tags FROM expenses WHERE id = $1", id)
	err = row.Scan()
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, Err{Message: "expense not found"})
	}

	stmt, err := db.Prepare("UPDATE expenses SET title=$1, amount=$2, note=$3, tags=$4 WHERE id=$5")
	_, err = stmt.Query(e.Title, e.Amount, e.Note, pq.Array(e.Tags), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query expense statment:" + err.Error()})
	}
	e.ID, err = strconv.Atoi(id)
	return c.JSON(http.StatusOK, e)
}
