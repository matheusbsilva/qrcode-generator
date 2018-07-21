package main

import (
    "testing"
    "bytes"
    "image/png"
)

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
    buffer := new(bytes.Buffer)
    GenerateQRCode(buffer, "555-0940")

    if buffer.Len() == 0 {
        t.Errorf("No QRCode generated")
    }

    _, err := png.Decode(buffer)

    if err != nil {
        t.Errorf("Generated QRCode is not a PNG: %s", err)
    }
}
