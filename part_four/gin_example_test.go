package partfour

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
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

func TestMain(m *testing.M) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}
	POSTGRES, err = gorm.Open("sqlite3", db)
	if err != nil {
		log.Fatalf("can't open gorm connection: %s", err)
	}
	POSTGRES.LogMode(true)
	sqlMock = mock
	code := m.Run()
	os.Exit(code)

}

func TestGetNameHandler(t *testing.T) {
	tests := [3]struct {
		name string
		id   string
	}{
		{
			name: "get one record by id = 1",
			id:   "1",
		},
		//{
		//	name: "get one record by id = 2",
		//	id:   "2",
		//}, {
		//	name: "get one record by id = 3",
		//	id:   "3",
		//},
	}
	g := gin.Default()
	v1 := g.Group("/v1")
	Register(v1)

	recordSQL := `SELECT * FROM items WHERE items.deleted_at IS NULL AND (id = $1) ORDER BY "items"."id" LIMIT 1`

	Convey(tests[0].name, t, func() {
		w := httptest.NewRecorder()
		sqlMock.ExpectQuery(FixRule(recordSQL)).WithArgs(tests[0].id).WillReturnRows(GetRowsForItem(NewItems()[0]))
		request, _ := http.NewRequest("GET", fmt.Sprintf("/v1/name/%s", tests[0].id), nil)
		g.ServeHTTP(w, request)
		So(w.Body.String(), ShouldContainSubstring, "XieWei")
		var result Item
		json.Unmarshal(w.Body.Bytes(), &result)
		So(result.ID, ShouldEqual, 1)
		So(result.Name, ShouldEqual, "XieWei")
	})
	//Convey(tests[1].name, t, func() {
	//	w := httptest.NewRecorder()
	//	sqlMock.ExpectQuery(FixRule(recordSQL)).WithArgs(tests[1].id).WillReturnRows(GetRowsForItem(NewItems()[1]))
	//	request, _ := http.NewRequest("GET", fmt.Sprintf("/v1/name/%s", tests[1].id), nil)
	//	g.ServeHTTP(w, request)
	//	//fmt.Println(w.Body.String())
	//	So(w.Body.String(), ShouldContainSubstring, "WuXiaoShen")
	//	var result Item
	//	json.Unmarshal(w.Body.Bytes(), &result)
	//	So(result.ID, ShouldEqual, 2)
	//	So(result.Name, ShouldEqual, "WuXiaoShen")
	//})
	//Convey(tests[2].name, t, func() {
	//	w := httptest.NewRecorder()
	//	sqlMock.ExpectQuery(FixRule(recordSQL)).WithArgs(tests[2].id).WillReturnError(fmt.Errorf("record not found"))
	//	request, _ := http.NewRequest("GET", fmt.Sprintf("/v1/name/%s", tests[2].id), nil)
	//	g.ServeHTTP(w, request)
	//	fmt.Println(w.Body.String())
	//	So(w.Body.String(), ShouldContainSubstring, "record not found")
	//})
}
