package main


import (
    "fmt"
    "io/ioutil"
    "image"
    "image/png"
    "bytes"
)

const permission = 0644

func main() {
	fmt.Println("Hello qr code")

    qrcode := GenerateQRCode("555-2368")
    ioutil.WriteFile("qrcode.png", qrcode, permission)
}

func GenerateQRCode(code string) []byte {
    img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
    buf := new(bytes.Buffer)
    _ = png.Encode(buf, img)

    return buf.Bytes()
}
