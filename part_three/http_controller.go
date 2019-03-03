package partthree

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetPageResponse(url string) (int, []byte, error) {
	request, _ := http.NewRequest("GET", url, nil)
	client := http.DefaultClient
	response, err := client.Do(request) //
	if err != nil {
		return response.StatusCode, nil, fmt.Errorf("client.Do fail")
	}
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	return response.StatusCode, result, err
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Results []Result

func GetTrending(values []byte) Results {
	var results = make([]Result, 0)
	stringReader := strings.NewReader(string(values))
	doc, _ := goquery.NewDocumentFromReader(stringReader)
	doc.Find("div.explore-content ol.repo-list li").Each(func(i int, selection *goquery.Selection) {
		var oneProject Result
		oneProject.Name, _ = selection.Find("div h3 a").Attr("href")
		oneProject.URL = fmt.Sprintf("https://github.com%s", oneProject.Name)
		results = append(results, oneProject)
	})
	return results
}

func GetTrendingTwo(url string) Results {
	_, values, _ := GetPageResponse(url) //
	var results = make([]Result, 0)
	stringReader := strings.NewReader(string(values))
	doc, _ := goquery.NewDocumentFromReader(stringReader)
	doc.Find("div.explore-content ol.repo-list li").Each(func(i int, selection *goquery.Selection) {
		var oneProject Result
		oneProject.Name, _ = selection.Find("div h3 a").Attr("href")
		oneProject.URL = fmt.Sprintf("https://github.com%s", oneProject.Name)
		results = append(results, oneProject)
	})
	return results
}
