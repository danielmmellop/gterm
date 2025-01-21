package grid

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Grid struct {
	widget.TextGrid
	lastRow  int
	onTapped func()
}

func New(onTapped func()) *Grid {
	grid := &Grid{}
	grid.ExtendBaseWidget(grid)

	grid.onTapped = onTapped

	return grid
}

func (grid *Grid) UpdateContentText(data []string) {
	rows := cmdResultToTextGridRow(data)
	for id, row := range rows {
		grid.SetRow(grid.lastRow+id, row)
		grid.Refresh() // TODO: Confirm if this is necessary to fix the row after a clean command
	}

	grid.lastRow += len(rows)
}

func (grid *Grid) Tapped(_ *fyne.PointEvent) {
	grid.onTapped()
}

func (grid *Grid) Clean() {
	emptyRow := cmdResultToTextGridRow([]string{""})
	rows := len(grid.Rows)
	for id := range rows {
		grid.SetRow(id, emptyRow[0])
	}
	grid.lastRow = 0
}

func cmdResultToTextGridRow(data []string) []widget.TextGridRow {
	rows := []widget.TextGridRow{}

	for _, rowText := range data {
		runeText := []widget.TextGridCell{}
		for _, runeValue := range rowText {
			runeText = append(runeText, widget.TextGridCell{Rune: rune(runeValue)})
		}
		rows = append(rows, widget.TextGridRow{Cells: runeText})
	}

	return rows
}
