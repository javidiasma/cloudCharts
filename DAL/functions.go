package DAL

import (
	"cloudCharts/config"
	"fmt"
)

var Countries map[int]string

func init() {
	Countries = make(map[int]string, 0)
	Countries[1] = "na_sales"
	Countries[2] = "eu_sales"
	Countries[3] = "jp_sales"
	Countries[4] = "other_sales"
	Countries[5] = "global_sales"
}

func GenerateBarItemsByName(name string) []float32 {
	items := make([]float32, 0)
	var tempCountry string
	for i := 0; i < 5; i++ {
		var tempValue float32
		tempCountry = Countries[i+1]
		tempCountry = "SUM(" + tempCountry + ")"
		config.DB.Select(tempCountry).Where("name = ?", name).Scan(&tempValue)
		items = append(items, tempValue)
	}
	fmt.Println(items)
	return items
}
