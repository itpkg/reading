package blog

import (
	"encoding/json"
	"errors"

	"github.com/itpkg/reading/api/config"
	"gopkg.in/olivere/elastic.v3"
)

type Dao struct {
	Client *elastic.Client `inject:""`
	Cfg    *config.Model   `inject:""`
}

const TYPE = "blog"

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

func (p *Dao) List(type_ string, from, size int) ([]string, error) {
	q := elastic.NewMatchQuery("type", type_)

	r, e := p.Client.Search().
		Index(p.Cfg.ElasticSearch.Index).
		Type(TYPE).
		From(from).Size(size).
		Query(q).
		Do()

	if e != nil {
		return nil, e
	}

	ids := make([]string, 0)

	if r.Hits != nil {
		for _, hit := range r.Hits.Hits {
			ids = append(ids, hit.Id)
		}
	}
	return ids, nil
}

func (p *Dao) Set(type_, title, body string) error {
	_, err := p.Client.Index().
		Index(p.Cfg.ElasticSearch.Index).Type(TYPE).Id(title).
		BodyJson(Item{Type: type_, Body: body}).Do()
	return err
}

func (p *Dao) Get(title string) (*Item, error) {
	r, e := p.Client.Get().Index(p.Cfg.ElasticSearch.Index).Type(TYPE).Id(title).Do()
	if e != nil {
		return nil, e
	}

	if r.Found {
		i := Item{Title: title}
		e := json.Unmarshal(*r.Source, &i)
		return &i, e
	} else {
		return nil, errors.New("not found")
	}
}
