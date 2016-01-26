package cms

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func (p *CmsEngine) dict(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if buf, err := exec.Command("sdcv", "--data-dir", "dict", r.URL.Query().Get("keyword")).Output(); err == nil {
		p.Html(w, strings.Replace(string(buf), "\n", "<br>", -1))
	} else {
		p.Abort(w, err)
	}

}
