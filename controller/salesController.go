package controller

import (
	"net/http"
)

type saleItem struct {
	Time    string  `json:"time"`
	SKU     string  `json:"sku"`
	Name    string  `json:"name"`
	Amount  int     `json:"amount"`
	Price   float64 `json:"price"`
	Total   float64 `json:"total"`
	Invoice string  `json:"invoice"`
}

type saleItems struct {
	Items []saleItem `json:"saleItems"`
}

func GetSalesItems(w http.ResponseWriter, r *http.Request) {
	if http.MethodGet == r.Method {
		WriteJSON(w, saleItems{Items: FetchSalesItems()})
	} else {
		WriteDefaultResponse(w)
	}
}

func FetchSalesItems() []saleItem {
	items := []saleItem{}
	rows, _ := DB.Query(`
		SELECT
			o.invoice_date,
			i.sku,
			(i.name || " " || "(" || i.size || "," || c.name || ")" ) AS name,
			d.amount,
			d.price,
			(d.amount * d.price) AS Total,
			o.invoice_id
		FROM orders o
			LEFT JOIN order_details d
				ON o.invoice_id = d.invoice_id
			LEFT JOIN items i 
				ON d.item_sku = i.sku
			LEFT JOIN colors c
				ON i.color = c.code;
	`)

	for rows.Next() {
		row := saleItem{}
		rows.Scan(
			&row.Time, &row.SKU, &row.Name, &row.Amount,
			&row.Price, &row.Total, &row.Invoice,
		)
		items = append(items, row)
	}
	rows.Close()

	return items
}
