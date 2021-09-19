package indices

import (
	"fmt"

	"github.com/gocolly/colly"
)

type StockIndex int

type Indices struct {
	name   string
	symbol string
}

const (
	DAX StockIndex = iota
	TECDAX
	MDAX
	SDAX
	DOWJONES
	SP500
	NASDAQ100
	ESTOXX
	SMI
	ATX
	CAC40
	NIKKEI225
)

var StockIndices = []StockIndex{
	DAX, TECDAX, MDAX, SDAX, DOWJONES, SP500, NASDAQ100, ESTOXX, SMI, ATX, CAC40, NIKKEI225,
}

func (s StockIndex) String() string {
	return [...]string{"DAX", "TecDAX", "M-DAX", "S-DAX", "DOW JONES", "S&P 500", "NASDAQ 100", "EuroStoxx", "Swiss Market Index", "Austrian Traded Index", "France CAC 40 Index", "Japan NIKKEI Index"}[s]
}

func (s StockIndex) UriParam() string {
	return [...]string{"dax", "tecdax", "mdax", "sdax", "dow_jones", "s&p_500", "nasdaq_100", "euro_stoxx_50", "smi", "atx", "cac_40", "nikkei_225"}[s]
}

func (s StockIndex) EnumIndex() int {
	return int(s)
}

func LoadSingleValues() {
	for _, index := range StockIndices {
		url := fmt.Sprintf("https://www.finanzen.net/index/%s/werte", index.UriParam())
		fmt.Printf("%d %s\n", index.EnumIndex(), index)

		collector := colly.NewCollector(colly.AllowedDomains("finanzen.net", "www.finanzen.net"))

		collector.OnHTML("div.table-quotes div.table-responsive .table tbody td:nth-child(1) a", func(element *colly.HTMLElement) {
			fmt.Println(element.Text)
		})

		collector.OnRequest(func(request *colly.Request) {
			fmt.Println("Visiting", request.URL.String())
		})

		collector.Visit(url)
	}
}
