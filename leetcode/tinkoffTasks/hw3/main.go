package main

import (
	"context"
	"fmt"
	"hw3/domain"
	"hw3/generator"
	"math"
	"os"
	"os/signal"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var tickers = []string{"AAPL", "SBER", "NVDA", "TSLA"}

func initCandle(arg interface{}, period domain.CandlePeriod) domain.Candle {
	switch val := arg.(type) {
	case domain.Price:
		ts, err := domain.PeriodTS(period, val.TS)
		if err != nil {
			logger.Info(err)
		}
		return domain.Candle{
			Ticker: val.Ticker,
			Period: period,
			Open:   val.Value,
			High:   val.Value,
			Low:    val.Value,
			Close:  val.Value,
			TS:     ts,
		}
	case domain.Candle:
		ts, err := domain.PeriodTS(period, val.TS)
		if err != nil {
			logger.Info(err)
		}
		return domain.Candle{
			Ticker: val.Ticker,
			Period: period,
			Open:   val.Open,
			High:   val.High,
			Low:    val.Low,
			Close:  val.Close,
			TS:     ts,
		}
	}
	return domain.Candle{}
}

func formCandles(prices <-chan domain.Price, wg *sync.WaitGroup) <-chan domain.Candle {
	out := make(chan domain.Candle)
	wg.Add(1)

	candles := make(map[string]domain.Candle)

	go func() {
		defer wg.Done()

		// TODO openTime part
		for price := range prices {
			if candle, ok := candles[price.Ticker]; ok {
				candle.High = math.Max(candle.High, price.Value)
				candle.Low = math.Min(candle.Low, price.Value)
				candle.Close = price.Value
			} else {
				candles[price.Ticker] = initCandle(price, domain.CandlePeriod1m)
			}
		}

		for _, c := range candles {
			out <- c
		}
		close(out)
	}()
	return out
}

func saveCandleCSV(candle domain.Candle) {
	file, err := os.OpenFile(fmt.Sprintf("candles_%s.csv", candle.Period),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		logger.Info(err)
	}

	_, err = file.WriteString(fmt.Sprintf("%v\n", candle))
	if err != nil {
		logger.Info(err)
	}
	file.Close()
}

func saveCandles(candles <-chan domain.Candle, wg *sync.WaitGroup) <-chan domain.Candle {
	out := make(chan domain.Candle)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for candle := range candles {
			out <- candle
			saveCandleCSV(candle)
		}
		close(out)
	}()
	return out
}

func handleCandles(candles <-chan domain.Candle, wg *sync.WaitGroup,
	period domain.CandlePeriod) <-chan domain.Candle {
	out := make(chan domain.Candle)
	wg.Add(1)

	newCandles := make(map[string]domain.Candle)

	go func() {
		defer wg.Done()

		// TODO openTime part
		for candle := range candles {
			if existCandle, ok := newCandles[candle.Ticker]; ok {
				existCandle.High = math.Max(existCandle.High, candle.High)
				existCandle.Low = math.Min(existCandle.Low, candle.Low)
				existCandle.Close = candle.Close
			} else {
				newCandles[candle.Ticker] = initCandle(candle, period)
			}
		}

		for _, c := range newCandles {
			out <- c
		}

		close(out)
	}()

	return out
}

var logger = log.New()

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	pg := generator.NewPricesGenerator(generator.Config{
		Factor:  10,
		Delay:   time.Millisecond * 500,
		Tickers: tickers,
	})

	logger.Info("start prices generator...")
	prices := pg.Prices(ctx)

	wg := sync.WaitGroup{}

	oneMinCandles := formCandles(prices, &wg)
	twoMinPrices := saveCandles(oneMinCandles, &wg)

	twoMinCandles := handleCandles(twoMinPrices, &wg, domain.CandlePeriod2m)
	tenMinPrices := saveCandles(twoMinCandles, &wg)

	tenMinCandles := handleCandles(tenMinPrices, &wg, domain.CandlePeriod10m)
	nextMinPrices := saveCandles(tenMinCandles, &wg)

	go func() { // ENDING
		for _ = range nextMinPrices {
		}
	}()

	<-stop
	cancel()
	wg.Wait()

	logger.Info("end")
}

