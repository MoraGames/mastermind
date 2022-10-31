package main

import (
	"github.com/nsf/termbox-go"
)

func main(){
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	mainMenu()
}

const (
	Arrow_Left IndicatorStyle = "->"
	Arrow_Right IndicatorStyle = "<-"
)

func mainMenu() {
	menu, err := NewMenu(
		"-- Menù Principale -- ----------------------------",
		"Utilizza i tasti freccia per orientarti nel menù.\nPremi invio per selezionare.",
		[]Option{"Nuova  Partita", "Classifica  Record", "Impostazioni", "Esci dal Gioco"},
		SelectionStyle{Arrow_Left, Arrow_Right},
		0,
	)
	if err != nil {
		panic(err)
	}

	switch menu.Handler() {
		case 0:
			//NewGame()
		case 1:
			//ShowRecords()
		case 2:
			//EditSettings()
		case 3:
			//Quit
	}
}