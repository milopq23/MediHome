package service

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryService interface {
	UploadFile(file multipart.File, filename string) (string, error)
	HandelResquest(r *http.Request) (string, error)
}

type cloudinaryService struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinaryService(cld *cloudinary.Cloudinary) CloudinaryService {
	return &cloudinaryService{cld: cld}
}

func (s *cloudinaryService) HandelResquest(r *http.Request) (string, error){
	err := r.ParseMultipartForm(10<<20)	
	if err != nil {
		return "", err
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()

	return s.UploadFile(file, header.Filename)
}

func (s *cloudinaryService) UploadFile(file multipart.File, filename string) (string, error){
	uploadResult, err := s.cld.Upload.Upload(context.Background(),file,uploader.UploadParams{
		PublicID: filename,
	})
	if err != nil {
		return "", err
	}
	return uploadResult.SecureURL, nil
}




