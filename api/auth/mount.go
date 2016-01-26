package auth

import (
	"github.com/itpkg/reading/api/core"
)

func (p *AuthEngine) Mount(rt core.Router) {

	rt.GET("/admin/users", p.getAdminUsers)
	rt.GET("/admin/site/info", p.getAdminSiteInfo)
	rt.POST("/admin/site/info", p.postAdminSiteInfo)
	rt.GET("/admin/site/seo", p.getAdminSiteSeo)
	rt.POST("/admin/site/seo", p.postAdminSiteSeo)
	rt.GET("/admin/site/secrets", p.getAdminSiteSecrets)
	rt.POST("/admin/site/secrets", p.postAdminSiteSecrets)

	rt.POST("/admin/notices", p.saveNotice)
	rt.DELETE("/admin/notice/:id", p.destroyNotice)

	rt.GET("/admin/locales", p.locales)
	rt.POST("/admin/locales", p.saveLocale)
	rt.DELETE("/admin/locale/:id", p.destroyLocale)

	rt.GET("/users/logs", p.getUsersLogs)

	rt.GET("/oauth/sign_in", p.getSignIn)
	rt.POST("/oauth/sign_in", p.postSignIn)
}
