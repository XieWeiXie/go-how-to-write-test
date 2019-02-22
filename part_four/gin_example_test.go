package partfour

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

var (
	sqlMock sqlmock.Sqlmock
)

func FixRule(value string) string {
	return fmt.Sprintf("[" + value + "]")
}

func GetRowsForItem(item Item) *sqlmock.Rows {
	fields := []string{"id", "created_at", "updated_at", "deleted_at", "name"}
	rows := sqlmock.NewRows(fields)
	rows.AddRow(item.ID, item.CreatedAt, item.UpdatedAt, item.DeletedAt, item.Name)
	return rows

}

func TestMain(t *testing.M) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}
	POSTGRES, err = gorm.Open("sqlite3", db)
	if err != nil {
		log.Fatalf("can't open gorm connection: %s", err)
	}
	POSTGRES.LogMode(false)
	sqlMock = mock

}

func TestGetNameHandler(t *testing.T) {
	tests := [2]struct {
		name string
		id   string
	}{
		{
			name: "get one record by id = 1",
			id:   "1",
		},
		{
			name: "get one record by id = 2",
			id:   "2",
		},
	}
	g := gin.Default()
	v1 := g.Group("/v1")
	Register(v1)

	Convey(tests[0].name, t, func() {
		recordSQL := `SELECT * FROM items WHERE items.deleted_at IS NULL AND (id = $1) ORDER BY "items"."id" LIMIT 1`
		sqlMock.ExpectQuery(FixRule(recordSQL)).WithArgs(tests[0].id).WillReturnRows(GetRowsForItem(NewItems()[0]))

		w := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", fmt.Sprintf("/v1/name/%s", tests[0]), nil)
		g.ServeHTTP(w, request)
		fmt.Println(w.Body.String())
	})
	Convey(tests[1].name, t, func() {

	})
}
