package indices

import "fmt"

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

func (s StockIndex) EnumIndex() int {
	return int(s)
}

func LoadSingeValues() {
	for _, index := range StockIndices {
		fmt.Printf("%d %s\n",index.EnumIndex(), index)
	}
}
