package bigbluebutton

import (
	"regexp"
	"strconv"

	pdfapi "github.com/pdfcpu/pdfcpu/pkg/api"
	pdf "github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

const SizeRegEx = "[+-]?([0-9]*[.])?[0-9]+"

type PDFReader struct {
	info         []string
	e            error
	pageSizeInfo string
}

func (p *PDFReader) readSize(file string) ([]string, error) {
	p.info, p.e = pdfapi.InfoFile(file, pdf.NewDefaultConfiguration())
	p.pageSizeInfo = p.info[2]
	return p.info, p.e
}

func (p *PDFReader) getPageWidth() float64 {
	r := regexp.MustCompile(SizeRegEx)
	width := 0.00
	matches := r.FindAllString(p.pageSizeInfo, -1)
	width, _ = strconv.ParseFloat(matches[0], 64)

	return width
}

func (p *PDFReader) getPageHeight() float64 {
	r := regexp.MustCompile(SizeRegEx)
	heigth := 0.00
	matches := r.FindAllString(p.pageSizeInfo, -1)
	heigth, _ = strconv.ParseFloat(matches[1], 64)
	return heigth
}
