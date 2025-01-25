package usecases

import (
	"log"
	"strconv"
	"strings"

	"github.com/matheuspsantos/purchase-wex/src/core/dtos"
	"github.com/matheuspsantos/purchase-wex/src/core/models"
	"github.com/matheuspsantos/purchase-wex/src/core/utils"
	"github.com/matheuspsantos/purchase-wex/src/infra/repository"
)

func GetPurchaseTransactionByIdUseCase(id string, cur string) (*dtos.PurchaseTrxWithExchangeDTOResponse, error) {
	purchase := models.Purchase{}
	repository.FindPurchaseById(id, &purchase)

	actualData := utils.ConvertTimeToYearMonthDayFormat(purchase.TransactionDate)
	retro_date := purchase.TransactionDate.AddDate(0, -6, 0)
	oldDate := utils.ConvertTimeToYearMonthDayFormat(retro_date)

	log.Println(actualData)
	log.Println(oldDate)

	// tem que fazer a consulta no repositorio para pegar do tesouro
	res, err := repository.GetFromFederalTreasury(actualData, oldDate, cur)
	if err != nil {
		return nil, err
	}
	// se n tiver valor no delta de 6 meses tem q lançar o erro descrito no teste
	log.Print("Converting values to response")
	var pres dtos.PurchaseTrxWithExchangeDTOResponse
	pres.ID = uint64(purchase.ID)
	pres.Description = purchase.Description
	pres.OriginalDollarAmount = purchase.Amount
	pres.TransactionDate = strings.Split(purchase.TransactionDate.String(), " ")[0]
	pres.ExchangeRateUsed = convertStrToFloat(res.Data[0].ExchangeRate)
	pres.ConvertedAmount = convertStrToFloat(res.Data[0].ExchangeRate) * purchase.Amount
	// o valro tem que ser arredondado para duas casas decimais
	return &pres, nil
}

func convertStrToFloat(str string) float64 {
	conv, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Print(err)
	}
	return utils.RoundFloat(conv, 2)
}
