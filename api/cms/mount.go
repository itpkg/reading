package cms

import (
	"github.com/itpkg/reading/api/core"
)

func (p *CmsEngine) Mount(rt core.Router) {
	rt.GET("/video/channels/:id", p.showChannel)
	rt.GET("/video/channels", p.listChannel)
	rt.GET("/video/playlist/:id", p.showPlaylist)
	rt.GET("/video/playlist", p.listPlaylist)
	rt.GET("/video/items/:id", p.showVideo)
	rt.GET("/video/items", p.listVideo)

	rt.GET("/cms/books/:id", p.showBook)
	rt.GET("/cms/books", p.listBook)

	rt.GET("/cms/articles/:id", p.showArticle)
	rt.GET("/cms/articles", p.listArticle)
	rt.GET("/cms/tags/:id", p.showTag)
	rt.GET("/cms/tags", p.listTag)

	rt.GET("/attachments", p.listAttachment)
	rt.POST("/attachments", p.saveAttachment)
	rt.DELETE("/attachment/:id", p.removeAttachment)

	rt.POST("/dict", p.dict)
}