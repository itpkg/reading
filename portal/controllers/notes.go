package controllers

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"

	"fmt"
	"github.com/astaxie/beego"
	"github.com/itpkg/reading/portal/utils"
)

const notesRoot = "tmp/notes"

type NotesController struct {
	BaseController
}

func (p *NotesController) Prepare() {
	p.BaseController.Prepare()

	var notes map[string]string
	const key = "cache://notes"
	obj := utils.BM.Get(key)
	var err error
	if obj == nil {
		notes = make(map[string]string)
		root := path.Join(notesRoot, p.Lang)
		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if info.Mode().IsRegular() {
				name := info.Name()
				notes[path[len(root)+1:len(path)-3]] = name[0 : len(name)-3]
			}
			return nil
		})
		var buf []byte
		if buf, err = utils.ToJson(notes); err == nil {
			err = utils.BM.Put(key, buf, time.Hour*24)
		}
	} else {
		err = utils.FromJson(obj.([]byte), &notes)
	}

	if err == nil {
		p.Data["items"] = notes
	} else {
		beego.Error(err)
		p.Abort("500")
	}

}

// @router /notes [get]
func (p *NotesController) Index() {
	p.TplName = "notes/index.tpl"
}

// @router /notes/* [get]
func (p *NotesController) Show() {
	name := p.Ctx.Input.Param(":splat")
	key := fmt.Sprintf("cache://notes/%s/%s", p.Lang, name)

	var buf []byte
	var err error

	obj := utils.BM.Get(key)
	if obj == nil {
		if buf, err = ioutil.ReadFile(path.Join(
			notesRoot,
			p.Lang,
			name+".md"),
		); err == nil {
			err = utils.BM.Put(key, buf, time.Hour*24)
		}
	} else {
		buf = obj.([]byte)
	}

	if err == nil {
		p.Data["body"] = string(buf)
		p.TplName = "notes/show.tpl"
	} else {
		beego.Error(err)
		p.Abort("404")
	}

}
