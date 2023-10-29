package stringutil

import (
	"net/url"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/consts"
	"github.com/savi2w/wesley-chan/model"
)

func GetFileName(fileName string) string {
	fileName = strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))

	// Replace any sequence of non-alphanumeric characters with a single underscore
	re := regexp.MustCompile("[^a-zA-Z0-9_]+")
	fileName = re.ReplaceAllString(fileName, "_")

	return fileName
}

func GetFileExt(fileName string) string {
	ext := filepath.Ext(fileName)
	runes := []rune(ext)

	if len(runes) >= consts.MaxFileExtensionSize {
		return string(runes[:consts.MaxFileExtensionSize])
	}

	return ext
}

func GetFileURL(cfg *config.Config, file *model.File) string {
	return "https://" + cfg.AWSConfig.FileBucketName + ".s3.amazonaws.com/" + url.PathEscape(file.ID) + "#" + url.PathEscape(file.OriginalFileName)
}
