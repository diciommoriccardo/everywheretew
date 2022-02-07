package models

import (
	"fmt"
	"time"
)

type Product struct {
	Name           string
	Tags           []string
	Description    string
	ConversationId string
	ExternalId     string
	Source         string
	Country        string
	AvgAge         float32
	TotalEarnings  float32
	PublishDate    time.Time
}

func (t Product) String() string {
	return fmt.Sprintf("\n{"+
		"\n\tname: %s,"+
		"\n\ttags: %s, "+
		"\n}", t.Name, t.Tags)
}

// globalSearchByProductName returns the number of tweets present globally on the twitter platform (useful in order to calculate the product score)
func globalSearchByProductName(product Product) int {
	nTweets := 0

	// non riesco a trovare una api idonea allo scopo nella api docs v2, per il momento uso la "recent search api" gratuita che mostra solo i tweet recenti
	return nTweets
}

var InMemoryProductsDb = map[string][]Product{
	"61d9c7f193247cd87083247c": []Product{
		{
			Name:          "Prodotto 1",
			Tags:          []string{"Tag1"},
			Description:   "",
			Country:       "BARI",
			AvgAge:        34,
			TotalEarnings: 480,
			PublishDate:   time.Date(2021, 12, 04, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:          "Prodotto 4",
			Tags:          []string{"Tag4", "Tag2"},
			Description:   "",
			Country:       "FOGGIA",
			AvgAge:        32,
			TotalEarnings: 120,
			PublishDate:   time.Date(2021, 12, 05, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:          "Prodotto 8",
			Tags:          []string{"Tag8", "Tag1"},
			Description:   "",
			Country:       "ROMA",
			AvgAge:        45,
			TotalEarnings: 800,
			PublishDate:   time.Date(2021, 12, 01, 0, 0, 0, 0, time.UTC),
		},
	},
	"61d9d046b0556f25834f5c7a": []Product{
		{
			Name:          "Prodotto 10",
			Tags:          []string{"Tag10"},
			Description:   "",
			Country:       "PESCARA",
			AvgAge:        43,
			TotalEarnings: 1250,
			PublishDate:   time.Date(2021, 12, 14, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:          "Prodotto 40",
			Tags:          []string{"Tag40", "Tag21"},
			Description:   "",
			Country:       "FIRENZE",
			AvgAge:        56,
			TotalEarnings: 1760,
			PublishDate:   time.Date(2021, 12, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:          "Prodotto 80",
			Tags:          []string{"Tag80", "Tag10"},
			Description:   "",
			Country:       "TORINO",
			AvgAge:        26,
			TotalEarnings: 570,
			PublishDate:   time.Date(2021, 12, 01, 0, 0, 0, 0, time.UTC),
		},
	},
	"61d9d07b63564d1ec3798083": []Product{
		{
			Name:          "Prodotto 15",
			Tags:          []string{"Tag15"},
			Description:   "",
			Country:       "BARI",
			AvgAge:        39,
			TotalEarnings: 950,
			PublishDate:   time.Date(2021, 12, 13, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:          "Prodotto 80",
			Tags:          []string{"Tag80", "Tag10"},
			Description:   "",
			Country:       "TORINO",
			AvgAge:        26,
			TotalEarnings: 570,
			PublishDate:   time.Date(2021, 12, 01, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:          "Prodotto 16",
			Tags:          []string{"Tag16", "Tag1"},
			Description:   "",
			Country:       "MATERA",
			AvgAge:        32,
			TotalEarnings: 1980,
			PublishDate:   time.Date(2022, 01, 10, 0, 0, 0, 0, time.UTC),
		},
	},
}
