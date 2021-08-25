package utils

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/cheggaaa/pb/v3"
	"github.com/pkg/errors"
)

func DownloadWithProgress(url string, path string) error {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return errors.Wrap(err, "get url fail")
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "read response body fail")
		}
		return errors.Errorf("status code %d, %s", resp.StatusCode, string(content))
	}

	bar := pb.Full.Start64(resp.ContentLength)
	defer bar.Finish()

	barReader := bar.NewProxyReader(resp.Body)
	defer barReader.Close()

	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return errors.Wrapf(err, "open file fail")
	}

	_, err = io.Copy(file, barReader)
	if err != nil {
		return errors.Wrap(err, "read response fail")
	}

	bar.Finish()
	return nil
}
