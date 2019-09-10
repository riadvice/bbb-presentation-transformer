package bigbluebutton

import (
	"fmt"
	"strings"
	"testing"
)

func TestPresentationSize(t *testing.T) {
	var reader = PDFReader{}
	var info, _ = reader.readSize("default.pdf")

	if len(info) != 20 {
		t.Errorf("Wrong Info Array size \nGot => %d", len(info))
	} else {
		t.Logf("Found %d information", len(info))
	}

	table := []struct {
		output, expected, name string
	}{
		{strings.Trim(info[2], " "), "Page size: 1920.00 x 1080.00 points", "Full size info"},
		{fmt.Sprintf("%.2f", reader.getPageWidth()), fmt.Sprintf("%.2f", 1920.00), "PDF Page Width"},
		{fmt.Sprintf("%.2f", reader.getPageHeight()), fmt.Sprintf("%.2f", 1080.00), "PDF Page Height"},
	}

	for _, i := range table {
		if i.output != i.expected {
			t.Errorf("%s :: Output: %s correct: %s", i.name, i.output, i.expected)
		} else {
			t.Logf("%s :: Pass", i.name)
		}
	}

	/*

		if  !=  {
			t.Errorf("Wrong Document Page Height \nGot => %f", reader.getPageHeight())
		}
	*/
}
