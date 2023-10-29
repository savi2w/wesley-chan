package stringutil

import (
	"net/url"

	"github.com/savi2w/wesley-chan/config"
)

func GetFileURL(cfg *config.Config, fileKey string) string {
	return "https://" + cfg.AWSConfig.FileBucketName + ".s3.amazonaws.com/" + url.PathEscape(fileKey)
}
