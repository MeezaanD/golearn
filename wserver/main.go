package main

import (
	"flag"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type TableData struct {
	RandomData [][]int
	QRCode     string
}

var addr = flag.String("addr", ":1718", "http service address")

var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	rand.Seed(time.Now().UnixNano())
	// Generate random data
	data := make([][]int, 10)
	for i := range data {
		data[i] = make([]int, 10)
		for j := range data[i] {
			data[i][j] = rand.Intn(100)
		}
	}

	qrCode := req.FormValue("s")
	tableData := TableData{
		RandomData: data,
		QRCode:     qrCode,
	}

	templ.Execute(w, tableData)
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .QRCode}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.QRCode}}" />
<br>
{{.QRCode}}
<br>
<br>
{{end}}
<h1>Random Data Table</h1>
<table border="1">
	<tr>
		<th>Column 1</th>
		<th>Column 2</th>
		<th>Column 3</th>
		<th>Column 4</th>
		<th>Column 5</th>
		<th>Column 6</th>
		<th>Column 7</th>
		<th>Column 8</th>
		<th>Column 9</th>
		<th>Column 10</th>
	</tr>
	{{range .RandomData}}
	<tr>
		{{range .}}
		<td>{{.}}</td>
		{{end}}
	</tr>
	{{end}}
</table>
<br>
<form action="/" name=f method="GET">
    <input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
    <input type=submit value="Show QR" name=qr>
</form>
</body>
</html>
`
