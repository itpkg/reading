package auth

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (p *AuthEngine) getUsersLogs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user, err := p.Session.User(r)
	if err == nil {
		var logs []Log
		p.Db.Where("user_id = ?", user.ID).Limit(60).Order("id DESC").Find(&logs)
		p.Render.JSON(w, http.StatusOK, logs)
	} else {
		p.Abort(w, err)
	}
}
