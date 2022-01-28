package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/flopp/go-findfont"

	"gtools/app/content"
	"gtools/app/menu"
	"gtools/app/vars"
)

func main() {
	a := app.NewWithID(vars.AppTitle)
	a.SetIcon(&fyne.StaticResource{StaticName: vars.AppTitle, StaticContent: []byte{}})
	LogLifecycle(a)

	window := a.NewWindow(vars.AppTitle)
	window.Resize(fyne.NewSize(920, 500))
	window.SetMaster()
	window.CenterOnScreen()
	vars.TopWindow = window

	// 菜单栏
	menu.SetMenu(a, window)
	// 内容
	content.SetContent(a, window)

	window.ShowAndRun()
}

func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		// 楷体:simkai.ttf
		// 黑体:simhei.ttf
		if strings.Contains(path, "STHeiti") {
			log.Println("使用字体: ", path)
			if err := os.Setenv("FYNE_FONT", path); err != nil {
				return
			}
			break
		}
	}
	fmt.Println("=============")
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
