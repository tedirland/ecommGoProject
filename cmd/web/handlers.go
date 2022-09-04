package main

import (
	"net/http"
)

func (app *applicaiton) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, *r, "terminal", nil); err != nil {
		app.errorLog.Println(err)
	}

}
