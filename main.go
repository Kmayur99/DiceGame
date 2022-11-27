package main

import (
	img "./ImgByteCode"
	mc "./maincontent"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.LightTheme())
	myWindow := myApp.NewWindow("ROLLING WITH DICE")
	mycanvas := myWindow.Canvas()
	hello := widget.NewLabel("**WELLENFANGT CREATIONS PRESENT**")
	left := widget.NewButtonWithIcon("", theme.NewThemedResource(img.ResourceLeft1aPng, nil), func() {})
	right := widget.NewButtonWithIcon("", theme.NewThemedResource(img.ResourceRightPng, nil), func() {})
	exit := widget.NewButtonWithIcon("EXIT", theme.NewThemedResource(img.ResourceExitPng, nil), func() { myApp.Quit() })
	start := widget.NewButtonWithIcon("START", theme.NewThemedResource(img.ResourceD6MinPng, nil), func() {
		go mc.Showcontent(mycanvas)
	})
	history := widget.NewButtonWithIcon("PLAYER $ HISTORY", theme.NewThemedResource(img.ResourceHistoryPng, nil), func() {})
	image := canvas.NewImageFromResource(img.Resource12Png)
	image.FillMode = canvas.ImageFillOriginal
	myWindow.SetContent(widget.NewVBox(
		fyne.NewContainerWithLayout(layout.NewGridLayout(1), hello),
		fyne.NewContainerWithLayout(layout.NewGridLayout(1), start),
		fyne.NewContainerWithLayout(layout.NewGridLayout(1), history),
		fyne.NewContainerWithLayout(layout.NewGridLayout(1), image),
		fyne.NewContainerWithLayout(layout.NewGridLayout(4), left, right, new(fyne.Container), exit)),
	)
	myWindow.Resize(fyne.NewSize(100, 100))
	myWindow.ShowAndRun()
}
