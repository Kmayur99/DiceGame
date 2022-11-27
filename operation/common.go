package operation

import (
	"math/rand"
	"time"

	img "../ImgByteCode"
	"fyne.io/fyne"
)

func Dice() (*fyne.StaticResource, int) {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 6
	rolpdice := rand.Intn(max-min+1) + min

	switch rolpdice {
	case 1:
		return img.ResourceD1MinPng, rolpdice
	case 2:
		return img.ResourceD2MinPng, rolpdice
	case 3:
		return img.ResourceD3MinPng, rolpdice
	case 4:
		return img.ResourceD4MinPng, rolpdice
	case 5:
		return img.ResourceD5MinPng, rolpdice
	case 6:
		return img.ResourceD6MinPng, rolpdice
	}
	/*
		TODO
		Contaier should contain some user friendly message
	*/
	return img.ResourceD1MinPng, 1
}

func RolePolygonDice() (*fyne.StaticResource, int) {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 12
	rolpdice := rand.Intn(max-min+1) + min

	switch rolpdice {
	case 1:
		return img.Resource1Png, rolpdice
	case 2:
		return img.Resource2Png, rolpdice
	case 3:
		return img.Resource3Png, rolpdice
	case 4:
		return img.Resource4Png, rolpdice
	case 5:
		return img.Resource5Png, rolpdice
	case 6:
		return img.Resource6Png, rolpdice
	case 7:
		return img.Resource7Png, rolpdice
	case 8:
		return img.Resource8Png, rolpdice
	case 9:
		return img.Resource9Png, rolpdice
	case 10:
		return img.Resource10Png, rolpdice
	case 11:
		return img.Resource11Png, rolpdice
	case 12:
		return img.Resource12aPng, rolpdice
	}
	/*
		TODO
		Contaier should contain some user friendly message
	*/
	return img.Resource12aPng, 1
}
