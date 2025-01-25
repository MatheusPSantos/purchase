package models

type RateOfExchange struct {
	CountryCurrencyDesc string `json:"country_currency_desc"`
	Currency            string `json:"currency"`
	ExchangeRate        string `json:"exchange_rate"`
	RecordDate          string `json:"record_date"`
}

type Rates struct {
	Data []RateOfExchange `json:"data"`
}

// exmplo de resposta
// {
// 	"data": [
// 		{
// 			"country_currency_desc": "Brazil-Real",
// 			"currency": "Real",
// 			"exchange_rate": "6.184",
// 			"record_date": "2024-12-31"
// 		}
// 	],
// }
