package stringutil

import (
	"net/url"
	"path/filepath"

	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/consts"
)

func GetFileExt(fileName string) string {
	ext := filepath.Ext(fileName)
	runes := []rune(ext)

	if len(runes) >= consts.MaxFileExtensionSize {
		return string(runes[:consts.MaxFileExtensionSize])
	}

	return ext
}

func GetFileURL(cfg *config.Config, fileID string) string {
	return "https://" + cfg.AWSConfig.FileBucketName + ".s3.amazonaws.com/" + url.PathEscape(fileID)
}
