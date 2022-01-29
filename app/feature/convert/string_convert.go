package convert

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"regexp"
	"strconv"
)

type OctalNumberConvert struct {
	fyne.Window

	encode *widget.Entry
	decode *widget.Entry
	clear  *widget.Button
	trans  *widget.Button
}

func NewOctalNumberConvert() *OctalNumberConvert {
	numberConvert := &OctalNumberConvert{}

	numberConvert.encode = widget.NewMultiLineEntry()
	numberConvert.decode = widget.NewMultiLineEntry()
	numberConvert.trans = widget.NewButton("解码并复制到粘贴板", numberConvert.transFunc)
	numberConvert.clear = widget.NewButton("清空", numberConvert.clearFunc)
	return numberConvert
}

func (o *OctalNumberConvert) clearFunc() {
	o.encode.Text = ""
	o.decode.Text = ""
	o.encode.Refresh()
	o.decode.Refresh()
	o.encode.FocusGained()
}

func (o *OctalNumberConvert) transFunc() {
	o.decode.Text = convert(o.encode.Text)
	o.decode.Refresh()
	o.decode.FocusGained()
	o.Clipboard().SetContent(o.decode.Text)
}

func (o *OctalNumberConvert) reverseTransFunc() {

}

func (o *OctalNumberConvert) NewBorder(win fyne.Window) fyne.CanvasObject {
	o.Window = win

	o.encode.Wrapping = fyne.TextWrapWord
	o.encode.SetPlaceHolder("粘贴你的八进制字符串到这里...")

	o.decode.Wrapping = fyne.TextWrapWord
	o.decode.SetPlaceHolder("粘贴字符串到这里...")

	center := container.New(layout.NewGridWrapLayout(fyne.NewSize(720, 180)), o.encode, o.decode)
	bottom := container.New(layout.NewHBoxLayout(), o.trans, o.clear)
	return container.NewBorder(nil, bottom, nil, nil, center)
}

func convert(in string) string {
	s := []byte(in)
	reg := regexp.MustCompile(`\\[0-7]{3}`)

	out := reg.ReplaceAllFunc(s,
		func(b []byte) []byte {
			i, _ := strconv.ParseInt(string(b[1:]), 8, 0)
			return []byte{byte(i)}
		})
	return string(out)
}
