package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"gtools/app/content"
	"gtools/app/menu"
	"gtools/app/vars"
	"log"
)

func main() {
	a := app.NewWithID(vars.AppTitle)
	a.SetIcon(&fyne.StaticResource{StaticName: vars.AppTitle, StaticContent: []byte{}})
	LogLifecycle(a)

	window := a.NewWindow(vars.AppTitle)
	window.Resize(fyne.NewSize(900, 480))
	window.SetMaster()
	window.CenterOnScreen()
	vars.TopWindow = window

	// 菜单栏
	menu.SetMenu(a, window)
	// 内容
	content.SetContent(a, window)

	window.ShowAndRun()
}

func LogLifecycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		log.Println("Go Tools - Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		log.Println("Go Tools - Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Go Tools - Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Go Tools - Lifecycle: Exited Foreground")
	})
}

