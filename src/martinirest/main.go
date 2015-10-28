package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	. "./data"
)

func main() {
	m := martini.Classic()

	// Use template renderer
	m.Use(render.Renderer())

	//START GET OMIT
	m.Get("/book/:title", func(params martini.Params) string {
		title := params["title"]

		b := new(Book)
		b.Get(title)

		ret, _ := json.MarshalIndent(b, "", "    ")
		return string(ret)
	})
	//END GET OMIT

	// START POST OMIT
	m.Post("/book", func(req *http.Request) string {
		req.ParseForm()
		params := req.Form

		b := new(Book)
		b.Title = params["title"][0]
		b.Author = params["author"][0]
		b.Published, _ = time.Parse("02/01/2006", params["date"][0])

		b.Save()

		ret, _ := json.MarshalIndent(b, "", "    ")
		return string(ret)
	})
	//END POST OMIT

	// Get a display view for the book
	m.Get("/display/book/:title", func(params martini.Params, r render.Render) {
		title := params["title"]

		b := new(Book)
		b.Get(title)

		r.HTML(200, "example", b)

	})

	m.Run()
}
