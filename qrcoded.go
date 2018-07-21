package main


import (
    "fmt"
    "io/ioutil"
)

const permission = 0644

func main() {
	fmt.Println("Hello qr code")

    qrcode := GenerateQRCode("555-2368")
    ioutil.WriteFile("qrcode.png", qrcode, permission)
}

func GenerateQRCode(code string) []byte {
    return nil
}
