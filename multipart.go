package calcapi

import (
	calc "goa-ddb/gen/calc"
	"mime/multipart"
)

// CalcAddresumeDecoderFunc implements the multipart decoder for service "calc"
// endpoint "addresume". The decoder must populate the argument p after
// encoding.
func CalcAddresumeDecoderFunc(mr *multipart.Reader, p *[]*calc.Resume) error {
	// Add multipart request decoder logic here
	return nil
}

// CalcAddresumeEncoderFunc implements the multipart encoder for service "calc"
// endpoint "addresume".
func CalcAddresumeEncoderFunc(mw *multipart.Writer, p []*calc.Resume) error {
	// Add multipart request encoder logic here
	return nil
}