/* package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"hw3/domain"
	"hw3/generator"
	"math"
	"os"
	"os/signal"
	"sync"
	"time"
)

func FormCandleFromPrice(wg *sync.WaitGroup, prices <-chan domain.Price) <-chan domain.Candle {
	out := make(chan domain.Candle)
	var openTime time.Time
	candles := make(map[string]domain.Candle)
	go func() {
		defer func() {
			for _, c := range candles {
				out <- c
			}
			close(out)
			wg.Done()
		}()
		for price := range prices {
			time, _ := domain.PeriodTS(domain.CandlePeriod1m, price.TS)
			if openTime.Before(time) {
				openTime = time
				for _, c := range candles {
					out <- c
				}
				candles = make(map[string]domain.Candle)
			}
			if _, ok := candles[price.Ticker]; !ok {
				candles[price.Ticker] = domain.Candle{price.Ticker, domain.CandlePeriod1m, price.Value, price.Value,
					price.Value, price.Value, openTime}
			}
			candle := candles[price.Ticker]
			candle.High = math.Max(candle.High, price.Value)
			candle.Low = math.Min(candle.Low, price.Value)
			candle.Close = price.Value
			candles[price.Ticker] = candle
		}
	}()
	return out
}

func FormCandle(wg *sync.WaitGroup, candles <-chan domain.Candle, period domain.CandlePeriod) <-chan domain.Candle {
	out := make(chan domain.Candle)
	var openTime time.Time
	newCandles := make(map[string]domain.Candle)
	go func() {
		defer func() {
			for _, c := range newCandles {
				out <- c
			}
			close(out)
			wg.Done()
		}()
		for candle := range candles {
			time, _ := domain.PeriodTS(period, candle.TS)
			if openTime.Before(time) {
				openTime = time
				for _, c := range newCandles {
					out <- c
				}
				newCandles = make(map[string]domain.Candle)
			}
			if _, ok := newCandles[candle.Ticker]; !ok {
				newCandles[candle.Ticker] = domain.Candle{candle.Ticker, period, candle.Open, candle.High,
					candle.Low, candle.Close, openTime}
			}
			c := newCandles[candle.Ticker]
			c.High = math.Max(candle.High, c.High)
			c.Low = math.Min(candle.Low, c.Low)
			c.Close = candle.Close
			newCandles[candle.Ticker] = c
		}
	}()
	return out
}

func FormCandle2m(wg *sync.WaitGroup, candles <-chan domain.Candle) <-chan domain.Candle {
	return FormCandle(wg, candles, domain.CandlePeriod2m)
}

func FormCandle10m(wg *sync.WaitGroup, candles <-chan domain.Candle) <-chan domain.Candle {
	return FormCandle(wg, candles, domain.CandlePeriod10m)
}

func save(candle domain.Candle) {
	f, err := os.OpenFile(fmt.Sprintf("candles_%s.csv", candle.Period), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("%v\n", candle))
}

func Save(wg *sync.WaitGroup, candles <-chan domain.Candle) <-chan domain.Candle {
	out := make(chan domain.Candle)
	go func() {
		defer wg.Done()
		for candle := range candles {
			out <- candle
			save(candle)
		}
		close(out)
	}()
	return out
}

var tickers = []string{"AAPL", "SBER", "NVDA", "TSLA"}

func main() {
	logger := log.New()
	ctx, cancel := context.WithCancel(context.Background())
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	wg := sync.WaitGroup{}

	pg := generator.NewPricesGenerator(generator.Config{
		Factor:  10,
		Delay:   time.Millisecond * 500,
		Tickers: tickers,
	})

	logger.Info("start prices generator...")
	wg.Add(6)
	prices := pg.Prices(ctx)
	candles := FormCandleFromPrice(&wg, prices)
	candles = Save(&wg, candles)
	candles = FormCandle2m(&wg, candles)
	candles = Save(&wg, candles)
	candles = FormCandle10m(&wg, candles)
	candles = Save(&wg, candles)

	go func() {
		for _ = range candles {
		}
	}()
	<-stop
	cancel()
	wg.Wait()
	logger.Info("all goroutines terminated")
	logger.Info("main process terminated")
} */
