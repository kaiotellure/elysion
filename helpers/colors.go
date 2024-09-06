package helpers

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"net/http"

	"github.com/cascax/colorthief-go"
)

func GetImagePrimaryColorFromURL(url string) (string, error) {

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)
	if err != nil {
		return "", err
	}

	primary, err := colorthief.GetColor(img)
	if err != nil {
		return "", err
	}

	converted, _ := color.NRGBAModel.Convert(primary).(color.NRGBA)
	return fmt.Sprintf("#%02x%02x%02x", converted.R, converted.G, converted.B), nil
}
