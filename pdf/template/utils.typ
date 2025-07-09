#let accent = rgb("#D5F5D9")
#let gray = rgb("#1111")

#let plane = {image("assets/plane.svg", width: 14pt, fit: "stretch")}
#let bed = {image("assets/bed.svg", width: 14pt, fit: "stretch")}

#let title(name, name_en, size: 7pt, flexed: false) = {
  strong(name)
  if flexed {"\n"} else {" / "}
  text(size: size, style: "italic", name_en)
}


#let entry(text, text_en, t, t2) = {
  let shape() = {
    polygon(
      fill: accent,
      (0%, 0pt),
      (10%, 0pt),
      (7%, 3%),
      (0%,  3%),
    )
  }
  place(dx: -10%,
    shape()
  )
  line(
    start: (-10%, 0pt),
    length: 120%,
    stroke: 0.5pt + accent,
  )
  title(text, text_en)
  if t2 != none {
    grid(
      columns: (1fr, 1fr),
      column-gutter: 30pt,
      t, t2
    )
  } else {
    t
  }
  "\n"
}
