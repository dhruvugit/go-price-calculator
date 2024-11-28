package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager `json:"-"`
	TaxRate           float64                 `json:"tax_rate"`
	InputPrices       []float64               `json:"input_prices"`
	TaxIncludedPrices map[string]string       `json:"tax_included_prices"`
}

/*
As package name is price so we would avoid writing just New() for constructor else
we will prefer to use proper naming for what is actually creates
*/
func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Println("Result for Process is : ", result)
	job.TaxIncludedPrices = result
	/*filemanager.WriteJSON("job.json", job) will write to same file again and again
	which will eventually override values hence using a variable file name for each
	job as per tax rate and did some calculation to avoid the decimal tax in name*/

	//filemanager.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
	job.IOManager.WriteResult(job)

}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}
	price, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = price
}
