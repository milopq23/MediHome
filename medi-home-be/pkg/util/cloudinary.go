package util

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func UltiCloudinary() (*cloudinary.Cloudinary, error){
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	// Kiểm tra biến môi trường có đủ không
	if cloudName == "" || apiKey == "" || apiSecret == "" {
		log.Fatal("Thiếu thông tin cấu hình Cloudinary (env CLOUDINARY_CLOUD_NAME / API_KEY / API_SECRET)")
	}

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		log.Fatalf("Lỗi khi khởi tạo Cloudinary: %v", err)
	}

	return cld, err
}