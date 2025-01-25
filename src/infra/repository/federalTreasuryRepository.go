package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/matheuspsantos/purchase-wex/src/core/models"
	"github.com/pkg/errors"
)

const (
	url = "https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange"
)

func GetFromFederalTreasury(actualDate string, oldDate string, currency string) (*models.Rates, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Printf("Could not make request because error: %v", err)
		return nil, err
	}

	var (
		res  *http.Response
		body []byte
	)

	request.Header.Add("Content-type", "application/json")

	q := request.URL.Query()
	fd, ft := mountQuery(actualDate, oldDate, currency)
	q.Add("format", "json")
	q.Add("page[size]", "1")
	q.Add("sort", "-record_date") // pegar a data mais atual
	q.Add("fields", fd)
	q.Add("filter", ft)
	request.URL.RawQuery = q.Encode()
	log.Println(request.URL.RawQuery)
	res, err = http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var rates models.Rates
	if err = json.Unmarshal(body, &rates); err != nil {
		fmt.Errorf("error while enconding treasury api response. %w", err)
		return nil, err
	}
	log.Println("API response: ", rates)
	if len(rates.Data) == 0 {
		return nil, errors.New("The purchase cannot be converted to the target currency.")
	}

	return &rates, nil
}

// monta a query
// ad -> data atual
// od -> dat amais antiga
func mountQuery(actualDate, oldDate string, currency string) (fields string, filters string) {
	fields = "country_currency_desc,currency,exchange_rate,record_date"

	filters = fmt.Sprintf(
		"record_date:lte:%s,record_date:gte:%s,country_currency_desc:eq:%s",
		actualDate,
		oldDate,
		currency)
	return
}
