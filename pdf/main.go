package pdf

import "fmt"

func main() {
	pdf := NewPDF("Novo", "Muito Mau", 100.0, 200.0)

	pdf.GeneratePDF()

	f := pdf.GetPDFFile()

	fmt.Println("Please read the", f.Name(), "file")
}
