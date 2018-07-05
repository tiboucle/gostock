package controller

import (
	"net/http"
)

type item struct {
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

type items struct {
	Items []item `json:"items"`
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	if http.MethodGet == r.Method {
		WriteJSON(w, items{Items: FetchItems()})
	} else {
		WriteDefaultResponse(w)
	}
}

func FetchItems() []item {
	i := []item{}
	rows, _ := DB.Query(`
		SELECT 
		 sku, 
		 (i.name || " " || "(" || size || "," || c.name || ")" ) AS name,
		 stock
		FROM items i 
			LEFT JOIN colors c
				ON i.color = c.code;
	`)

	for rows.Next() {
		row := item{}
		rows.Scan(&row.SKU, &row.Name, &row.Stock)
		i = append(i, row)
	}
	rows.Close()

	return i
}
