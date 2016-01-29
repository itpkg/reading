package controllers

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

const notesRoot = "tmp/notes"

type NotesController struct {
	BaseController
}

func (p *NotesController) Prepare(){
	p.BaseController.Prepare()

	notes := make(map[string]string)
	root := path.Join(notesRoot, p.Lang)
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() {
			name := info.Name()
			notes[path[len(root)+1:len(path)-3]] = name[0 : len(name)-3]
		}
		return nil
	})
	p.Data["notes"] = notes

}

// @router /notes [get]
func (p *NotesController) Index() {
	p.TplName = "notes/index.tpl"
}

// @router /notes/* [get]
func (p *NotesController) Show() {
	if buf, err := ioutil.ReadFile(path.Join(
		notesRoot,
		p.Lang,
		p.Ctx.Input.Param(":splat")+".md"),
	); err == nil {
		p.Data["body"] = string(buf)
		p.TplName = "notes/show.tpl"
	} else {
		p.Abort("404")
	}

}
