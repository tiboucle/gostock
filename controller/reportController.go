package controller

import (
	"fmt"
	strftime "github.com/jehiah/go-strftime"
	"net/http"
	"time"
)

type itemValue struct {
	SKU       string  `json:"sku"`
	Name      string  `json:"name"`
	Amount    int     `json:"amount"`
	PriceRate float64 `json:"price"`
	Value     float64 `json:"value"`
}

type itemValueSummary struct {
	PrintDate   string  `json:"printdate"`
	TotalSKU    int     `json:"totalsku"`
	TotalAmount int     `json:"totalamount"`
	TotalValue  float64 `json:"totalvalue"`
}

type itemValueReport struct {
	Items   []itemValue      `json:"items"`
	Summary itemValueSummary `json:"itemValueSummary"`
}

type transactionReport struct {
	InvoiceId   string  `json:"orderid"`
	InvoiceDate string  `json:"timestamp"`
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	Amount      int     `json:"amount"`
	Price       float64 `json:"price"`
	Omzet       float64 `json:"omzet"`
	Purchase    float64 `json:"purchase"`
	Profit      float64 `json:"profit"`
}

type salesSummary struct {
	PrintDate   string  `json:"printdate"`
	StartDate   string  `json:"startdate"`
	EndDate     string  `json:"enddate"`
	TotalSales  int     `json:"totalsales"`
	TotalAmount int     `json:"totalamount"`
	TotalOmzet  float64 `json:"totalomzet"`
	TotalProfit float64 `json:"totalprofit"`
}

type salesReport struct {
	Items   []transactionReport `json:"items"`
	Summary salesSummary        `json:"summary"`
}

func validateDate(date string) (time.Time, bool) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return t, false
	}
	return t, true
}

func GetItemValueReport(w http.ResponseWriter, r *http.Request) {
	if http.MethodGet == r.Method {
		itemValueReport := GenerateItemValueReport()
		WriteJSON(w, itemValueReport)
	} else {
		WriteDefaultResponse(w)
	}
}

func GetSalesReport(w http.ResponseWriter, r *http.Request) {
	if http.MethodGet == r.Method {
		startDate := r.FormValue("startdate")
		from, ok := validateDate(startDate)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Should have startdate params in YYYY-MM-DD format.")
			return
		}
		endDate := r.FormValue("enddate")
		to, ok := validateDate(endDate)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Should have enddate params in YYYY-MM-DD format.")
			return
		}
		salesReport := GenerateSalesReport(startDate, endDate)
		salesReport.Summary.StartDate = strftime.Format("%d %B %Y", from)
		salesReport.Summary.EndDate = strftime.Format("%d %B %Y", to)
		WriteJSON(w, salesReport)
	} else {
		WriteDefaultResponse(w)
	}
}

func GenerateSalesReport(startDate string, endDate string) salesReport {
	transaction := []transactionReport{}
	rows, _ := DB.Query(`
		SELECT 
			d.invoice_id,
			o.invoice_date,
			d.item_sku,
			(i.name || " " || "(" || size || "," || c.name || ")" ) AS name,
			d.amount,
			d.price,
			(d.amount * d.price) AS omzet,
			pd.price,
			((d.amount * d.price) - (d.amount * pd.price)) AS profit		
		FROM order_details d
			LEFT JOIN orders o
				ON d.invoice_id = o.invoice_id 
			LEFT JOIN items i 
				ON d.item_sku = i.sku
			LEFT JOIN colors c
				ON i.color = c.code
			LEFT JOIN purchase_details pd
				ON d.item_sku = pd.item_sku
		WHERE o.invoice_date BETWEEN '` + startDate + ` 00:00:00' 
				AND '` + endDate + ` 23:59:59';
	`)

	summary := salesSummary{}
	i := 0
	for rows.Next() {
		row := transactionReport{}
		rows.Scan(
			&row.InvoiceId, &row.InvoiceDate, &row.SKU, &row.Name, &row.Amount,
			&row.Price, &row.Omzet, &row.Purchase, &row.Profit,
		)
		transaction = append(transaction, row)

		summary.TotalAmount += row.Amount
		summary.TotalOmzet += row.Omzet
		summary.TotalProfit += row.Profit

		if i > 0 && transaction[i].InvoiceId != transaction[i-1].InvoiceId {
			summary.TotalSales++
		} else if i == 0 {
			summary.TotalSales = 1
		}
		i++
	}
	summary.PrintDate = strftime.Format("%d %B %Y", time.Now())

	report := salesReport{
		Items:   transaction,
		Summary: summary,
	}
	rows.Close()

	return report
}

func GenerateItemValueReport() itemValueReport {
	items := []itemValue{}
	rows, _ := DB.Query(`
		SELECT 
			i.sku, 
			(i.name || " " || "(" || i.size || "," || c.name || ")" ) AS name, 
			sum(d.amount) AS amount, 
			(sum(d.price * d.amount) / sum(d.amount)) AS price, 
			(sum(d.amount) * (sum(d.price * d.amount)/ sum(d.amount))) AS value 
		FROM orders o 
			LEFT JOIN order_details d 
				ON o.invoice_id = d.invoice_id 
			LEFT JOIN items i 
				ON d.item_sku = i.sku
			LEFT JOIN colors c
				ON i.color = c.code
		GROUP BY i.sku;
	`)

	itemValueSummary := itemValueSummary{}
	for rows.Next() {
		row := itemValue{}
		rows.Scan(&row.SKU, &row.Name, &row.Amount, &row.PriceRate, &row.Value)
		items = append(items, row)

		itemValueSummary.TotalAmount += row.Amount
		itemValueSummary.TotalValue += row.Value
	}
	itemValueSummary.TotalSKU = len(items)
	itemValueSummary.PrintDate = strftime.Format("%d %B %Y", time.Now())

	itemValueReport := itemValueReport{
		Items:   items,
		Summary: itemValueSummary,
	}
	rows.Close()

	return itemValueReport
}
