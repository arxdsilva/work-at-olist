package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/arxdsilva/olist/record"

	"github.com/arxdsilva/olist/bill"
	"github.com/labstack/echo"
)

// Bill calculates a specific bill of a subscriber
// If a month is not specified, It'll be the last
// closed month
// month and year are strings of integers
func (s *Server) Bill(c echo.Context) (err error) {
	sub := c.Param("subscriber")
	month := c.QueryParam("month")
	year := c.QueryParam("year")
	if (month == "") || (year == "") {
		month = string(strconv.Itoa(int(time.Now().Month() - 1)))
		year = string(strconv.Itoa(time.Now().Year()))
	}
	bill := bill.New(month, year, sub)
	storedBill, err := s.Storage.BillFromID(bill.ID)
	// bill already exists
	if err == nil {
		calls, err := s.Storage.CallsFromBillID(storedBill.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		storedBill.Calls = calls
		return c.JSON(http.StatusOK, storedBill)
	}
	// create new bill
	records, err := s.Storage.RecordsFromBill(bill)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	m, err := strconv.Atoi(month)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	bill.ID = fmt.Sprintf("%s%v%s", sub, month, year)
	calls, err := callsFromRecords(records, m, bill.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	bill.Calls = calls
	err = s.Storage.SaveCalls(calls)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	err = s.Storage.SaveBill(bill)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, bill)
}

func callsFromRecords(rs []record.Record, monthOfReference int, billID string) (cs []bill.Call, err error) {
	f := filterRecordsPeriod(rs, monthOfReference)
	for _, v := range f {
		c, errC := callFromRecords(v)
		if errC != nil {
			return nil, errC
		}
		c.BillID = billID
		cs = append(cs, c)
	}
	return
}

func callFromRecords(rs []record.Record) (c bill.Call, err error) {
	buff := make(map[string]record.Record)
	for _, v := range rs {
		buff[v.Type] = v
	}
	startTime, err := time.Parse(time.RFC3339, buff["start"].TimeStamp)
	if err != nil {
		return
	}
	endTime, err := time.Parse(time.RFC3339, buff["end"].TimeStamp)
	if err != nil {
		return
	}
	price, err := bill.Calculate(startTime, endTime)
	if err != nil {
		return
	}
	duration := endTime.Sub(startTime)
	c.CallPrice = price
	c.CallStartDate = startTime.Day()
	c.Destination = buff["start"].Destination
	c.CallStartTime = startTime.Format(time.Kitchen)
	c.CallDuration = duration.String()
	return
}

func filterRecordsPeriod(rs []record.Record, monthOfReference int) (rFiltered map[string][]record.Record) {
	rFiltered = make(map[string][]record.Record)
	recordsMap := make(map[string][]record.Record)
	for _, r := range rs {
		recordsMap[r.CallID] = append(recordsMap[r.CallID], r)
	}
	for k, v := range recordsMap {
		if (v[0].Type == "end") && (v[0].Month == monthOfReference) {
			rFiltered[k] = append(rFiltered[k], v[0])
			rFiltered[k] = append(rFiltered[k], v[1])
		} else if (v[1].Type == "end") && (v[1].Month == monthOfReference) {
			rFiltered[k] = append(rFiltered[k], v[0])
			rFiltered[k] = append(rFiltered[k], v[1])
		}
	}
	return
}
