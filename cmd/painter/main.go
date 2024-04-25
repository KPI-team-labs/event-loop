package main

import (
	"net/http"

	"github.com/KPI-team-labs/event-loop/painter"
	"github.com/KPI-team-labs/event-loop/painter/lang"
	"github.com/KPI-team-labs/event-loop/ui"
)

func main() {
	var (
		pv ui.Visualizer // Візуалізатор створює вікно та малює у ньому.

		// Потрібні для частини 2.
		opLoop painter.Loop // Цикл обробки команд.
		state  = lang.NewUistate()
		parser = lang.ParserWithState(state) // Парсер команд.
	)

	//pv.Debug = true
	pv.Title = "Simple painter"

	pv.OnScreenReady = opLoop.Start
	opLoop.Receiver = &pv

	go func() {
		http.Handle("/", lang.HttpHandler(&opLoop, parser))
		_ = http.ListenAndServe("localhost:17000", nil)
	}()

	pv.Main()
	opLoop.StopAndWait()
}
