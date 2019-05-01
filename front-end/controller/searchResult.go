package controller

import (
	"distributed-web-crawler/crawler/models"
	"distributed-web-crawler/front-end/model"
	"distributed-web-crawler/front-end/view"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/valyala/fastjson"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(
	template string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

// localhost:8888/search?q=xxx&from=20
func (h SearchResultHandler) ServeHTTP(
	writer http.ResponseWriter, req *http.Request) {

	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))

	if err != nil {
		from = 0
	}

	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.view.Render(writer, page)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h SearchResultHandler) getSearchResult(
	q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	result.Query = q

	var s string
	if q != "" {
		s = fmt.Sprintf("http://localhost:9200/dating_profile/_search?q=%s&from=%d", q, from)
	} else {
		s = fmt.Sprintf("http://localhost:9200/dating_profile/_search?from=%d", from)
	}

	fmt.Printf("Request: %s\n", s)
	resp, err := http.Get(s)

	//resp, err := h.client.Search().
	//	Query(elastic.NewQueryStringQuery(q)).
	//	From(from).
	//	Do(context.Background())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {

		}
		bodyString := string(bodyBytes)

		var p fastjson.Parser
		v, err := p.Parse(bodyString)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Body: %v\n", v)

		hitsObj := v.GetObject("hits")
		fmt.Printf("hits object: %v\n", hitsObj)

		hits := hitsObj.Get("total").GetInt()
		fmt.Printf("hits: %d\n", hits)
		result.Hits = hits

		fmt.Printf("from: %d\n", from)
		result.Start = from

		hitsArray, _ := hitsObj.Get("hits").Array()

		for _, item := range hitsArray {

			sourceObj := item.GetObject("_source")
			payload := sourceObj.Get("Payload")

			profile := models.Profile{}
			profile.Name = payload.Get("Name").String()
			profile.Age = payload.GetInt("Age")
			profile.Education = payload.Get("Education").String()
			profile.Marriage = payload.Get("Marriage").String()
			profile.Income = payload.Get("Income").String()
			profile.Height = payload.GetInt("Height")
			profile.Gender = payload.Get("Gender").String()
			result.Items = append(result.Items, profile)
		}

	}

	fmt.Printf("result: %v+", result)

	return result, nil
}
