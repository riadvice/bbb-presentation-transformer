// bbb-presentation-transformer project main.go
package main

import (
	"fmt"

	"bytes"
	"io"
	"io/ioutil"

	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/gofpdi"
)

func main() {
	// for testing purposes, get an arbitrary template pdf as stream
	rs, _ := getTemplatePdf()

	// create a new Importer instance
	imp := gofpdi.NewImporter()

	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr: "pt",
		Size:    gofpdf.SizeType{Wd: 1920, Ht: 1080},
	})

	tpl := imp.ImportPageFromStream(pdf, &rs, 1, "/MediaBox")
	pageSizes := imp.GetPageSizes()
	nrPages := len(imp.GetPageSizes())
	fmt.Println("Pages =>", nrPages)

	// add all pages from template pdf
	for i := 1; i <= nrPages; i++ {
		pdf.AddPage()
		if i > 1 {
			tpl = imp.ImportPageFromStream(pdf, &rs, i, "/MediaBox")
		}
		imp.UseImportedTemplate(pdf, tpl, 0, 0, pageSizes[i]["/MediaBox"]["w"], pageSizes[i]["/MediaBox"]["h"])
	}
	// output
	//fileStr := PdfFile(baseStr + ".pdf")
	//	example.Filename("output")
	pdf.OutputFileAndClose("output.pdf")
	//	example.Summary(err, fileStr)

	fmt.Println("PDF Ready!")
}

func getTemplatePdf() (io.ReadSeeker, error) {
	data, err := ioutil.ReadFile("default.pdf")
	return bytes.NewReader(data), err
}
