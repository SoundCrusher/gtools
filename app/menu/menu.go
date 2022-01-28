package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"gtools/app/vars"
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

	vars.Dark = fyne.NewMenuItem("Dark", func() {
		a.Settings().SetTheme(theme.DarkTheme())
		vars.Dark.Checked = true
		vars.Light.Checked = false
	})
	vars.Light = fyne.NewMenuItem("Light", func() {
		a.Settings().SetTheme(theme.LightTheme())
		vars.Light.Checked = true
		vars.Dark.Checked = false
	})
	themeItem := fyne.NewMenuItem("Theme", func() {})
	themeItem.ChildMenu = fyne.NewMenu("", vars.Dark, vars.Light)

	// Help -> 关于
	about := fyne.NewMenuItem("About", func() {})

	// 主菜单
	settings := fyne.NewMenu("Setting", themeItem)
	helps := fyne.NewMenu("Help", help, about)

	vars.Light.Action()
	window.SetMainMenu(fyne.NewMainMenu(settings, helps))
}
