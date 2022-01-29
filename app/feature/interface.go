package feature

import "fyne.io/fyne/v2"

// Creator 窗口创建
type Creator interface {

	// NewBorder 窗口创建
	NewBorder(win fyne.Window) fyne.CanvasObject
}
