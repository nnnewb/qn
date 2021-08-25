package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

func Upload(bucket, key, path string, credential *auth.Credentials) error {
	ret := storage.PutRet{}
	policy := storage.PutPolicy{
		Scope:   fmt.Sprintf("%s:%s", bucket, key),
		Expires: uint64(time.Now().Unix()) + uint64(time.Hour.Seconds()),
	}
	token := policy.UploadToken(credential)
	uploader := storage.NewFormUploader(&storage.Config{})

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return errors.Wrap(err, "open upload file fail")
	}

	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return errors.Wrap(err, "check upload file stat fail")
	}

	err = uploader.Put(context.TODO(), &ret, token, key, file, info.Size(), &storage.PutExtra{})
	if err != nil {
		return errors.Wrap(err, "putfile fail")
	}

	return nil
}
