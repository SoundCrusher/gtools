package content

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"gtools/app/data"
	"gtools/app/vars"
)

func BuildTree(a fyne.App, window fyne.Window, title *widget.Label, intro *widget.Label, content *fyne.Container) *widget.Tree {
	return &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return data.TreeNodeMetadata[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := data.TreeNodeMetadata[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := data.TreeNodeMappings[uid]
			if !ok {
				fyne.LogError("Missing contentBorder panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
		},
		OnSelected: func(uid string) {
			if t, ok := data.TreeNodeMappings[uid]; ok {
				a.Preferences().SetString("currentTutorial", uid)
				changePanel(a, window, title, intro, content, t)
			}
		},
	}
}

func changePanel(a fyne.App, window fyne.Window, title *widget.Label, intro *widget.Label, content *fyne.Container, t data.TreeNode) {
	if fyne.CurrentDevice().IsMobile() {
		child := a.NewWindow(t.Title)
		vars.TopWindow = child
		child.SetContent(t.View(vars.TopWindow))
		child.Show()
		child.SetOnClosed(func() {
			vars.TopWindow = window
		})
		return
	}

	title.SetText(t.Title)
	intro.SetText(t.Intro)
	content.Objects = []fyne.CanvasObject{t.View(window)}
	content.Refresh()
}