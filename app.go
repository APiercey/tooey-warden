// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

// +build ignore

package main

import (
	"encoding/json"
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
	"log"
	"os"
	"os/exec"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	bwItems := getBitWardenItems()

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

// ReadOp Perform a read from the application state
type ReadOp struct {
	key  int
	resp chan int
}

// WriteOp perform a read to the application state
type WriteOp struct {
	key  int
	val  int
	resp chan bool
}

// ApplicationState the state of the application
type ApplicationState struct {
	Items []BwItems
}

// Application contains the structure of the application data, including state and read and write queues.
type Application struct {
	State      ApplicationState
	WriteQueue chan WriteOp
	ReadQueue  chan ReadOp
}

func createApplicationState() Application {
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

func getItems() map[string]string {
	dict := make(map[string]string)
	dict["Gitlab"] = "asd"
	dict["Github"] = "ghj"
	dict["bitbucket"] = "uuu"

	return dict
}

func readBitWardenItems() []byte {
	cmd := exec.Command("bw", "list", "items")
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		os.Exit(1)
	}

	return out
}

// BwLogin Bitwarden login information
type BwLogin struct {
	Username string
	Password string
}

// BwItem struct of a Bitwarden item
type BwItem struct {
	ID    string
	Name  string
	Login BwLogin
}

// BwItemCollection a collection of BwItems
type BwItemCollection struct {
	Collection []BwItem
}

func getBitWardenItems() []BwItem {
	data := readBitWardenItems()

	bwItems := make([]BwItem, 0)
	json.Unmarshal(data, &bwItems)

	return bwItems
}
