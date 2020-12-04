package model

type Review struct {
	TelephoneNumber string          `json:"telephone_number"`
	SearchCount     int             `json:"search_count"`
	ReportedCount   int             `json:"reported_count"`
	LastSearch      string          `json:"last_search"`
	Comments        []ReviewComment `json:"comments"`
}
