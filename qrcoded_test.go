package main

import (
    "testing"
    "bytes"
    "image/png"
    "errors"
)

type ErrorWriter struct{}

func (e *ErrorWriter) Write(b []byte) (int, error) {
    return 0, errors.New("Expected error")
}

func TestGenerateQRCodePropagatesErrors(t *testing.T) {
    w := new(ErrorWriter)
    err := GenerateQRCode(w, "555-32432")

    if err == nil || err.Error() != "Expected error" {
        t.Errorf("Error not propagated correctly, go %v", err)
    }
}

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
