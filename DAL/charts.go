package DAL

import (
	"cloudCharts/DAO"
	"cloudCharts/config"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"os"
)

func First(firstGame string, secondGame string, firstGameSales []float32, secondGameSales []float32) string {
	bar := charts.NewBar()

	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Compare sales of two games on a chart",
	}))
	var firstGameModel DAO.History
	var secondGameModel DAO.History
	config.DB.Where(&DAO.History{Name: firstGame}).Find(&firstGameModel)
	config.DB.Where(&DAO.History{Name: secondGame}).Find(&secondGameModel)

	bar.SetXAxis([]string{firstGameModel.Name, secondGameModel.Name}).
		AddSeries("first", generateFirstGameBarItems(firstGameSales)).
		AddSeries("second", generateFirstGameBarItems(secondGameSales))
	f, _ := os.Create("firstBar.html")
	_ = bar.Render(f)

	return f.Name()
}

func generateFirstGameBarItems(gameSales []float32) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 5; i++ {
		items = append(items, opts.BarData{Value: gameSales[i], Label: &opts.Label{Show: true}})
	}
	return items
}
