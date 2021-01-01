package main

import (
	"os"

	"github.com/hitalos/paginator"
)

func main() {
	totalOfItems := 256
	pagination := paginator.New(totalOfItems)
	pagination.SetActualPage(7)           // optional - default 1
	pagination.SetPageLimit(10)           // optional - default 10
	pagination.SetPageRange(5)            // optional - default 5
	pagination.SetPagePath("page/")       // optional - default "page/"
	pagination.SetPrefix("/admin/posts/") // optional - default ""
	pagination.Paginate()

	os.Stdout.WriteString(pagination.String())
}
