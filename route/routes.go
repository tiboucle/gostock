package route

import (
	"fmt"
	"net/http"
)

func Run(p int) (e error) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", p), nil)
	return err
}
