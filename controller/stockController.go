package controller

import (
	"net/http"
)

type incomingItem struct {
	Time     string `json:"time"`
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Amount   int    `json:"amount"`
	Received int    `json:"received"`
	Price    int    `json:"price"`
	Total    int    `json:"total"`
	Invoice  string `json:"invoice"`
	Note     string `json:"note"`
}

type incomingItems struct {
	Items []incomingItem `json:"incomingitems"`
}

func GetIncomingItems(w http.ResponseWriter, r *http.Request) {
	if http.MethodGet == r.Method {
		WriteJSON(w, incomingItems{Items: FetchIncomingItems()})
	} else {
		WriteDefaultResponse(w)
	}
}

func FetchIncomingItems() []incomingItem {
	items := []incomingItem{}
	rows, _ := DB.Query(`
		SELECT
			ii.time_received,
			i.sku,
			(i.name || " " || "(" || i.size || "," || c.name || ")" ) AS name,
			d.amount,
			ii.item_received,
			d.price,
			(d.amount * d.price) AS Total,
			ii.invoice_id,
			ii.note
		FROM incoming_items ii
			LEFT JOIN items i 
				ON ii.item_sku = i.sku
			LEFT JOIN colors c
				ON i.color = c.code
			LEFT JOIN purchases p
				ON ii.invoice_id = p.invoice_id
			LEFT JOIN purchase_details d
				ON p.invoice_id = d.invoice_id;
	`)

	for rows.Next() {
		row := incomingItem{}
		rows.Scan(
			&row.Time, &row.SKU, &row.Name, &row.Amount, &row.Received,
			&row.Price, &row.Total, &row.Invoice, &row.Note,
		)
		items = append(items, row)
	}
	rows.Close()

	return items
}
