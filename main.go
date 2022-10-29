package main

import (
	"github.com/MoraGames/terminal"
	"github.com/nsf/termbox-go"
)

func main(){
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	menu, err := NewMenu(
		"-- Menù Principale -- ----------------------------",
		"Utilizza i tasti freccia per orientarti nel menù.",
		[]Option{"Nuova  Partita", "Classifica  Record", "Impostazioni", "Esci dal Gioco"},
		Arrows,
		0,
	)
	if err != nil {
		panic(err)
	}

	for exit := false; exit == false; {
		menu.Print()
		switch event := termbox.PollEvent(); event.Type {
			case termbox.EventKey:
				switch event.Key {
					case termbox.KeyArrowLeft:
						//fmt.Println("<-")
					case termbox.KeyArrowDown:
						menu.IncraseSelection()
						//fmt.Println("\\/")
					case termbox.KeyArrowRight:
						//fmt.Println("->")
					case termbox.KeyArrowUp:
						menu.DecraseSelection()
						//fmt.Println("/\\")
					case termbox.KeyEsc:
						//fmt.Println("Esc")
						menu.SetSelection(3)
					case termbox.KeyEnter:
						if menu.GetSelection() == 3 {
							exit = true
						}
					default:
						//if event.Ch == 'r' {
						//	fmt.Println("r")
						//}
					}
			case termbox.EventError:
				panic(event.Err)
		}
	}

	terminal.Clear()
}