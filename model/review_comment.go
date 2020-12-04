package model

type ReviewCommentCategory string

const (
	Fraud          ReviewCommentCategory = "Bedrägeri"
	Other                                = "Annat"
	Sales                                = "Telefonförsäljning"
	MarketResearch                       = "Marknadsundersökning"
)

func ParseReviewCommentCategory(categoryString string) ReviewCommentCategory {
	switch categoryString {
	case "Bedrägeri":
		return Fraud
	case "Telefonförsäljning":
		return Sales
	case "Marknadsundersökning":
		return MarketResearch
	default:
		return Other
	}
}

type ReviewComment struct {
	PostDate string                `json:"post_date"`
	Content  string                `json:"content"`
	Category ReviewCommentCategory `json:"category"`
}
