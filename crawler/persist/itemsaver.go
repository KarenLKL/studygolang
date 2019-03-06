package persist

import (
	"context"
	"encoding/json"
	"github.com/KarenLKL/studygolang/crawler/engine"
	"github.com/KarenLKL/studygolang/crawler/model"
	"github.com/KarenLKL/studygolang/crawler/zhenai/parser"
	"log"
)

type ItemSaver struct {
	ElasticClient *elastic.Client
	Index         string
}

// item入库处理
func (i ItemSaver) Saver() (chan engine.Item, error) {
	itemSaverChan := make(chan engine.Item)
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	go func() {
		count := 0
		for {
			item := <-itemSaverChan
			userDetail := parser.ParseUserDetail(item.UserInfo)
			if userDetail != nil {
				log.Printf("save item:got count:#%d,value:%v", count, userDetail)
			}
			item.UserInfo = userDetail
			err := save(client, item, i.Index)
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

func (i ItemSaver) All(esType string) model.Response {
	result, err := i.ElasticClient.Get().Index(i.Index).Type(dbType).Do(context.Background())
	if err != nil {
		return model.Response{}
	}
	var response model.Response
	log.Print(result)
	return response
}
