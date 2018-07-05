package route

import (
	"fmt"
	"gostock/controller"
	"gostock/service"
	"net/http"
)

func Run(p int) (e error) {
	RegisterDatabase()
	RegisterRoute()
	err := http.ListenAndServe(fmt.Sprintf(":%d", p), nil)
	return err
}

func RegisterDatabase() {
	controller.DB = service.InitConnection()
}

func RegisterRoute() {
	http.HandleFunc("/api/items", controller.GetItems)
	http.HandleFunc("/api/incoming/items", controller.GetIncomingItems)
	http.HandleFunc("/api/sales/items", controller.GetSalesItems)
	http.HandleFunc("/api/sales/item/value/report", controller.GetItemValueReport)
	http.HandleFunc("/api/sales/item/report", controller.GetSalesReport)
}
