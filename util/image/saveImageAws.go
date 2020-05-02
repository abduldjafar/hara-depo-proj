package image

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"hara-depo-proj/config"
	"log"
	"net/http"
)

func connectAWS() *session.Session {
	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(baseConfig.Aws.Region),
		})
	if err != nil {
		panic(err)
	}
	return sess
}

func UploadPhotoAws(r *http.Request, file, kodeuser, jenis, path string) string {
	sess := connectAWS()

	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	filE, header, err := r.FormFile(file)
	if err != nil {
		// Do your error handling here
		log.Println(err.Error())
	}
	defer filE.Close()

	filename := header.Filename

	uploader := s3manager.NewUploader(sess)

	data, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(baseConfig.Aws.Bucketname),                            // Bucket to be used
		Key:    aws.String(kodeuser + "/" + jenis + "/" + path + "/" + filename), // Name of the file to be saved
		Body:   filE,                                                             // File
	})
	if err != nil {
		// Do your error handling here
		log.Println(err.Error())
	}

	return data.Location

}
