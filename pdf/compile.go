package main

import (
	"fmt"
	"os"
	"os/exec"
)

func (p *PDFData) GeneratePDF() error {
	filePath := "template/input.typ"

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	defer os.Remove(filePath)

	p.fillPDF(f)
	return p.compilePDF()
}

func (p *PDFData) compilePDF() error {
	output := fmt.Sprintf("out/orcamento_%s_%d.pdf", p.GeneralInfo.CustomerName, p.GeneralInfo.QuotationNumber)
	fmt.Println(output)
	err := exec.Command("typst", "compile", "template/main.typ", output).Run()
	if err != nil {
		return err
	}
	return nil
}
