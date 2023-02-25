package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"time"
)

type OperationObject struct {
	Type  interface{} `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Id    interface{} `json:"id,omitempty"`
	Date  interface{} `json:"created_at,omitempty"`
}
type Billing struct {
	Company   string          `json:"company,omitempty"`
	Operation OperationObject `json:"operation,omitempty"`
	Type      interface{}     `json:"type,omitempty"`
	Value     interface{}     `json:"value,omitempty"`
	Id        interface{}     `json:"id,omitempty"`
	Date      interface{}     `json:"created_at,omitempty"`
}

type CompanyStats struct {
	Company              string        `json:"company"`
	ValidOperationsCount int           `json:"valid_operations_count"`
	Balance              int           `json:"balance"`
	InvalidOperations    []interface{} `json:"invalidOperations,omitempty"`
}

const EnvVarFile = "file"

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Failed to open file", err)
	}
	return file
}

func handleInput() *os.File {
	var flagFile = flag.String("file", "", ".json filename expected")
	flag.Parse()

	if *flagFile != "" {
		return openFile(*flagFile)
	}

	if envFile, exist := os.LookupEnv(EnvVarFile); exist != false {
		return openFile(envFile)
	}

	var userFile string
	if _, err := fmt.Scan(&userFile); err != nil {
		log.Fatalln("Failed to scan filename")
	} else {
		return openFile(userFile)
	}

	return nil
}

func getType(operationType ...interface{}) string {
	for _, arg := range operationType {
		if arg != nil {
			switch _type := arg.(type) {
			case string:
				if _type == "+" || _type == "income" {
					return "+"
				}
				if _type == "outcome" || _type == "-" {
					return "-"
				}
			default:
				return ""
			}
		}
	}
	return ""
}

func isInteger(f float64) bool {
	return math.Abs(f-float64(int(f))) <= 1e-16
}

func getValue(operationValue ...interface{}) int {
	for _, arg := range operationValue {
		if arg != nil {
			switch value := arg.(type) {
			case string:
				if strValue, err := strconv.Atoi(value); err != nil {
					return 0
				} else {
					return strValue
				}
			case float64:
				if !isInteger(value) {
					return 0
				}
				return int(value)
			default:
				return 0
			}
		}
	}
	return 0
}

func getId(operationId ...interface{}) string {
	for _, arg := range operationId {
		if arg != nil {
			switch id := arg.(type) {
			case float64:
				return strconv.Itoa(int(id))
			case string:
				return id
			default:
				return ""
			}
		}
	}
	return ""
}

func getDate(operationDate ...interface{}) time.Time {
	for _, arg := range operationDate {
		if arg != nil {
			switch date := arg.(type) {
			case string:
				loc, err := time.LoadLocation("Europe/Moscow")
				if err != nil {
					log.Fatalln("Wrong location", err)
				}
				time_, err := time.ParseInLocation(time.RFC3339, date, loc)
				if err != nil {

					return time.Time{}
				}
				return time_
			}
		}
	}
	return time.Time{}
}

func updateBalance(stats *CompanyStats, type_ string, value int) {
	if type_ == "+" {
		stats.Balance += value
	} else {
		stats.Balance -= value
	}
}

func writeFile(stats *[]CompanyStats) {
	file, err := os.Create("out.json")
	if err != nil {
		log.Fatalln(err)
	}

	enc := json.NewEncoder(file)
	enc.SetIndent("", "\t")
	if err = enc.Encode(&stats); err != nil {
		log.Fatalln(err)
	}

	file.Close()
}

func handleCompanies(billInfo []Billing) {
	sort.Slice(billInfo, func(i, j int) bool {
		return getDate(billInfo[i].Operation.Date,
			billInfo[i].Date).Before(getDate(billInfo[j].Operation.Date, billInfo[j].Date))
	})
	companies := make(map[string][]Billing)
	for _, v := range billInfo {
		if v.Company != "" && getId(v.Operation.Id, v.Id) != "" &&
			!(getDate(v.Operation.Date, v.Date).IsZero()) {
			companies[v.Company] = append(companies[v.Company], v)
		}
	}

	stats := make([]CompanyStats, len(companies), len(companies))
	c := 0
	for k, val := range companies {
		stats[c] = CompanyStats{Company: k}
		for _, v := range val {
			type_ := getType(v.Operation.Type, v.Type)
			value := getValue(v.Operation.Value, v.Value)
			if type_ == "" || value == 0 {
				stats[c].InvalidOperations = append(stats[c].InvalidOperations,
					getId(v.Operation.Id, v.Id))
				continue
			}
			stats[c].ValidOperationsCount++
			updateBalance(&stats[c], type_, value)
		}
		c++
	}
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Company < stats[j].Company
	})
	writeFile(&stats)
}

func main() {
	if file := handleInput(); file != nil {
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Println("Failed to close file")
			}
		}(file)
		var billInfo []Billing
		var decoder *json.Decoder = json.NewDecoder(file)
		if err := decoder.Decode(&billInfo); err != nil {
			log.Fatalln("Failed to decode", err)
		}
		handleCompanies(billInfo)
	} else {
		log.Fatalln("No file to handle")
	}
}
