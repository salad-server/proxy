package routes

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/salad-server/proxy/util"
)

var index = 0

func infoLog(user, path string, values url.Values) {
	values.Del("h")
	values.Del("u")
	log.Printf(
		"[%s] %s?%s\n", user, path, values.Encode(),
	)
}

func (router *Router) cycleUser() [2]string {
	if index+1 >= len(router.Users) {
		index = 0
	} else {
		index++
	}

	return router.Users[index]
}

func (router *Router) Bancho(w http.ResponseWriter, r *http.Request) {
	query := r.URL.RequestURI()
	ban, err := url.Parse(query)

	if err != nil {
		util.BadRequest(w)
		return
	}

	values := ban.Query()
	user := router.cycleUser()

	values.Del("h")
	values.Del("u")
	values.Set("u", user[0])
	values.Set("h", user[1])

	rurl := fmt.Sprintf("https://osu.ppy.sh%s?%s", r.URL.Path, values.Encode())
	res, err := http.Get(rurl)

	if err != nil {
		util.InternalError(w)
		log.Println(err)

		return
	}

	data, _ := ioutil.ReadAll(res.Body)
	infoLog(user[0], r.URL.Path, values)
	w.Write(data)
}
