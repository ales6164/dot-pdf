package main

import (
	"log"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"bytes"

	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		if len(url) == 0 {
			body, _ := ioutil.ReadAll(r.Body)
			r.Body.Close()

			buf := ExampleNewPDFGenerator(body)
			buf.WriteTo(w)
		} else {
			buf := ExampleNewPDFGeneratorURL(url)
			buf.WriteTo(w)
		}
	})

	http.Handle("/", router)

	//router.Run(":" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func ExampleNewPDFGenerator(html []byte) *bytes.Buffer {

	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	/*pdfg.MarginTop.Set(0)
	pdfg.MarginBottom.Set(0)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)*/

	pdfg.OutputFile = ""

	// Add one page from an URL
	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(html)))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	return pdfg.Buffer()
}

func ExampleNewPDFGeneratorURL(url string) *bytes.Buffer {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	/*pdfg.MarginTop.Set(0)
	pdfg.MarginBottom.Set(0)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)*/

	pdfg.OutputFile = ""

	// Add one page from an URL
	pdfg.AddPage(wkhtmltopdf.NewPage(url))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	return pdfg.Buffer()
}
