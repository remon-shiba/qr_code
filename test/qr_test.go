package test

import (
	"fmt"
	"testing"

	qr "github.com/skip2/go-qrcode"
)

func TestExampleEncode(t *testing.T) {
	if png, err := qr.Encode("Roldan Polintang", qr.Medium, 256); err != nil {
		t.Errorf("Error: %s", err.Error())
	} else {
		fmt.Printf("PNG is %d bytes long", len(png))
	}

	
}
