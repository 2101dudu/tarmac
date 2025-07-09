#import "input.typ": *
#import "utils.typ": *
#import "tables.typ": *

#set text(font: "Arial", size: 8pt)

#set page(
  margin: (
    top: 4cm,
    x: 1.5cm,
    bottom: 2cm,
  ),
  header: [
    #grid(
      columns: (5fr, 4fr, 3fr),
      rows: (60pt),
      gutter: 3pt,
        grid.cell(
      align: left + horizon,
        image("assets/logo_horiz.svg", width: 100%),
      ),
      "",
      grid.cell(
        align: right + horizon,
        "Orçamento / " + quotationNumber
      ),
    )
    #v(-6%)
  ],
  footer: context [
    #line(
      length: 100%
    )
    #set text(size: 5pt)
    #set align(left)
    [
      Viajar Tours – Clube Viajar, Lda | Rua Damião de Góis, 9 s/l Esq. Tras. | 4050 - 225 Porto
      \
      NIPC 503542016 Capital Social €250 000.00 € - Matrícula da conservatória e Local: 503542016 - Registo RNAVT: 2146
      \
      Tel.: +351 225 025 805 (Chamada Para Rede Fixa Nacional) | www.viajartours.pt | E-Mail: info\@viajartours.pt
    ]
    #place(dy: -10pt, dx: 98%)[
      #counter(page).display(
        "1 / 1",
        both: true,
      )
    ]
  ]
)

#set table(
   fill: (x, y) =>
    if y == 0 {accent}
    else if calc.even(y)  { gray },
   gutter: 0pt,
   stroke: 0.1pt,
)




#entry("INFORMAÇÃO GERAL", "General Information", info, date)

#entry("SERVIÇOS", "Services", services, none)

#entry("PREÇOS", "Prices", prices, none)

#entry("POLÍTICAS DE CANCELAMENTO", "Cancel Policies", policies, none)

#entry("CONDIÇÕES", "Conditions", conditions, none)
#entry("DESCRIÇÃO", "Description", description, none)
#entry("PROGRAMA", "Program", program, none)
#entry("OBSERVAÇÕES", "Observations", observations, none)
