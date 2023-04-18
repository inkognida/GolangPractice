package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	logrus "github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type ValCurs struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Date       string     `xml:"Date,attr"`
	Name       string     `xml:"name,attr"`
	Currencies []Currency `xml:"Valute"`
}

type Currency struct {
	ID       string `xml:"ID,attr"`
	NumCode  uint16 `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  uint   `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

type MaxMinCurrency struct {
	Value float64
	Name  string
	Date  time.Time
}

type CurrencyRate struct {
	Value float64
	Date  time.Time
}

type CurrencyInfo struct {
	Name         string
	AverageValue float64
	MaxValue     CurrencyRate
	MinValue     CurrencyRate
	Counter      int
}

func MaxMinCurrencyRate(ValCursInfo map[string]ValCurs) (MaxMinCurrency, MaxMinCurrency, map[string]*CurrencyInfo) {
	maxValue := math.SmallestNonzeroFloat32
	minValue := math.MaxFloat32

	maxRate := MaxMinCurrency{}
	minRate := MaxMinCurrency{}
	currencies := make(map[string]*CurrencyInfo, 0)

	for _, v := range ValCursInfo {
		for _, c := range v.Currencies {
			value, _ := strconv.ParseFloat(strings.ReplaceAll(c.Value, ",", "."), 32)
			t, _ := time.Parse("02.01.2006", v.Date)

			// MAX and MIN currency
			if value > maxValue {
				maxValue = value

				maxRate.Name = c.Name
				maxRate.Value = value
				maxRate.Date = t
			}
			if value < minValue {
				minValue = value

				minRate.Name = c.Name
				minRate.Value = value
				minRate.Date = t
			}
			// MAX and MIN every currency INFO
			if _, ok := currencies[c.Name]; !ok {
				currencies[c.Name] = &CurrencyInfo{
					Name:         c.Name,
					AverageValue: value,
					MaxValue: CurrencyRate{
						Value: value,
						Date:  t,
					},
					MinValue: CurrencyRate{
						Value: value,
						Date:  t,
					},
					Counter: 1,
				}
			} else {
				if value > currencies[c.Name].MaxValue.Value {
					currencies[c.Name].MaxValue.Value = value
					currencies[c.Name].MaxValue.Date = t
				}
				if value < currencies[c.Name].MinValue.Value {
					currencies[c.Name].MinValue.Value = value
					currencies[c.Name].MinValue.Date = t
				}
				currencies[c.Name].AverageValue += value
				currencies[c.Name].Counter++
			}
		}
	}
	for k, _ := range currencies {
		currencies[k].AverageValue = currencies[k].AverageValue / float64(currencies[k].Counter)
	}
	return maxRate, minRate, currencies
}

func CBTask() {
	logger := logrus.New()

	ValCursInfo := make(map[string]ValCurs, 0)
	now := time.Now()
	for i := 0; i < 90; i++ {
		t := now.AddDate(0, 0, -i)
		url := fmt.Sprintf("http://www.cbr.ru/scripts/XML_daily_eng.asp?date_req=%s",
			t.Format("02/01/2006"))
		resp, err := http.Get(url)
		if err != nil {
			logger.Log(logrus.FatalLevel, err)
			os.Exit(1)
		}
		logger.Infoln("URL:", url)
		if _, ok := ValCursInfo[t.Format("02/01/2006")]; !ok {
			tmpValCurs := ValCurs{}
			data, _ := ioutil.ReadAll(resp.Body)
			r := bytes.NewReader(data)
			d := xml.NewDecoder(r)
			d.CharsetReader = charset.NewReaderLabel
			err = d.Decode(&tmpValCurs)
			if err != nil {
				logger.Log(logrus.DebugLevel, err)
			}
			ValCursInfo[t.Format("02/01/2006")] = tmpValCurs
		}

		resp.Body.Close()
	}
	max, min, currencyInfo := MaxMinCurrencyRate(ValCursInfo)

	logger.Println(max, min)
	for _, v := range currencyInfo {
		logger.Println(v.Name, v.MaxValue, v.MinValue, v.AverageValue)
	}
}

func main() {
	// first task
	// 1:MR
	// 2:IF FLAG
	// 3:GOTO 5
	// 4:GOTO 1
	// 5:MR
	// 6:MR
	// 7:GOTO 5

	// Если допустимо !FLAG
	// 1: MR
	// 2: if !FLAG
	// 3: goto 1
	// 4: MR
	// 5: MR
	// 6: goto 4

	// second task
	//x = берем шар из ящика ЧБ
	//Если x черный:
	//	чб ящик - черные шары
	//	белый ящик - чб шары
	//	черный ящик - белые шары
	//Иначе:
	//	чб ящик - белые шары
	//	белый ящик - черные шары
	//	черный ящик - чб шар

	// Т.к. все надписи заведомо ложны, то - если x черный - ящики Б и Ч не подходят под условие "заведомо ложны" ->
	// ЧБ ящик содержит черные шары (единственный верный вариант), Б ящик не
	// содержит черные шары и для него ложно содержание белых -> Б ящик - чб шары. Аналогично если х белый.

	// third task
	CBTask()
}
