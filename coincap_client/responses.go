package coincap

import "fmt"

type assetsResponse struct { // слайс со всеми криптовалютами
	Assets    []Asset `json:"data"`
	Timestamp int64   `json:"timestamp"`
}

type assetResponse struct { // одна криптовалюта по id
	Asset     Asset `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type Asset struct { // описание столбцов каждой криптовалюты
	ID            string `json:"id"`
	Rank          string `json:"rank"`
	Symbol        string `json:"symbol"`
	Name          string `json:"name"`
	Supply        string `json:"supply"`
	MaxSupply     string `json:"maxSupply"`
	MarketCapUsd  string `json:"marketCapUsd"`
	VolumeUsd24Hr string `json:"volumeUsd24Hr"`
	PriceUsd      string `json:"priceUsd"`
}

func (d Asset) Info() string {
	return fmt.Sprintf("Id ~ %s | Rank ~ %s | Symbol ~ %s | Name ~ %s | Supply ~ %s | MaxSupply ~ %s | MaketCapUsd ~ %s | VolumeUsd24Hr ~ %s | PriceUsd ~ %s|\n",
		d.ID, d.Rank, d.Symbol, d.Name, d.Supply, d.MaxSupply, d.MarketCapUsd, d.VolumeUsd24Hr, d.PriceUsd)
}
