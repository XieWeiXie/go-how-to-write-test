package partthree

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	. "bou.ke/monkey"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetPageResponse(t *testing.T) {
	var client *http.Client
	guard := PatchInstanceMethod(reflect.TypeOf(client), "Do", func(*http.Client, *http.Request) (*http.Response, error) {
		var response http.Response
		response.StatusCode = 400

		return &response, fmt.Errorf("%s", "http fail")
	})
	defer guard.Unpatch()
	tests := [2]struct {
		name  string
		url   string
		want1 int
		want2 []byte
		want3 error
	}{
		{
			name:  "not ok",
			url:   "http://www.baidu.com",
			want1: 400,
			want2: []byte{},
			want3: fmt.Errorf("%s", "http fail"),
		},
		{
			name:  "not ok",
			url:   "http://www.hao123.com",
			want1: 400,
			want2: []byte{},
			want3: fmt.Errorf("%s", "http fail"),
		},
	}
	Convey(tests[0].name, t, func() {
		code, _, err := GetPageResponse(tests[0].url)
		So(code, ShouldEqual, tests[0].want1)
		//So(result, ShouldEqual, tests[0].want2)
		So(err, ShouldNotBeNil)
	})
	Convey(tests[1].name, t, func() {
		code, _, err := GetPageResponse(tests[1].url)
		So(code, ShouldEqual, tests[1].want1)
		//So(result, ShouldEqual, tests[1].want2)
		So(err, ShouldNotBeNil)

	})
}

func TestGetTrending(t *testing.T) {
	_, values, _ := GetPageResponse("https://github.com/trending/go?since=daily")
	tests := []struct {
		name   string
		values []byte
	}{
		{
			name:   "test get trending with mock",
			values: []byte(PageString),
		},
		{
			name:   "test get trending",
			values: values,
		},
	}
	for _, test := range tests {
		Convey(test.name, t, func() {
			fmt.Println(GetTrending(test.values))
		})
	}

}

func TestGetTrendingTwo(t *testing.T) {

	guard := Patch(GetPageResponse, func(_ string) (int, []byte, error) {
		return 200, []byte(PageString), nil
	})
	defer guard.Unpatch()

	tests := []struct {
		name string
		url  string
	}{
		{
			name: "get trending",
		},
	}
	Convey(tests[0].name, t, func() {
		results := GetTrendingTwo(tests[0].url)
		So(results, ShouldNotBeNil)
		So(len(results), ShouldEqual, 1)
		So(results[0].URL, ShouldContainSubstring, "github.com")
	})
}
