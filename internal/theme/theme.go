package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var PRIMARY_COLOR = color.RGBA{R: 230, G: 216, B: 223, A: 1}
var SECONDARY_COLOR = color.RGBA{R: 0x27, G: 0x26, B: 0x39, A: 0x1}

type AppTheme struct {
	fyne.Theme
}

func New() fyne.Theme {
	return &AppTheme{theme.DefaultTheme()}
}

func (at *AppTheme) Color(n fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground:
		return SECONDARY_COLOR
	case theme.ColorNameInputBackground:
		return SECONDARY_COLOR
	case theme.ColorNamePrimary:
		return PRIMARY_COLOR
	}

	return at.Theme.Color(n, theme.VariantDark)
}
