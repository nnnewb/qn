package utils

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/pkg/errors"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/cobra"
)

func Upload(cmd *cobra.Command, bucket, key, path string, credential *auth.Credentials) error {
	ret := storage.PutRet{}
	policy := storage.PutPolicy{
		Scope:   fmt.Sprintf("%s:%s", bucket, key),
		Expires: uint64(time.Now().Unix()) + uint64(time.Hour.Seconds()),
	}
	token := policy.UploadToken(credential)
	uploader := storage.NewResumeUploaderV2(&storage.Config{})

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return errors.Wrap(err, "open upload file fail")
	}

	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return errors.Wrap(err, "check upload file stat fail")
	}

	mutex := sync.Mutex{}
	partsize, err := cmd.Flags().GetInt64("partsize")
	if err != nil {
		return err
	}

	total := info.Size()
	progress := pb.
		New64(total).
		SetRefreshRate(time.Second).
		SetWidth(79).
		Set(pb.Bytes, true).
		Start()
	err = uploader.Put(context.Background(), &ret, token, key, file, total, &storage.RputV2Extra{
		PartSize: partsize,
		Notify: func(partNumber int64, ret *storage.UploadPartsRet) {
			mutex.Lock()
			defer mutex.Unlock()
			if total-partsize < 0 {
				progress.Add64(total)
			} else {
				progress.Add64(partsize)
			}
			total -= partsize
		},
		NotifyErr: func(partNumber int64, err error) {
			mutex.Lock()
			defer mutex.Unlock()
			progress.SetErr(err)
		},
	})
	progress.Finish()
	if err != nil {
		return errors.Wrap(err, "putfile fail")
	}

	return nil
}