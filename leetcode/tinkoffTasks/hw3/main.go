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

	var openTime time.Time

	candles := make(map[string]domain.Candle)

	go func() {
		defer wg.Done()

		for price := range prices {
			time_, _ := domain.PeriodTS(domain.CandlePeriod1m, price.TS)
			if openTime.Before(time_) {
				openTime = time_
				for _, c := range candles {
					out <- c
				}
				candles = make(map[string]domain.Candle)
			}

			if candle, ok := candles[price.Ticker]; ok {
				candle.High = math.Max(candle.High, price.Value)
				candle.Low = math.Min(candle.Low, price.Value)
				candle.Close = price.Value
				candles[price.Ticker] = candle
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

	var openTime time.Time

	newCandles := make(map[string]domain.Candle)

	go func() {
		defer wg.Done()

		for candle := range candles {
			// some time solution
			time_, _ := domain.PeriodTS(period, candle.TS)
			if openTime.Before(time_) {
				openTime = time_
				for _, c := range newCandles {
					out <- c
				}
				newCandles = make(map[string]domain.Candle)
			}
			if existCandle, ok := newCandles[candle.Ticker]; ok {
				existCandle.High = math.Max(existCandle.High, candle.High)
				existCandle.Low = math.Min(existCandle.Low, candle.Low)
				existCandle.Close = candle.Close
				newCandles[candle.Ticker] = existCandle
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
