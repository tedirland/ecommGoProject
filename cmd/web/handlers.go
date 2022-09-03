package main

import "net/http"

func (app *applicaiton) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Hit the handler")
}
