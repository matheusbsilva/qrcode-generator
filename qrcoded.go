package main


import (
    "fmt"
    "image"
    "image/png"
    "os"
    "io"
)

const permission = 0644

func main() {
	fmt.Println("Hello qr code")

    file, _ := os.Create("qrcode.png")
    defer file.Close()

    GenerateQRCode(file, "555-2368")
}

func GenerateQRCode(w io.Writer, code string) error {
    img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
    return png.Encode(w, img)
}
