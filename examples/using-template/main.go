package main

import (
	"os"
	"text/template"

	"github.com/hitalos/paginator"
)

func main() {
	tmpl, _ := template.New("").Parse(`<ul class="pagination">
		{{- range .Pages }}
	<li{{ if .Actual }} class="actual"{{ end }}>` +
		`<a href="{{ .Link }}">{{ .Title }}</a>` +
		`</li>
		{{- end }}
</ul>`)

	totalOfItems := 256
	pagination := paginator.New(totalOfItems)
	pagination.SetActualPage(7)           // optional - default 1
	pagination.SetPageLimit(10)           // optional - default 10
	pagination.SetPageRange(5)            // optional - default 5
	pagination.SetPagePath("page/")       // optional - default "page/"
	pagination.SetPrefix("/admin/posts/") // optional - default ""
	pagination.Paginate()

	_ = tmpl.Execute(os.Stdout, pagination)
}
