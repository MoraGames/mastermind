package main

import (
	"fmt"
	"github.com/MoraGames/terminal"
	"github.com/nsf/termbox-go"
)

type Option string
type Menu struct {
	title string
	description string
	options []Option

	style SelectionStyle
	selectedOption int
	defaultSelectedOption int
}

func NewMenu(title, description string, options []Option, style SelectionStyle, defaultSelectedOption int) (*Menu, error) {
	if defaultSelectedOption < 0 || defaultSelectedOption >= len(options) {
		return nil, fmt.Errorf("the chosen selection (%d) is not applicable. The menu tolerates the solutions between [0; %d)", defaultSelectedOption, len(options))
	}

	var m Menu

	m.title = title
	m.description = description
	m.options = options
	m.style = style
	m.selectedOption = defaultSelectedOption
	m.defaultSelectedOption = defaultSelectedOption

	return &m, nil
}

func (m *Menu) IncraseSelection() {
	m.selectedOption++
	if m.selectedOption == len(m.options) {
		m.selectedOption = 0
	}
}

func (m *Menu) DecraseSelection() {
	m.selectedOption--
	if m.selectedOption == -1 {
		m.selectedOption = len(m.options) - 1
	}
}

func (m *Menu) SetSelection(selectedOption int) error {
	if selectedOption < -1 || selectedOption >= len(m.options) {
		return fmt.Errorf("the chosen selection (%d) is not applicable. The menu tolerates the solutions between [0; %d)", selectedOption, len(m.options))
	}
	if selectedOption == -1 {
		m.selectedOption = m.defaultSelectedOption
	} else {
		m.selectedOption = selectedOption
	}
	return nil
}

func (m Menu) GetSelection() int {
	return m.selectedOption
}

func (m Menu) Print() {
	terminal.Clear()
	fmt.Println()

	fmt.Println(m.title)
	fmt.Println(m.description + "\n")

	maxLength := -1
	for _, option := range m.options {
		if len := len(option); len > maxLength {
			maxLength = len
		}
	}

	for i := 0; i < len(m.options); i++ {
		lengthDifference := maxLength - len(m.options[i])

		var modifier string
		for s := 0; s < lengthDifference/2; s++ {
			modifier += " "
		}
		spacing := [2]string{"", ""}
		for i2 := 0; i2 < 2; i2++ {
			for s := 0; s < m.style[i2].Length(); s++ {
				spacing[i2] += " "
			}
		}

		if m.selectedOption == i {
			fmt.Print(m.style[0])
		}else{
			fmt.Print(spacing[0])
		}
		fmt.Print(" ", modifier, m.options[i], modifier, " ")
		if lengthDifference%2 == 1 {
			fmt.Print(" ")
		}
		if m.selectedOption == i {
			fmt.Print(m.style[1])
		}else{
			fmt.Print(spacing[1])
		}
		fmt.Println()
	}

	fmt.Println()
}

func (m Menu) Handler() int {
	for exit := false; exit == false; {
		m.Print()
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			switch event.Key {
			case termbox.KeyArrowDown:
				m.IncraseSelection()
			case termbox.KeyArrowUp:
				m.DecraseSelection()
			case termbox.KeyEnter:
				exit = true
			}
		case termbox.EventError:
			panic(event.Err)
		}
	}
	terminal.Clear()

	return m.GetSelection()
}

/*
	-- Menù Principale -- -----------------------------------
	Utilizza i tasti freccia per orientarti nel menù.

	<-   Nuova  Partita   ->
	   Classifica  Record
	      Impostazioni
	     Esci dal Gioco
*/