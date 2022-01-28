package vars

import (
	"fmt"
	"fyne.io/fyne/v2"
)

var (
	TopWindow fyne.Window
	Light     *fyne.MenuItem
	Dark      *fyne.MenuItem
)

var (
	MatchError = fmt.Errorf(MatchErr)
)
