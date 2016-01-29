package controllers

import (
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

const langKey = "locale"

type LocaleController struct {
	beego.Controller
	i18n.Locale
}

func (p *LocaleController) Prepare() {
	// 1. Check URL arguments.
	lang := p.Input().Get(langKey)

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = p.Ctx.GetCookie(langKey)
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := p.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			lang = al[:5] // Only compare first 5 letters.
		}
	}

	// 4. Default language is English.
	if len(lang) == 0 || !i18n.IsExist(lang) {
		lang = "en-US"
	}

	// Save language information in cookies.
	if lang != p.Ctx.GetCookie(langKey) {
		p.Ctx.SetCookie(langKey, lang)
	}

	// Set language properties.
	beego.Trace("GET LANGUAGE: ", lang)
	p.Lang = lang
	p.Data["Lang"] = lang
	p.Data["Languages"] = i18n.ListLangs()
}

func init() {
	filepath.Walk("conf/locales", func(path string, info os.FileInfo, err error) error {
		name := info.Name()
		if info.Mode().IsRegular() && filepath.Ext(name) == ".ini" {
			lang := name[:len(name)-4]
			beego.Trace("Loading language: " + lang)
			return i18n.SetMessage(lang, path)
		}
		return nil
	})

	beego.AddFuncMap("i18n", i18n.Tr)
}
