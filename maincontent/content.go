package maincontent

import (
	"fmt"
	"log"
	"strconv"

	img1 "../ImgByteCode"
	op "../operation"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/Knetic/govaluate"
)

func evaluate(s1 string, w *widget.Label) {
	expression, err := govaluate.NewEvaluableExpression(s1)
	if err == nil {
		result, err := expression.Evaluate(nil)
		if err == nil {
			w.SetText((strconv.FormatFloat(result.(float64), 'f', -1, 64)))
		}
	}

	if err != nil {
		log.Println("Error in calculation", err)
		w.SetText("error")
	}
}

func Showcontent(c fyne.Canvas) {
	dice1 := canvas.NewImageFromResource(img1.Resource12Png)
	dice1.FillMode = canvas.ImageFillOriginal
	dice2 := canvas.NewImageFromResource(img1.Resource12Png)
	dice2.FillMode = canvas.ImageFillOriginal
	gSession := widget.NewEntry()
	gSession.SetPlaceHolder("X")
	button1 := widget.NewButtonWithIcon("ROLL-THE-DICE", theme.NewThemedResource(img1.ResourceBut12MinPng, nil), func() {
		dice1.Resource, _ = op.RolePolygonDice()
		dice2.Resource, _ = op.RolePolygonDice()
		dice1.Refresh()
		dice2.Refresh()
	})

	dice3 := canvas.NewImageFromResource(img1.Resource5aMinPng)
	dice3.FillMode = canvas.ImageFillOriginal
	dice4 := canvas.NewImageFromResource(img1.Resource5aMinPng)
	dice4.FillMode = canvas.ImageFillOriginal
	dice5 := canvas.NewImageFromResource(img1.Resource5aMinPng)
	dice5.FillMode = canvas.ImageFillOriginal
	// Y
	gSessionY := widget.NewEntry()
	gSessionY.SetPlaceHolder("Y")
	// Z
	gSessionZ := widget.NewEntry()
	gSessionZ.SetPlaceHolder("Z")
	button2 := widget.NewButtonWithIcon("ROLL-THE-DICE", theme.NewThemedResource(img1.ResourceButicon1Png, nil), func() {
		_, val1 := op.Dice()
		_, val2 := op.Dice()
		_, val3 := op.Dice()
		dice3.Resource, _ = op.Dice()
		dice4.Resource, _ = op.Dice()
		dice5.Resource, _ = op.Dice()
		dice3.Refresh()
		dice4.Refresh()
		dice5.Refresh()
		gSession.SetText(strconv.Itoa(val1))
		gSessionY.SetText(strconv.Itoa(val2))
		gSessionZ.SetText(strconv.Itoa(val3))
	})

	gSessionAriMticOper1 := widget.NewEntry()
	gSessionAriMticOper1.SetPlaceHolder("OP")

	gSessionAriMticOper2 := widget.NewEntry()
	gSessionAriMticOper2.SetPlaceHolder("OP")

	addition := widget.NewButton("+", func() {
		gSessionAriMticOper1.SetText("+")
		gSessionAriMticOper2.SetText("+")
	})
	subtraction := widget.NewButton("-", func() {
		gSessionAriMticOper1.SetText("-")
		gSessionAriMticOper2.SetText("-")
	})
	multiply := widget.NewButton("*", func() {
		gSessionAriMticOper1.SetText("*")
		gSessionAriMticOper2.SetText("*")
	})
	modulo := widget.NewButton("%", func() {
		gSessionAriMticOper1.SetText("%")
		gSessionAriMticOper2.SetText("%")
	})
	lb1 := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	lb2 := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	lb3 := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	verify := widget.NewButtonWithIcon("Verify", theme.NewThemedResource(img1.ResourceCheckPng, nil), func() {
		ss := gSession.Text + gSessionAriMticOper1.Text + gSessionY.Text + gSessionAriMticOper2.Text + gSessionZ.Text
		evaluate(ss, lb1)

		_, valA := op.RolePolygonDice()
		_, valB := op.RolePolygonDice()
		lb2.SetText(strconv.Itoa(valA * valB))

		i1, err := strconv.Atoi(lb1.Text)
		if err == nil {
			fmt.Println(i1)
		}
		i2, err := strconv.Atoi(lb2.Text)
		if err == nil {
			fmt.Println(i1)
		}

		if i1 == i2 {
			lb3.SetText("SUCCESS")
		} else {
			lb3.SetText("FAIL")
		}

	})

	container := fyne.NewContainerWithLayout(layout.NewGridLayout(5),
		gSession, gSessionAriMticOper1, gSessionY, gSessionAriMticOper2, gSessionZ)
	container1 := fyne.NewContainerWithLayout(layout.NewGridLayout(4),
		addition, subtraction, modulo, multiply)
	container2 := fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		verify)

	grid := fyne.NewContainerWithLayout(layout.NewGridLayout(1), container, container1, container2)

	grid1 := widget.NewHBox(
		button1,
		fyne.NewContainerWithLayout((layout.NewGridWrapLayout(fyne.NewSize(95, 95))), dice1, dice2),
		button2,
		fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(62, 62)), dice3, dice4, dice5),
		grid)
	c.SetContent(fyne.NewContainerWithLayout(layout.NewCenterLayout(), widget.NewVBox(grid1, lb1, lb2, lb3)))
}
