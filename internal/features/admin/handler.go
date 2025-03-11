package admin

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type handler struct {
	userCase UserCase
}

func NewHandler(userCase UserCase) *handler {
	return &handler{
		userCase: userCase,
	}
}

var templateFuncs = template.FuncMap{
	// Subtract function for pagination
	"subtract": func(a, b int) int {
		return a - b
	},
	// Add function for pagination
	"add": func(a, b int) int {
		return a + b
	},
	// Generate page range for pagination
	"pageRange": func(currentPage, totalPages int) []int {
		// Show up to 5 pages centered around current page
		start := max(currentPage-2, 1)

		end := start + 4
		if end > totalPages {
			end = totalPages
			// Adjust start to show 5 pages if possible
			start = max(end-4, 1)
		}

		var pages []int
		for i := start; i <= end; i++ {
			pages = append(pages, i)
		}
		return pages
	},
	// Format pagination info text
	"paginationInfo": func(meta Metadata) string {
		start := (meta.CurrentPage-1)*meta.PageSize + 1
		end := min(start+meta.PageSize-1, meta.TotalRecords)
		if meta.TotalRecords == 0 {
			start = 0
		}
		return fmt.Sprintf("%d-%d", start, end)
	},
}

func (h *handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	req := PaginationRequest{
		Page:  page,
		Limit: limit,
	}

	response, err := h.userCase.GetProducts(r.Context(), req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	templates := template.Must(template.New("").Funcs(templateFuncs).ParseGlob("web/admin/*.html"))

	err = templates.ExecuteTemplate(w, "list_product.html", response)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}
