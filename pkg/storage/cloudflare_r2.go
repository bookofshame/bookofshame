package storage

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsCfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/bookofshame/bookofshame/pkg/config"
	"github.com/bookofshame/bookofshame/pkg/constants"
	"github.com/bookofshame/bookofshame/pkg/logging"
	"go.uber.org/zap"
)

type CloudflareR2 struct {
	ctx        context.Context
	logger     *zap.SugaredLogger
	bucketName string
	client     *s3.Client
}

func NewCloudflareR2(ctx context.Context, cfg config.Config) *CloudflareR2 {
	logger := logging.FromContext(ctx)
	t, err := awsCfg.LoadDefaultConfig(ctx,
		awsCfg.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.R2AccessKey, cfg.R2SecretAccessKey, "")),
		awsCfg.WithRegion("auto"),
	)

	if err != nil {
		logger.Fatal(err)
	}

	client := s3.NewFromConfig(t, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", cfg.R2AccountId))
	})

	return &CloudflareR2{
		ctx:        ctx,
		logger:     logger,
		bucketName: cfg.R2BucketName,
		client:     client,
	}
}

func (r CloudflareR2) Upload(file io.Reader, fileName string) error {
	_, err := r.client.PutObject(r.ctx, &s3.PutObjectInput{
		Bucket: aws.String(r.bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})

	if err != nil {
		r.logger.Error("Could not upload file (name: %s). Error: %w", fileName, err)
		return fmt.Errorf("failed to upload file (name: %s)", fileName)
	}

	r.logger.Debugf("Uploaded file (name: %s).", fileName)

	return nil
}

func (r CloudflareR2) UploadLarge(file io.Reader, fileName string) error {
	uploader := manager.NewUploader(r.client, func(u *manager.Uploader) {
		u.PartSize = constants.PartSize
	})

	_, err := uploader.Upload(r.ctx, &s3.PutObjectInput{
		Bucket: aws.String(r.bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})

	if err != nil {
		r.logger.Error("Could not upload file (name: %s). Error: %w", fileName, err)
		return fmt.Errorf("failed to upload file (name: %s)", fileName)
	}

	r.logger.Debugf("Uploaded large file (name: %s).", fileName)

	return nil
}

func (r CloudflareR2) Download(fileName string) ([]byte, error) {
	result, err := r.client.GetObject(r.ctx, &s3.GetObjectInput{
		Bucket: aws.String(r.bucketName),
		Key:    aws.String(fileName),
	})

	if err != nil {
		r.logger.Errorf("Could not download file (name: %s). Error: %w", fileName, err)
		return nil, fmt.Errorf("could not download file (name: %s)", fileName)
	}

	defer result.Body.Close()

	buffer, err := io.ReadAll(result.Body)

	if err != nil {
		r.logger.Errorf("Could not read file (name: %s). Error: %w", fileName, err)
		return nil, fmt.Errorf("could not read file (name: %s)", fileName)
	}

	r.logger.Debugf("Downloaded file (name: %s).", fileName)

	return buffer, nil
}

func (r CloudflareR2) DownloadLarge(fileName string) ([]byte, error) {
	downloader := manager.NewDownloader(r.client, func(u *manager.Downloader) {
		u.PartSize = constants.PartSize
	})

	buffer := manager.NewWriteAtBuffer([]byte{})
	_, err := downloader.Download(r.ctx, buffer, &s3.GetObjectInput{
		Bucket: aws.String(r.bucketName),
		Key:    aws.String(fileName),
	})

	if err != nil {
		r.logger.Errorf("Could not download large file (name: %s). Error: %w", fileName, err)
		return nil, fmt.Errorf("could not download large file named: %s", fileName)
	}

	r.logger.Debugf("Downloaded large file (name: %s).", fileName)

	return buffer.Bytes(), err
}

func (r CloudflareR2) Delete(fileNames []string) error {
	var objectIds []types.ObjectIdentifier
	for _, key := range fileNames {
		objectIds = append(objectIds, types.ObjectIdentifier{Key: aws.String(key)})
	}

	result, err := r.client.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
		Bucket: aws.String(r.bucketName),
		Delete: &types.Delete{Objects: objectIds},
	})

	if err != nil {
		if len(fileNames) == 1 {
			r.logger.Errorf("Couldn't delete file name(%s). Error: %w", r.bucketName, err)
			return fmt.Errorf("couldn't delete file (name: %s)", r.bucketName)
		}

		r.logger.Errorf("Couldn't delete files. Error: %w", err)
		return fmt.Errorf("couldn't delete files")
	}

	r.logger.Debugf("Deleted %d files", len(result.Deleted))

	return nil
}
