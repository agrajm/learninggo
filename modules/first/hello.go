package first

import (
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)
//Hello sys Hello World
func Hello() string {
	return quote.Hello()
}

func Proverb() string {
	return quoteV3.Concurrency()
}
