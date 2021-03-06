package mws

import (
	"encoding/xml"
	"log"
	"net/url"
	"time"

	"github.com/haoxins/tools/v2"
	"github.com/thoas/go-funk"
)

// ReportType The Amazon report type
type ReportType string

const (
	// SettlementReport ...
	SettlementReport ReportType = "_GET_V2_SETTLEMENT_REPORT_DATA_FLAT_FILE_V2_"
	// InventoryReport ...
	InventoryReport = "_GET_FLAT_FILE_OPEN_LISTINGS_DATA_"
)

// GetReportListParams ...
type GetReportListParams struct {
	ReportType ReportType
	StartTime  time.Time
	EndTime    time.Time
}

// GetReportList Get all report ids if has next
func (seller *Seller) GetReportList(params GetReportListParams) []string {
	nextToken := ""
	var ids []string

	for {
		result := seller.getReportList(params.ReportType, params.StartTime, params.EndTime, nextToken)
		for _, v := range result.ReportInfos {
			ids = append(ids, v.ReportID)
		}

		if !result.HasNext {
			break
		}

		nextToken = result.NextToken
	}

	uids := funk.UniqString(ids)

	return uids
}

// GetReportList Get seller report list by report type
func (seller *Seller) getReportList(reportType ReportType, startTime time.Time, endTime time.Time, nextToken string) *GetReportListResult {
	params := seller.genGetReportListParams(reportType, startTime, endTime, nextToken)

	raw, err := seller.requestReports(params)
	tools.PanicError(err)

	if nextToken == "" {
		data := GetReportListResponse{}
		err = xml.Unmarshal(raw, &data)
		if err != nil {
			// TODO
			log.Println(string(raw))
			return nil
		}

		return &data.GetReportListResult
	}

	data := GetReportListByNextTokenResponse{}
	err = xml.Unmarshal(raw, &data)
	if err != nil {
		// TODO
		log.Println(string(raw))
		return nil
	}

	return &data.GetReportListResult
}

// GetReportByID Get seller report by report id
func (seller *Seller) GetReportByID(reportID string) string {
	params := seller.genGetReportParams(reportID)

	raw, err := seller.requestReports(params)
	tools.PanicError(err)
	text := string(raw)
	return text
}

func (seller *Seller) genGetReportListParams(reportType ReportType, startTime time.Time, endTime time.Time, nextToken string) string {
	v := url.Values{}

	seller.addBasicParams(&v)

	if nextToken != "" {
		v.Add("Action", "GetReportListByNextToken")
		v.Add("NextToken", nextToken)
	} else {
		v.Add("Action", "GetReportList")
	}

	v.Add("ReportTypeList.Type.1", string(reportType))
	v.Add("AvailableFromDate", startTime.Format(time.RFC3339))
	v.Add("AvailableToDate", endTime.Format(time.RFC3339))
	v.Add("MaxCount", "100")
	v.Add("Version", "2009-01-01")

	s := v.Encode()

	return s
}

func (seller *Seller) genGetReportParams(reportID string) string {
	v := url.Values{}

	seller.addBasicParams(&v)

	v.Add("Action", "GetReport")
	v.Add("ReportId", reportID)
	v.Add("Version", "2009-01-01")

	s := v.Encode()

	return s
}

func (seller *Seller) requestReports(params string) ([]byte, error) {
	// According to the document, this should be POST
	// But, only GET works
	return seller.get(ReportsPath, params)
}
