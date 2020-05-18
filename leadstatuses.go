package insightly

import (
	"fmt"
	"strconv"
)

// LeadStatus stores LeadStatus from Insightly
//
type LeadStatus struct {
	LEAD_STATUS_ID int    `json:"LEAD_STATUS_ID"`
	LEAD_STATUS    string `json:"LEAD_STATUS"`
	DEFAULT_STATUS bool   `json:"DEFAULT_STATUS"`
	STATUS_TYPE    int    `json:"STATUS_TYPE"`
	FIELD_ORDER    int    `json:"FIELD_ORDER"`
}

// GetLeadStatuses returns all leadStatuses
//
func (i *Insightly) GetLeadStatuses() ([]LeadStatus, error) {
	urlStr := "%sLeadStatuses?skip=%s&top=%s"
	skip := 0
	top := 500
	rowCount := top

	leadStatuses := []LeadStatus{}

	for rowCount >= top {
		url := fmt.Sprintf(urlStr, i.apiURL, strconv.Itoa(skip), strconv.Itoa(top))
		//fmt.Println(url)

		ls := []LeadStatus{}

		err := i.Get(url, &ls)
		if err != nil {
			return nil, err
		}

		for _, l := range ls {
			leadStatuses = append(leadStatuses, l)
		}

		rowCount = len(ls)
		skip += top
	}

	if len(leadStatuses) == 0 {
		leadStatuses = nil
	}

	return leadStatuses, nil
}
