package persist

import (
	"context"
	"encoding/json"
	"github.com/KarenLKL/studygolang/crawler/engine"
	"github.com/KarenLKL/studygolang/crawler/front/model"
	"github.com/KarenLKL/studygolang/crawler/model"
	"github.com/KarenLKL/studygolang/crawler/zhenai/parser"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"reflect"
	"strconv"
)

type ItemSaver struct {
	ElasticClient *elastic.Client
	Index         string
}

// item入库处理
func (i ItemSaver) Saver() (chan engine.Item, error) {
	itemSaverChan := make(chan engine.Item)
	go func() {
		count := 0
		for {
			item := <-itemSaverChan
			userDetail := parser.ParseUserDetail(item.UserInfo)
			if userDetail != nil {
				log.Printf("save item:got count:#%d,value:%v", count, userDetail)
			}
			item.UserInfo = userDetail
			err := i.save(item)
			if err != nil {
				log.Printf("error save item:%+v %s", userDetail, err.Error())
			}
			count++
		}
	}()
	return itemSaverChan, nil
}

const (
	dbIndex = "dating"
	dbType  = "zhenai"
)

func (i ItemSaver) save(row engine.Item) error {

	indexService := i.ElasticClient.Index().Index(i.Index).BodyJson(row)
	if len(row.Type) > 0 {
		indexService.Type(row.Type)
	}
	if len(row.Id) > 0 {
		indexService.Id(row.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	log.Printf("saved item: %v+", row)
	return nil
}

func (i ItemSaver) get(id string) (model.UserDetail, error) {
	result, err := i.ElasticClient.Get().Index(i.Index).Type(dbType).Id(id).Do(context.Background())
	if err != nil {
		return model.UserDetail{}, err
	}
	var userDetail model.UserDetail
	err = json.Unmarshal(*result.Source, &userDetail)
	if err != nil {
		return model.UserDetail{}, err
	}
	return userDetail, nil
}

func (i ItemSaver) All(esType string, q string, from int) (page.SearchResult, error) {
	search := i.ElasticClient.Search(i.Index)
	if len(q) > 0 {
		search = search.Query(elastic.NewQueryStringQuery(q))
	}
	searchResult, err := search.From(from).Do(context.Background())
	if err != nil {
		return page.SearchResult{}, err
	}
	var sr page.SearchResult

	if len(q) > 0 {
		sr.Q = "q=" + q
		if from != 0 {
			sr.Q += "&from=" + strconv.Itoa(from)
		}
	} else {
		if from != 0 {
			sr.Q = "from=" + strconv.Itoa(from)
		}
	}

	sr.Hits = searchResult.Hits.TotalHits
	sr.Start = from
	sr.Items = searchResult.Each(reflect.TypeOf(engine.Item{}))
	return sr, nil
}
