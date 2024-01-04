package controller

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
)

type (
	QRData struct {
		Fullname string `json:"fullname"`
		Folder   string `json:"folder"`
	}
)

func GenerateQR(c *fiber.Ctx) error {
	qrData := &[]QRData{}
	c.BodyParser(qrData)

	for _, v := range *qrData {
		fmt.Println("Name:", v.Fullname)
		folderPath := fmt.Sprintf("./qr/codes/%v", v.Folder)
		CreateDirectory(folderPath)
		GenerateQRCode(v.Fullname, v.Folder)
	}

	return c.JSON(fiber.Map{
		"header":  c.GetRespHeaders(),
		"message": "success",
		"data":    qrData,
	})
}

func GenerateQRCode(name, folder string) error {
	// now := time.Now()
	qrCode, _ := qrcode.New(name, qrcode.Highest)
	fileName := fmt.Sprintf("./qr/codes/%v/%v.png", folder, name)

	err := qrCode.WriteFile(256, fileName)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func GenerateQRWithLogo(c *fiber.Ctx) error {
	now := time.Now()
	// Text or data you want to encode in the QR code.
	qrData := &QRData{}
	c.BodyParser(qrData)

	// Generate the QR code.
	qr, err := qrcode.New(qrData.Fullname, qrcode.Highest)
	if err != nil {
		fmt.Println("qrErr:", err)
		log.Fatal(err)
	}

	// Encode the QR code as a PNG image.
	qrImg := qr.Image(256)

	// Load your logo image.
	logoFile, err := os.Open("assets/logo.png")
	if err != nil {
		fmt.Println("openErr:", err)
		log.Fatal(err)
	}

	// fmt.Println("logoFile:", logoFile)
	defer logoFile.Close()

	logoImg, _, err := image.Decode(logoFile)
	if err != nil {
		fmt.Println("decodErr:", err)
		log.Fatal(err)
	}

	// fmt.Println("logoIMG:", logoImg)
	// Calculate the position to overlay the logo.
	logoSize := 64 // Adjust the size of the logo as needed.
	x := (qrImg.Bounds().Dx() - logoSize) / 2
	y := (qrImg.Bounds().Dy() - logoSize) / 2

	// Create a new image to overlay the QR code and logo.
	finalImg := image.NewRGBA(qrImg.Bounds())
	draw.Draw(finalImg, qrImg.Bounds(), qrImg, image.Point{}, draw.Over)

	// Overlay the logo on the QR code.
	logoRect := image.Rect(x, y, x+logoSize, y+logoSize)

	draw.Draw(finalImg, logoRect, logoImg, image.Point{}, draw.Over)

	// Save the final image to a file.
	fileName := fmt.Sprintf("./qr/logo/%v.png", now.UnixMilli())
	outputFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("creatErr:", err)
		log.Fatal(err)
	}
	defer outputFile.Close()

	if err := png.Encode(outputFile, finalImg); err != nil {
		fmt.Println("encodErr:", err)
		log.Fatal(err)
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    qr,
	})
}

func CreateDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}

func TestSpecialChar(c *fiber.Ctx) error {
	// request := fiber.Map{}
	// if parsErr := c.BodyParser(request); parsErr != nil {
	// 	return c.JSON(fiber.Map{
	// 		"result": parsErr.Error(),
	// 	})
	// }

	param1 := c.Params("param1")
	param2 := c.Params("param2")
	return c.JSON(fiber.Map{
		"result1": param1,
		"result2": param2,
	})
}
