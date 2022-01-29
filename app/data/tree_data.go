package data

import (
	"fyne.io/fyne/v2"
	"gtools/app/feature"
)

type TreeNode struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
}

var (
	TreeNodeMetadata = map[string][]string{
		"":            {"Curl", "More"},
		"collections": {"list", "table", "tree"},
		"containers":  {"apptabs", "border", "box", "center", "doctabs", "grid", "scroll", "split"},
		"widgets":     {"accordion", "button", "card", "entry", "form", "input", "progress", "text", "toolbar"},
	}

	// TreeNodeMappings defines the metadata for each tutorial
	TreeNodeMappings = map[string]TreeNode{
		"Curl": {"Curl", "格式化你的 cURL 命令", feature.MakeCurlBorder},
		"More": {"More", "See more...", feature.MakeBoxLayout},
	}
)
