package entity

type Filter struct {
	Limit        *int
	Offset       *int
	CategoryId   *int64
	Language     *string
	DiscountOnly *bool
	PriceMin     *float64
	PriceMax     *float64
	Rating       *int8
	SellerId     *int64
	BrandId      *int64
	Search       *string
	SortBy       *string
	SortOrder    *string
	PopularOnly  *bool
	Status       *bool
	Order        *string
}
