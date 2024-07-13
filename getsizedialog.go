package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type numericalEntry struct {
	widget.Entry
}

func newNumericalEntry() *numericalEntry {
	entry := &numericalEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *numericalEntry) TypedRune(r rune) {
	if (r >= '0' && r <= '9') || r == '.' {
		e.Entry.TypedRune(r)
	}
}
func (s *Size) getSizeDialog(win fyne.Window) {
	width := newNumericalEntry()
	heighth := newNumericalEntry()

	getSizeDialog := dialog.NewForm("Custom Size", "OK", "Cancel",
		[]*widget.FormItem{
			{Text: "Heighth", Widget: heighth},
			{Text: "Width", Widget: width},
		},
		func(b bool) {

			s.heighth, _ = strconv.ParseFloat(heighth.Text, 64)
			s.width, _ = strconv.ParseFloat(width.Text, 64)
			refreshStatus()
		},
		win)

	getSizeDialog.Show()

}
