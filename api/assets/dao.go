package assets

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/itpkg/reading/api/config"
	"gopkg.in/olivere/elastic.v3"
)

type Dao struct {
	Client *elastic.Client `inject:""`
	Cfg    *config.Model   `inject:""`
}

const TYPE = "assets"

func (p *Dao) Del(title string) error {
	_, err := p.Client.Delete().
		Index(p.Cfg.ElasticSearch.Index).
		Type(TYPE).
		Id(title).
		Do()
	return err
}

//func (p*Dao) Count(type_ string) (int,error) {
//	q := elastic.NewMatchQuery("type", type_)
//
//	r,e := p.Client.Count(p.Cfg.ElasticSearch.Index).
//	Type(TYPE).
//	Query(q).
//	Do()
//	return int(r),e
//
//}

func (p *Dao) List(type_ string, from, size int) (int64, []*Item, error) {
	var q elastic.Query
	if type_ == "" {
		q = elastic.NewMatchAllQuery()
	} else {
		q = elastic.NewMatchQuery("type", type_) //fixme
	}

	r, e := p.Client.Search().
		Index(p.Cfg.ElasticSearch.Index).
		Type(TYPE).
		From(from).Size(size).
		Query(q).
		Do()

	if e != nil {
		return 0, nil, e
	}

	items := make([]*Item, 0)

	for _, item := range r.Each(reflect.TypeOf(Item{})) {
		if t, ok := item.(Item); ok {
			items = append(items, &t)
		}
	}

	//	for _, hit := range r.Hits.Hits {
	//		ids = append(ids, hit.Id)
	//	}

	return r.Hits.TotalHits, items, nil
}

func (p *Dao) Set(type_, title, body string) error {
	item := Item{Title: title, Type: type_, Body: body}
	_, err := p.Client.Index().
		Index(p.Cfg.ElasticSearch.Index).Type(TYPE).Id(item.Id()).
		BodyJson(item).Do()
	return err
}

func (p *Dao) Get(id string) (*Item, error) {
	r, e := p.Client.Get().Index(p.Cfg.ElasticSearch.Index).Type(TYPE).Id(id).Do()
	if e != nil {
		return nil, e
	}

	if r.Found {
		i := Item{}
		e := json.Unmarshal(*r.Source, &i)
		return &i, e
	} else {
		return nil, errors.New("not found")
	}
}
