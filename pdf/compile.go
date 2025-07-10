package pdf

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sync"
)

var pdfGenMutex sync.Mutex

func (p *PDFData) GeneratePDF() (string, error) {
	pdfGenMutex.Lock()
	defer pdfGenMutex.Unlock()

	filePath := "pdf/template/input.typ"

	f, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	// defer os.Remove(filePath)

	p.fillPDF(f)
	return p.compilePDF()
}

func (p *PDFData) compilePDF() (string, error) {
	re := regexp.MustCompile(` `)
	output := fmt.Sprintf("out/pdf/orcamento_%s_%d.pdf", re.ReplaceAllString(p.GeneralInfo.CustomerName, "_"), p.GeneralInfo.QuotationNumber)
	fmt.Println(output)
	err := exec.Command("typst", "compile", "pdf/template/main.typ", output).Run()
	if err != nil {
		return "", err
	}
	return output, nil
}
