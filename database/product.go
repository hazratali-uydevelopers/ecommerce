package database

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

var ProductList []Product

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Fresh orange form Bangladesh.",
		Price:       100,
		ImgURL:      "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Fresh Apple form Bangladesh.",
		Price:       140,
		ImgURL:      "https://5.imimg.com/data5/AK/RA/MY-68428614/apple.jpg",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Orange",
		Description: "Fresh orange form Bangladesh.",
		Price:       100,
		ImgURL:      "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Orange",
		Description: "Fresh orange form Bangladesh.",
		Price:       100,
		ImgURL:      "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Orange",
		Description: "Fresh orange form Bangladesh.",
		Price:       100,
		ImgURL:      "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
	}
	prd6 := Product{
		ID:          6,
		Title:       "Mango",
		Description: "Fresh Mango form Bangladesh.",
		Price:       10000000,
		ImgURL:      "https://img.freepik.com/free-photo/mango-still-life_23-2151542114.jpg?w=360",
	}
	ProductList = append(ProductList, prd1, prd2, prd3, prd4, prd5, prd6)
}
