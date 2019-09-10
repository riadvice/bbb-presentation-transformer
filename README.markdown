sudo add-apt-repository ppa:longsleep/golang-backports

sudo apt-get update

sudo apt-get install golang-go

go get github.com/jung-kurt/gofpdf

go get github.com/jung-kurt/gofpdf/contrib/gofpdi

go get github.com/pdfcpu/pdfcpu

curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

dep ensure -update
