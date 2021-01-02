package paginator

import (
	"strconv"
)

type page struct {
	Title  string `json:"title"`
	Link   string `json:"link"`
	Actual bool   `json:"actual"`
}

// Paginator struct to represents a list of pages.
type Paginator struct {
	count      int
	pageLimit  int
	pageRange  int
	actualPage int
	prefixLink string
	pagePath   string
	Pages      []page `json:"pages"`
}

// New returns a new Paginator struct.
func New(count int) Paginator {
	return Paginator{
		count:      count,
		pageLimit:  10,
		pageRange:  5,
		prefixLink: "",
		pagePath:   "page/",
		actualPage: 1,
		Pages:      []page{},
	}
}

// SetPrefix set string to prepend links.
func (p *Paginator) SetPrefix(prefix string) {
	p.prefixLink = prefix
}

// SetPagePath set string to be a separator to prepend page links.
func (p *Paginator) SetPagePath(path string) {
	p.pagePath = path
}

// SetActualPage set number of actual page.
func (p *Paginator) SetActualPage(n int) {
	if n > 0 {
		p.actualPage = n
	}
}

// SetPageLimit set limit of items in a page.
func (p *Paginator) SetPageLimit(n int) {
	if n > 0 {
		p.pageLimit = n
	}
}

// SetPageRange set range to show pages before and after of actual.
func (p *Paginator) SetPageRange(n int) {
	if n > 0 {
		p.pageRange = n
	}
}

func (p *Paginator) calcPagesCount() int {
	pagesCount := p.count / p.pageLimit
	if p.count%p.pageLimit > 0 {
		pagesCount++
	}
	return pagesCount
}

func (p *Paginator) addFirstAndPrevious() {
	link := p.prefixLink
	p.Pages = append(p.Pages, page{"≪", link, false})
	if p.actualPage-1 != 1 {
		link = p.prefixLink + p.pagePath + strconv.Itoa(p.actualPage-1)
	}
	p.Pages = append(p.Pages, page{"<", link, false})
}

func (p *Paginator) addLinkToPageNumber(n int) {
	link := p.prefixLink
	if n != 1 {
		link = p.prefixLink + p.pagePath + strconv.Itoa(n)
	}

	p.Pages = append(p.Pages, page{strconv.Itoa(n), link, p.actualPage == n})
}

func (p *Paginator) addNextAndLast(pagesCount int) {
	link := p.prefixLink + p.pagePath
	p.Pages = append(p.Pages, page{">", link + strconv.Itoa(p.actualPage+1), false})
	p.Pages = append(p.Pages, page{"≫", link + strconv.Itoa(pagesCount), false})
}

// Paginate mount the list of links using previously configured attributes.
// Run before render.
func (p *Paginator) Paginate() {
	p.Pages = []page{}
	pagesCount := p.calcPagesCount()

	if p.actualPage > 1 {
		p.addFirstAndPrevious()
	}

	for i := p.actualPage - p.pageRange; i < p.actualPage+p.pageRange; i++ {
		if i < 1 || i > pagesCount {
			continue
		}

		p.addLinkToPageNumber(i)
	}

	if p.actualPage < pagesCount {
		p.addNextAndLast(pagesCount)
	}
}

func (p Paginator) String() string {
	html := "<ul>"
	for _, pg := range p.Pages {
		html += "<li"
		if pg.Actual {
			html += ` class="actual"`
		}
		html += `><a href="` + pg.Link + `">` + pg.Title + "</a></li>"
	}
	html += "</ul>"

	return html
}
