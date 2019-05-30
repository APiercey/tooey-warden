package main

import (
	"log"

	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	bwItems := state.getBitWardenItems()

	draw(bwItems)
	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID { // event string/identifier
			case "q", "<C-c>": // press 'q' or 'C-c' to quit
				return
			case "<Resize>":
				draw(bwItems)
			}
		}
	}
}
func draw(bwItems []BwItem) {
	l := createItemsList()
	l = populateItemsList(bwItems, l)

	f := createFilterControl()

	p := widgets.NewParagraph()
	p.Text = "Test"

	grid := createGrid()

	grid.Set(
		ui.NewRow(0.92,
			ui.NewCol(0.3, l),
			ui.NewCol(0.7, p),
		),
		ui.NewRow(0.08,
			ui.NewCol(1.0, f),
		),
	)

	ui.Render(grid)
}

func createGrid() *ui.Grid {
	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	return grid
}

func createItemsList() *widgets.List {
	l := widgets.NewList()
	l.Title = "All Items"
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false

	return l
}

func populateItemsList(bwItems []BwItem, l *widgets.List) *widgets.List {
	items := []string{}
	for _, elem := range bwItems {
		items = append(items, elem.Name)
	}

	l.Rows = items

	return l
}

func createFilterControl() *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Text = "Press / to filter (press ? for help)"
	p.Title = "Filter"

	return p
}
