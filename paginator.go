package paginator

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	defaultRange = 5
	defaultLimit = 10
)

type page struct {
	Title string `json:"title"`
	Link  string `json:"link"`
	Class string `json:"class"`
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
		pageLimit:  defaultLimit,
		pageRange:  defaultRange,
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
	p.Pages = append(p.Pages, page{"⇤", link, "first"})

	if p.actualPage-1 != 1 {
		link = p.prefixLink + p.pagePath + strconv.Itoa(p.actualPage-1)
	}

	p.Pages = append(p.Pages, page{"←", link, "previous"})
}

func (p *Paginator) addLinkToPageNumber(n int) {
	link := p.prefixLink
	if n != 1 {
		link = p.prefixLink + p.pagePath + strconv.Itoa(n)
	}

	class := ""
	if p.actualPage == n {
		class = "actual"
	}
	p.Pages = append(p.Pages, page{strconv.Itoa(n), link, class})
}

func (p *Paginator) addNextAndLast(pagesCount int) {
	link := p.prefixLink + p.pagePath
	p.Pages = append(p.Pages, page{"→", link + strconv.Itoa(p.actualPage+1), "next"})
	p.Pages = append(p.Pages, page{"⇥", link + strconv.Itoa(pagesCount), "last"})
}

// Paginate mount the list of links using previously configured attributes.
// Run before render.
func (p *Paginator) Paginate() {
	p.Pages = []page{}
	pagesCount := p.calcPagesCount()

	if p.actualPage > 1 {
		p.addFirstAndPrevious()
	}

	for i := p.actualPage - p.pageRange; i <= p.actualPage+p.pageRange; i++ {
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
	list := []string{}
	for _, pg := range p.Pages {
		classAttr := ""
		if pg.Class != "" {
			classAttr = fmt.Sprintf(" class=%q", pg.Class)
		}

		list = append(list, fmt.Sprintf("\n\t<li%s><a href=%q>%s</a></li>", classAttr, pg.Link, pg.Title))
	}

	return fmt.Sprintf("<ul class=\"paginator\">%s\n</ul>\n", strings.Join(list, ""))
}
