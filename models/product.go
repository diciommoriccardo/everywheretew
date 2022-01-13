package models

type Product struct {
	name           string
	tags           []string
	description    string
	conversationId string
	externalId     string
	source         string
}

// globalSearchByProductName returns the number of tweets present globally on the twitter platform (useful in order to calculate the product score)
func globalSearchByProductName(product Product) int {
	nTweets := 0

	// non riesco a trovare una api idonea allo scopo nella api docs v2, per il momento uso la "recent search api" gratuita che mostra solo i tweet recenti
	return nTweets
}
