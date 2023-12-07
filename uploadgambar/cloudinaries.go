package uploadgambar

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"par/config"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/unidoc/unipdf/v3/model"
)

type Uploads interface {
	Upload(file *multipart.FileHeader) (string, error)
	Destroy(publicID string) error
}
type cloudUpload struct {
	clds *cloudinary.Cloudinary
}

func NewCloud(cfg *config.AppConfig) Uploads {
	clds, err := cloudinary.NewFromParams(cfg.CLOUDINARY_CLOUD_NAME, cfg.CLOUDINARY_API_KEY, cfg.CLOUDINARY_API_SECRET)

	if err != nil {
		log.Println("init cloudinary gagal", err)
		return nil
	}

	return &cloudUpload{clds: clds}
}

// Upload implements Uploads.
func (cl *cloudUpload) Upload(file *multipart.FileHeader) (string, error) {
	src, errfile := file.Open()

	if errfile != nil {
		return "", errfile
	}
	defer src.Close()
	// buffer := make([]byte, 512)
	// _, errread := src.Read(buffer)

	// if errread != nil {
	// 	return "", errread
	// }

	// filetype := http.DetectContentType(buffer)
	// if filetype != "application/pdf" {
	// 	return " ", errors.New("file bukan pdf")
	// }
	_, errpdf := model.NewPdfReader(src)
	if errpdf != nil {
		return "", errors.New("file anda bukan pdf")
	}
	src.Seek(0, 0)

	publicID := fmt.Sprintf("%d-%s", int(file.Size), time.Now().Format("20060102-150405"))
	uploadParams := uploader.UploadParams{
		PublicID:     publicID,
		Folder:       os.Getenv("CLOUDINARY_UPLOAD_FOLDER"),
		ResourceType: "raw",
		Format:       "pdf",
	}

	uploadResult, err := cl.clds.Upload.Upload(context.Background(), src, uploadParams)
	if err != nil {
		return "", errors.New("tidak berhasil upload,check your connection")
	}

	return uploadResult.SecureURL, nil
}

// Destroy implements Uploads.
func (*cloudUpload) Destroy(publicID string) error {
	panic("unimplemented")
}

func GetPublicID(secureURL string) string {
	// Proses filter Public ID dari SecureURL(avatar)
	urls := strings.Split(secureURL, "/")
	urls = urls[len(urls)-2:]                               // array [file, random_name.extension]
	noExtension := strings.Split(urls[len(urls)-1], ".")[0] // remove ".extension", result "random_name"
	urls = append(urls[:1], noExtension)                    // new array [file, random_name]
	publicID := strings.Join(urls, "/")                     // "file/random_name"

	return publicID
}
