#import "input.typ": *
#import "utils.typ": *

#let info = table(
  columns: (1fr, 1fr),
  inset: (x: 10pt),
  align: horizon,
  title("Titular", "Holder", flexed: true), title("Produto", "Product", flexed: true), 
  customerName, prodName
)

#let date = {
  table(
    columns: (1fr, 1fr),
    inset: (x: 10pt),
    align: horizon,
    title("Data de Criação", "Creation Date", flexed: true), title("Datas de Viagem", "Travel Dates", flexed: true),
    today, dateIn + " / " + dateOut
  )
} 


#let services = table(
  columns: (1fr, 10fr, 1fr, 1fr, 1.5fr, 1.5fr),
  inset: (x, y) => if x == 0 {0pt} else {(x: 7pt, y:5pt)},
  align: (center + horizon, left + horizon, center + horizon, center + horizon, center + horizon, center + horizon),
  table.header("", title("Descrição", "Description", flexed: true), title("QTD", "QTY", flexed: true),  title("Estado", "Status", flexed: true), title("De", "From", flexed: true), title("Até", "To", flexed: true)),
  ..(allServices),
)

#let prices = {
  table(
    columns: (10fr, 1.5fr, 1.5fr, 1.5fr),
    inset: (x: 7pt, y: 5pt),
    align: (left + horizon, center + horizon, left + horizon, left + horizon),
    table.header(title("Descrição", "Description", flexed: true), title("QTD", "QTY", flexed: true),  title("PVP Unid.", "Unit Gross", flexed: true), title("PVP Total", "Total Gross", flexed: true),
    ),
    ..(allPrices),
  )
  rect(
    fill: gray,
    width: 100%,
    height: 5%,
    inset: (x: 0pt, y: 10pt),
    [
      #grid(
        align: center + top,
        columns: (10fr, 1.5fr, 1.5fr, 1.5fr),
        "", "",  title("PVP Total", "Total Gross", flexed: true), strong(totalPrice)
        
      )
    ]
  )
}


#let policies = {
  table(
    columns: (1fr, 1fr, 1fr, 1fr, 1fr),
    inset: (x: 7pt, y: 5pt),
    align: left + horizon,
    table.header(title("Tipo", "Type", flexed: true), title("Serviço", "Service", flexed: true),  title("De", "From", flexed: true), title("Até", "To", flexed: true), title("Valor", "Price", flexed: true)
    ),
    ..(allPolicies),
  )
}
