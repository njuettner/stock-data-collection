package earnings

type Earnings struct {
	Symbol         string `json:"symbol"`
	AnnualEarnings []struct {
		FiscalDateEnding string `json:"fiscalDateEnding"`
		ReportedEPS      string `json:"reportedEPS"`
	} `json:"annualEarnings"`
	QuarterlyEarnings []struct {
		FiscalDateEnding   string `json:"fiscalDateEnding"`
		ReportedDate       string `json:"reportedDate"`
		ReportedEPS        string `json:"reportedEPS"`
		EstimatedEPS       string `json:"estimatedEPS"`
		Surprise           string `json:"surprise"`
		SurprisePercentage string `json:"surprisePercentage"`
	} `json:"quarterlyEarnings"`
}

type Calendar []struct {
	Currency         string `json:"currency"`
	Estimate         string `json:"estimate"`
	FiscalDateEnding string `json:"fiscalDateEnding"`
	Name             string `json:"name"`
	ReportDate       string `json:"reportDate"`
	Symbol           string `json:"symbol"`
}
