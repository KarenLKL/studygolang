package persist

import (
	"context"
	"encoding/json"
	"github.com/KarenLKL/studygolang/crawler/engine"
	"github.com/KarenLKL/studygolang/crawler/model"
	"github.com/KarenLKL/studygolang/crawler/zhenai/parser"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

// item入库处理
func ItemSaver(index string) (chan engine.Item, error) {
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
			err := save(client, item, index)
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

func save(client *elastic.Client, row engine.Item, index string) error {

	indexService := client.Index().Index(index).BodyJson(row)
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

func get(client *elastic.Client, index, id string) (model.UserDetail, error) {
	result, err := client.Get().Index(index).Type(dbType).Id(id).Do(context.Background())
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
