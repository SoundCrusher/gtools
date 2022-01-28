package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"log"
)

func SetMenu(a fyne.App, window fyne.Window) {
	// Help -> help
	// Help -> Documents
	help := fyne.NewMenuItem("Help", func() {})
	help.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Help", func() {
			log.Println("Help -> Help")
		}),
		fyne.NewMenuItem("Documents", func() {
			log.Println("Help -> Documents")
		}))

	// Setting -> Theme
	var light *fyne.MenuItem
	var dark *fyne.MenuItem
	dark = fyne.NewMenuItem("Dark", func() {
		a.Settings().SetTheme(theme.DarkTheme())
		dark.Checked = true
		light.Checked = false
	})
	light = fyne.NewMenuItem("Light", func() {
		a.Settings().SetTheme(theme.LightTheme())
		light.Checked = true
		dark.Checked = false
	})
	themeItem := fyne.NewMenuItem("Theme", func() {})
	themeItem.ChildMenu = fyne.NewMenu("", dark, light)

	// Help -> 关于
	about := fyne.NewMenuItem("About", func() {})

	// 主菜单
	settings := fyne.NewMenu("Setting", themeItem)
	helps := fyne.NewMenu("Help", help, about)

	light.Action()
	window.SetMainMenu(fyne.NewMainMenu(settings, helps))
}