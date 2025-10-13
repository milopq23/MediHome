package service

import (
	"context"

	// "log"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

type CloudinaryService interface {
	UploadFile(file multipart.File, filename string, folder string) (string, error)
		HandelRequest(c *gin.Context) (string, error)
	// UploadMultiFileHandleRequest(r *http.Request) ([]string, error)
	UploadMultiFileHandleRequestFromGin(c *gin.Context) ([]string, error)
}

type cloudinaryService struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinaryService(cld *cloudinary.Cloudinary) CloudinaryService {
	return &cloudinaryService{cld: cld}
}

func (s *cloudinaryService) HandelRequest(c *gin.Context) (string, error) {

	folder := c.PostForm("folder")
	if folder == "" {
		folder = "default"
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return "", err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	return s.UploadFile(file, fileHeader.Filename, folder)
}

func (s *cloudinaryService) UploadFile(file multipart.File, filename string, folder string) (string, error) {
	uploadResult, err := s.cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		PublicID: filename,
		Folder:   folder,
	})
	if err != nil {
		return "", err
	}
	return uploadResult.SecureURL, nil
}

func (s *cloudinaryService) UploadMultiFileHandleRequestFromGin(c *gin.Context) ([]string, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}

	files := form.File["files"]

	folder := c.PostForm("folder")
	if folder == "" {
		folder = "default"
	}

	var urls []string
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			continue
		}

		url, err := s.UploadFile(file, fileHeader.Filename, folder)
		file.Close()

		if err != nil {

			continue
		}

		urls = append(urls, url)
	}

	return urls, nil
}
