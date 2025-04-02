package admin

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type Handler interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
	GetProductDetails(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	userCase UserCase
}

func NewHandler(userCase UserCase) Handler {
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

var listProductTemplate *template.Template
var productDetailTempl *template.Template

func init() {
	dir, err := os.Getwd()
	log.Println(dir)
	if err != nil {
		log.Fatal(err)
	}
	listProductTemplatePath := filepath.Join(dir, "web/admin/list_product.html")
	listProductTemplate = template.Must(template.New("list_product.html").Funcs(templateFuncs).ParseFiles(listProductTemplatePath))

	productDetailTemplatePath := filepath.Join(dir, "web/admin/product_detail.html")
	productDetailTempl = template.Must(template.New("product_detail.html").Funcs(productDetailsTemplateFuncs).ParseFiles(productDetailTemplatePath))
}

func (h *handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 5
	}

	categoryID, err := strconv.Atoi(r.URL.Query().Get("category_id"))
	if err != nil {
		categoryID = 0
	}

	searchKey := r.URL.Query().Get("search_key")

	req := PaginationRequest{
		Page:       page,
		Limit:      limit,
		CategoryId: categoryID,
		SearchKey:  searchKey,
	}

	response, err := h.userCase.GetProducts(r.Context(), req)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	err = listProductTemplate.Execute(w, response)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func (h *handler) GetProductDetails(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Something errors", http.StatusInternalServerError)
		return
	}

	productDetails, err := h.userCase.GetProductDetails(r.Context(), id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	if err := productDetailTempl.Execute(w, productDetails); err != nil {
		log.Println(err)
		http.Error(w, "Something errors", http.StatusInternalServerError)
	}
}

var productDetailsTemplateFuncs = template.FuncMap{
	"formatPrice": func(price float64) string {
		return strconv.FormatFloat(price, 'f', 2, 64)
	},
}
