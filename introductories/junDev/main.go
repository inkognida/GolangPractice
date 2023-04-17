package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"math"
	"net/http"
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

type CurrencyRate struct {
	Value float32
	Name  string
	Date  time.Time
}

type AverageRate struct {
	Name    string
	Value   float64
	Counter int
}

func MaxMinCurrencyRate(ValCursInfo map[string]ValCurs) (CurrencyRate, CurrencyRate) {
	maxValue := math.SmallestNonzeroFloat32
	minValue := math.MaxFloat32

	maxRate := CurrencyRate{}
	minRate := CurrencyRate{}
	for _, v := range ValCursInfo {
		for _, c := range v.Currencies {
			value, _ := strconv.ParseFloat(strings.ReplaceAll(c.Value, ",", "."), 32)
			if value > maxValue {
				maxValue = value

				maxRate.Name = c.Name
				maxRate.Value = float32(value)
				t, _ := time.Parse("02.01.2006", v.Date)
				maxRate.Date = t
			}
			if value < minValue {
				minValue = value

				minRate.Name = c.Name
				minRate.Value = float32(value)
				t, _ := time.Parse("02.01.2006", v.Date)
				minRate.Date = t
			}
		}
	}

	return maxRate, minRate
}

func AverageRUB(ValCursInfo map[string]ValCurs) []AverageRate {
	currencies := make(map[string]*AverageRate, 0)

	for _, v := range ValCursInfo {
		for _, p := range v.Currencies {
			value, _ := strconv.ParseFloat(strings.ReplaceAll(p.Value, ",", "."), 32)
			if _, ok := currencies[p.Name]; !ok {
				currencies[p.Name] = &AverageRate{
					Name:    p.Name,
					Value:   value,
					Counter: 1,
				}
			} else {
				currencies[p.Name].Value += value
				currencies[p.Name].Counter++
			}
		}
	}

	curs := make([]AverageRate, 0)
	for k, v := range currencies {
		curs = append(curs, AverageRate{
			Name:  k,
			Value: v.Value / float64(v.Counter),
		})
	}

	return curs
}

func main() {
	logger := logrus.New()

	ValCursInfo := make(map[string]ValCurs, 0)
	now := time.Now()
	for i := 0; i < 90; i++ {
		t := now.AddDate(0, 0, -i)
		url := fmt.Sprintf("http://www.cbr.ru/scripts/XML_daily_eng.asp?date_req=%s", t.Format("02/01/2006"))
		resp, err := http.Get(url)
		if err != nil {
			logger.Log(logrus.DebugLevel, err)
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
	max, min := MaxMinCurrencyRate(ValCursInfo)
	averageRuble := AverageRUB(ValCursInfo)
	logger.Println(max, min, averageRuble)
}
