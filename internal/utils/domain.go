package utils

import (
	"github.com/pkg/errors"
	"github.com/qiniu/go-sdk/v7/storage"
)

func GetBucketDownloadDomain(bucket string, mgr *storage.BucketManager) (string, error) {
	domains, err := mgr.ListBucketDomains(bucket)
	if err != nil {
		return "", errors.Wrap(err, "list bucket domains failed")
	}

	var downloadDomain string
	for _, domain := range domains {
		downloadDomain = domain.Domain
	}

	if mgr.Cfg.UseHTTPS {
		downloadDomain = "https://" + downloadDomain
	} else {
		downloadDomain = "http://" + downloadDomain
	}

	return downloadDomain, nil
}
