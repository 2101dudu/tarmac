package pdf

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"regexp"
	"sync"
)

var pdfGenMutex sync.Mutex

func (p *PDFData) GeneratePDF() (string, error) {
	tmpFile, err := os.CreateTemp("pdf/template", "input_*.typ")
	if err != nil {
		return "", err
	}

	tmpPath := tmpFile.Name()
	defer os.Remove(tmpPath)
	defer tmpFile.Close()

	p.fillPDF(tmpFile)
	return p.compilePDF(tmpPath)
}

func (p *PDFData) compilePDF(inputPath string) (string, error) {
	re := regexp.MustCompile(` `)
	output := fmt.Sprintf("out/pdf/orcamento_%s_%d.pdf", re.ReplaceAllString(p.GeneralInfo.CustomerName, "_"), p.GeneralInfo.QuotationNumber)

	out, err := exec.Command("typst", "compile", inputPath, output).CombinedOutput()

	if err != nil {
		slog.Error("Typst Error", "Output", string(out))
		return "", fmt.Errorf("Error executing typst compile: %v, Output: %s", err, string(out))
	}
	return output, nil
}
