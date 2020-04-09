package insightly

import (
	"fmt"
	"strconv"
)

// OpportunityCategory stores OpportunityCategory from Insightly
//
type OpportunityCategory struct {
	CATEGORY_ID      int    `json:"CATEGORY_ID"`
	CATEGORY_NAME    string `json:"CATEGORY_NAME"`
	ACTIVE           bool   `json:"ACTIVE"`
	BACKGROUND_COLOR string `json:"BACKGROUND_COLOR"`
}

// GetOpportunityCategories returns all opportunitycategories
//
func (i *Insightly) GetOpportunityCategories() ([]OpportunityCategory, error) {
	return i.GetOpportunityCategoriesInternal()
}

// GetOpportunityCategoriesInternal is the generic function retrieving opportunitycategories from Insightly
//
func (i *Insightly) GetOpportunityCategoriesInternal() ([]OpportunityCategory, error) {
	urlStr := "%sOpportunityCategories?skip=%s&top=%s"
	skip := 0
	top := 500
	rowCount := top

	opportunityCategories := []OpportunityCategory{}

	for rowCount >= top {
		url := fmt.Sprintf(urlStr, i.apiURL, strconv.Itoa(skip), strconv.Itoa(top))
		//fmt.Println(url)

		oc := []OpportunityCategory{}

		err := i.Get(url, &oc)
		if err != nil {
			return nil, err
		}

		for _, o := range oc {
			opportunityCategories = append(opportunityCategories, o)
		}

		rowCount = len(oc)
		skip += top
	}

	if len(opportunityCategories) == 0 {
		opportunityCategories = nil
	}

	return opportunityCategories, nil
}
