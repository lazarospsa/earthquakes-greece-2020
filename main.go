package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"time"
)

func readCSVFromURL(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	url := "https://lithos.geo.auth.gr/fdsnws/event/1/query?starttime=2020-01-01T00%3A00%3A00&endtime=2021-01-01T00%3A00%3A00&eventtype=earthquake&includeallmagnitudes=true&format=csv&formatted=true&nodata=404"
	data, err := readCSVFromURL(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Date&Time, Latitude, Longitude, Depth, Magnitude, Description")

	for idx, row := range data {
		// skip header
		if idx == 0 {
			continue
		}

		/*if idx == 100 {
			break
		}*/
		date := row[1]
		parsetime, _ := time.Parse(time.RFC3339, date)
		timeStamp := parsetime.Format("2006-01-02 15:04:05")
		fmt.Println(timeStamp, row[2], row[3], row[4], row[5], row[6])
		coord := row[2] + "," + row[3]
		println("https://www.google.gr/maps/@" + coord + ",18z")
		//println("https://www.google.gr/maps/@", row[2], ",", row[3], ",13z")
	}
}

//url := "https://lithos.geo.auth.gr/fdsnws/event/1/query?starttime=2019-01-01T00%3A00%3A00&endtime=2020-02-01T00%3A00%3A00&eventtype=earthquake&includeallmagnitudes=true&format=csv&formatted=true&nodata=404"

/*
			timeStampString := currentTime.Format("2006-01-02 15:04:05")
			layOut := "2006-01-02 15:04:05"
			timeStamp, err := time.Parse(layOut, timeStampString)


	fmt.Printf("Format: dd/mm/yy hh:MM:ss PM Mon - %s \n", t.In(loc).Format("02/01/06 03:04:05 PM Jan"))
	fmt.Printf("Format: dd/mm/yy hh:MM:ss PM Da Mon - %s \n", t.In(loc).Format("02/01/06 03:04:05 PM Mon Jan"))
	fmt.Printf("Format: dd/mm/yy hh:MM:ss PM Day Month - %s \n", t.In(loc).Format("02/01/06 03:04:05 PM Monay January"))
	fmt.Printf("Format: d/m/y h:M:s PM - %s\n", t.In(loc).Format("2/1/6 3:4:5 PM"))
	fmt.Printf("Format: _d/m/y h:M:s PM - %s\n", t.In(loc).Format("_2/1/6 3:4:5 PM"))
	fmt.Printf("Format: dd/mm/yy hh:MM:ss PM - %s \n", t.In(loc).Format("02/01/06 03:04:05 PM"))
	fmt.Printf("Format: dd/mm/yyyy hh:MM:ss PM - %s \n", t.In(loc).Format("02/01/2006 03:04:05 PM"))
	fmt.Printf("Format: dd/mm/yyyy hh:MM:ss.ms PM - %s \n", t.In(loc).Format("02/01/2006 03:04:05.000 PM"))
	fmt.Printf("Format: dd/mm/yyyy hh:MM:ss.000ms PM - %s \n", t.In(loc).Format("02/01/2006 03:04:05.000000 PM"))
	fmt.Printf("Format: dd/mm/yyyy hh:MM:ss.000000ms PM - %s \n", t.In(loc).Format("02/01/2006 03:04:05.000000000 PM"))
	fmt.Printf("Format dd/mm/yyyy hh:MM:ss TZName - %s\n", t.In(loc).Format("02/01/2006 15:04:05 MST"))
	fmt.Printf("Format dd/mm/yyyy hh:MM:ss Z - %s\n", t.In(loc).Format("02/01/2006 15:04:05 Z7"))
	fmt.Printf("Format dd/mm/yyyy hh:MM:ss Z - %s\n", t.In(loc).Format("02/01/2006 15:04:05 Z07"))
	fmt.Printf("Format dd/mm/yyyy hh:MM:ss ZZ - %s\n", t.In(loc).Format("02/01/2006 15:04:05 Z0700"))
	fmt.Printf("Format dd/mm/yyyy hh:MM:ss Z:Z - %s\n", t.In(loc).Format("02/01/2006 15:04:05 Z07:00"))
	fmt.Printf("Format dd/mm/yyyy hh:MM:ss Z:Z - %s\n", t.In(loc).Format("02/01/2006 15:04:05 -07:00"))
*/
/* Output:
Times:
Format: dd/mm/yy hh:MM:ss PM Mon - 09/11/16 10:56:30 AM Nov
Format: dd/mm/yy hh:MM:ss PM Da Mon - 09/11/16 10:56:30 AM Wed Nov
Format: dd/mm/yy hh:MM:ss PM Day Month - 09/11/16 10:56:30 AM Monay November
Format: d/m/y h:M:s PM - 9/11/6 10:56:30 AM
Format: _d/m/y h:M:s PM -  9/11/6 10:56:30 AM
Format: dd/mm/yy hh:MM:ss PM - 09/11/16 10:56:30 AM
Format: dd/mm/yyyy hh:MM:ss PM - 09/11/2016 10:56:30 AM
Format: dd/mm/yyyy hh:MM:ss.ms PM - 09/11/2016 10:56:30.952 AM
Format: dd/mm/yyyy hh:MM:ss.000ms PM - 09/11/2016 10:56:30.952216 AM
Format: dd/mm/yyyy hh:MM:ss.000000ms PM - 09/11/2016 10:56:30.952216972 AM
Format dd/mm/yyyy hh:MM:ss TZName - 09/11/2016 10:56:30 IST
Format dd/mm/yyyy hh:MM:ss Z - 09/11/2016 10:56:30 Z7
Format dd/mm/yyyy hh:MM:ss Z - 09/11/2016 10:56:30 +02
Format dd/mm/yyyy hh:MM:ss ZZ - 09/11/2016 10:56:30 +0200
Format dd/mm/yyyy hh:MM:ss Z:Z - 09/11/2016 10:56:30 +02:00
Format dd/mm/yyyy hh:MM:ss Z:Z - 09/11/2016 10:56:30 +02:00
*/
