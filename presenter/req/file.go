package req

import "mime/multipart"

type File struct {
	Header *multipart.FileHeader
}
