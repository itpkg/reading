package auth

import (
	"github.com/itpkg/reading/api/core"
)

func (p *AuthEngine) Mount(rt core.Router) {
	rt.GET("/admin/roles", p.getAdminRoles)
	rt.POST("/admin/roles", p.postAdminRoles)
	rt.GET("/admin/site/info", p.getAdminSiteInfo)
	rt.POST("/admin/site/info", p.postAdminSiteInfo)
	rt.GET("/admin/site/seo", p.getAdminSiteSeo)
	rt.POST("/admin/site/seo", p.postAdminSiteSeo)
	rt.GET("/admin/site/secrets", p.getAdminSiteSecrets)
	rt.POST("/admin/site/secrets", p.postAdminSiteSecrets)

	rt.GET("/users/logs", p.getUsersLogs)

	rt.GET("/oauth/sign_in", p.getSignIn)
	rt.POST("/oauth/sign_in", p.postSignIn)
}
