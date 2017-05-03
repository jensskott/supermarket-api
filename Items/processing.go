package Items

import (
	"net/http"
	"strconv"
)

func FormToItem(r *http.Request) (Item, []string) {
	var item Item
	var errStr, quantityStr string
	var errs []string
	var err error

	item.Name, errStr = processFormField(r, "name")
	errs = appendError(errs, errStr)
	quantityStr, errStr = processFormField(r, "quantity")
	if len(errStr) != 0 {
		errs = append(errs, errStr)
	} else {
		item.Quantity, err = strconv.Atoi(quantityStr)
		if err != nil {
			errs = append(errs, "Parameter 'quantity' not an integer")
		}
	}
	return item, errs
}

func appendError(errs []string, errStr string) []string {
	if len(errStr) > 0 {
		errs = append(errs, errStr)
	}
	return errs
}

func processFormField(r *http.Request, field string) (string, string) {
	fieldData := r.PostFormValue(field)
	if len(fieldData) == 0 {
		return "", "Missing '" + field + "' parameter, cannot continue"
	}
	return fieldData, ""
}
