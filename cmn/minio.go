package cmn

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	Endpoint string
	Username string
	Password string
	Bucket   string
}

func (m *Minio) Upload(localPathFile string, minioObjectName string) error {
	ctx := context.Background()

	// 初始化
	minioClient, err := minio.New(m.Endpoint, &minio.Options{Creds: credentials.NewStaticV4(m.Username, m.Password, ""), Secure: false})
	if err != nil {
		return err
	}

	// 检查创建Bucket
	exists, err := minioClient.BucketExists(ctx, m.Bucket)
	if !(err == nil && exists) {
		err = minioClient.MakeBucket(ctx, m.Bucket, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}

	// 上传
	_, err = minioClient.FPutObject(ctx, m.Bucket, minioObjectName, localPathFile, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (m *Minio) Download(minioObjectName string, localPathFile string) error {
	ctx := context.Background()

	// 初始化
	minioClient, err := minio.New(m.Endpoint, &minio.Options{Creds: credentials.NewStaticV4(m.Username, m.Password, ""), Secure: false})
	if err != nil {
		return err
	}

	// 下载
	err = minioClient.FGetObject(ctx, m.Bucket, minioObjectName, localPathFile, minio.GetObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}
