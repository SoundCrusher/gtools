package data

import (
	"fyne.io/fyne/v2"
	"gtools/app/feature"
	"gtools/app/feature/convert"
	"gtools/app/feature/curl"
)

type TreeNode struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
}

var (
	TreeNodeMetadata = map[string][]string{
		"":            {"Curl", "ConvertOctalUtf8", "More"},
		"collections": {"list", "table", "tree"},
		"containers":  {"apptabs", "border", "box", "center", "doctabs", "grid", "scroll", "split"},
		"widgets":     {"accordion", "button", "card", "entry", "form", "input", "progress", "text", "toolbar"},
	}

	// TreeNodeMappings defines the metadata for each tutorial
	TreeNodeMappings = map[string]TreeNode{
		"Curl":             {"Curl", "格式化你的 cURL 命令", curl.NewFormatter().NewBorder},
		"ConvertOctalUtf8": {"ConvertOctalUtf8", "八进制字符串转换", convert.NewOctalNumberConvert().NewBorder},
		"More":             {"More", "See more...", feature.MakeBoxLayout},
	}
)
