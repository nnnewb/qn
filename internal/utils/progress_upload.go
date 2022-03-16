package utils

import (
	"context"
	"fmt"
	"os"
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

	partsize, err := cmd.Flags().GetInt64("partsize")
	if err != nil {
		return err
	}

	// https://developer.qiniu.com/kodo/1238/go#resume-upload-file
	// 服务端SDK中默认块大小为4MB，支持根据条件设置块大小
	// （要求除最后一块外，其他块大于等于1MB，小于等于1G），
	// 块与块之间可以并发上传，以提高上传效率。
	if partsize < 1024*1024*1024 || partsize >= 1024*1024*1024*1024 {
		return fmt.Errorf("分片大小 %d 必须大于等于 1MiB, 小于 1GiB", partsize)
	}

	progress := pb.
		New64(info.Size()).
		SetRefreshRate(time.Second).
		SetWidth(79).
		Set(pb.Bytes, true).
		Start()
	err = uploader.Put(context.Background(), &ret, token, key, file, info.Size(), &storage.RputV2Extra{
		PartSize: partsize,
		Notify: func(partNumber int64, ret *storage.UploadPartsRet) {
			if progress.Current()+partsize > info.Size() {
				progress.Add64(info.Size() - progress.Current())
			} else {
				progress.Add64(partsize)
			}
		},
		NotifyErr: func(partNumber int64, err error) {
			progress.SetErr(err)
		},
	})
	progress.Finish()
	if err != nil {
		return errors.Wrap(err, "putfile fail")
	}

	return nil
}
