#import "input.typ": *

#align(center)[== Vizela Viagens]

#table(
  columns: (auto, 1fr, auto),
  inset: 10pt,
  align: horizon,
  table.header(
    [*Categoria*], [*Descrição*], [*Preço*],
  ),
  "Voo",
  getFlightDesc(),
  getFlightPrice(),
  "Hotel",
  getHotelDesc(),
  getHotelPrice(),
)
