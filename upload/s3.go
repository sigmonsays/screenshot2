package upload

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sigmonsays/screenshot2/config"
	"github.com/sigmonsays/screenshot2/core"
)

type S3 struct {
}

func (me *S3) Upload(cfg *config.AppConfig, shortname *core.Shortname) error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-1"),
		Credentials: credentials.NewStaticCredentials(cfg.Upload.AccessKey, cfg.Upload.SecretKey, ""),
	})
	objectName := shortname.Value + ".jpg" // name in remote storage
	uploader := s3manager.NewUploader(sess)
	f, err := os.Open(shortname.LocalFile)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:          aws.String(cfg.Upload.Bucket),
		Key:             aws.String(objectName),
		ACL:             aws.String("public-read"),
		ContentEncoding: aws.String("image/jpg"),
		Body:            f,
	})
	if err != nil {
		return err
	}

	shortname.Url = fmt.Sprintf("http://%s.s3.amazonaws.com/%s", cfg.Upload.Bucket, objectName)

	log.Tracef("successfully uploaded %s bucket:%s", objectName, cfg.Upload.Bucket)

	return nil
}
