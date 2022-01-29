package curl

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"gtools/app/vars"
)

type Formatter struct {
	fyne.Window

	data    *Curl
	content *widget.Entry
	button  *widget.Button
}

type Curl struct {
	req         string
	options     []string
	longOptions []string
	header      []string
	data        string
}

func NewFormatter() *Formatter {
	// 只初始化组件，等 app 运行后，才做其他处理，不然会 panic
	f := &Formatter{}
	// 内容框
	f.content = widget.NewMultiLineEntry()
	// 按钮
	f.button = widget.NewButton("Click me.", f.clickFunc)

	return f
}

func (f *Formatter) clickFunc() {
	log.Printf("点击按钮, 输入的文字为: \n%s", f.content.Text)
	curl, err := matchCurl(f.content.Text)
	if err != nil {
		information := dialog.NewInformation("提示", err.Error(), f.Window)
		information.Show()
	} else {
		f.content.Text = curl
		f.content.Refresh()
	}
}

func (f *Formatter) NewBorder(win fyne.Window) fyne.CanvasObject {
	f.Window = win

	f.content.Wrapping = fyne.TextWrapWord
	f.content.SetPlaceHolder("粘贴你的 curl 命令到这里...")

	top := container.New(layout.NewGridWrapLayout(fyne.NewSize(720, 370)), f.content)
	bottom := container.New(layout.NewCenterLayout(), f.button)
	return container.NewVBox(top, bottom)
}

func (c *Curl) toString() string {
	options := strings.Join(c.options, " ")
	longOptions := strings.Join(c.longOptions, "\\\n\t\t")
	headers := strings.Join(c.header, "\\\n\t\t")
	return fmt.Sprintf("curl %s %s \\\n\t\t%s \\\n\t\t%s", options, c.req, headers, longOptions)
}

func matchCurl(curlStr string) (string, error) {
	index := strings.Index(strings.ToLower(curlStr), "curl ")
	if index < 0 {
		return "", vars.MatchError
	}
	log.Println("开始位置: ", index)
	curlStr = curlStr[index+4:]

	curl := &Curl{
		req:     "",
		options: make([]string, 0),
		header:  make([]string, 0),
		data:    "",
	}
	//解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile(`(-H|-h)(\s{0,10})(".*"|'.*')`)
	if reg1 == nil {
		fmt.Println("regexp err")
		return "", vars.MatchError
	}
	//根据规则提取关键信息
	for _, n := range reg1.FindAllString(curlStr, -1) {
		curlStr = strings.Replace(curlStr, n, "", 1)
		i := strings.Index(n, "-H ")
		y := strings.Index(n, "-h ")
		begin := math.Max(float64(i), float64(y))
		n = n[int(begin):]
		curl.header = append(curl.header, n)
		fmt.Println("header = ", n)
	}

	// --options
	reg3 := regexp.MustCompile(`\s+--\w+(-\w+){0,3}\s*(".*"|'.*'|\w+\s)?`)
	if reg3 == nil {
		fmt.Println("regexp err")
		return "", vars.MatchError
	}
	for _, op := range reg3.FindAllString(curlStr, -1) {
		curlStr = strings.Replace(curlStr, op, "", 1)
		op = op[strings.Index(op, "--"):]
		curl.longOptions = append(curl.longOptions, op)
		fmt.Println("--参数: ", op)
	}

	// -op
	reg4 := regexp.MustCompile(`\s+-\w+\s*(".*"|'.*'|\w+\s)?`)
	if reg4 == nil {
		fmt.Println("regexp err")
		return "", vars.MatchError
	}
	for _, op := range reg4.FindAllString(curlStr, -1) {
		curlStr = strings.Replace(curlStr, op, "", 1)
		op = op[strings.Index(op, "-"):]
		curl.options = append(curl.options, op)
		fmt.Println("-参数: ", op)
	}
	fmt.Println("剩下: ", curlStr)

	req := regexp.MustCompile("[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]")
	if req == nil {
		fmt.Println("regexp err")
		return "", vars.MatchError
	}
	curl.req = req.FindAllString(curlStr, -1)[0]
	// 取 url
	return curl.toString(), nil
}
