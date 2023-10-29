package imageutil

// func IsImage(src io.Reader) (bool, error) {
// 	// Read the first few bytes to identify the file type
// 	header := make([]byte, 261)

// 	_, err := src.Read(header)
// 	if err != nil {
// 		return false, err
// 	}

// 	// Reset reader to the beginning of the file
// 	if seeker, success := src.(io.Seeker); success {
// 		_, err = seeker.Seek(0, io.SeekStart)

// 		if err != nil {
// 			return false, err
// 		}
// 	}

// 	_, _, err = image.Decode(bytes.NewReader(header))
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }
