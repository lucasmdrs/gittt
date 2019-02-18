package main

import (
	"log"
	"net/http"

	"github.com/lucasmdrs/gittt"
)

func main() {
	g := gittt.Init()

	g.ListenAllEvents()

	always := g.ConditionBuilder(gittt.AnyEvent, g.ConditionAlways, nil)

	logPayload := g.ActionBuilder(g.ActionLogPayload, nil)

	always.AddAction(logPayload)

	g.AddConditions(always)

	http.HandleFunc("/webhook", g.Handler)

	log.Println("Wait for connections on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
