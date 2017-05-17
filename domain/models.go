package domain

type Product struct {
	GTIN        string `json:"gtin,omitempty"`
	Description string `json:"description,omitempty"`
	Contents    *struct {
		Quantity    float64 `json:"quantity,omitempty"`
		QuantityUOM string  `json:"quantityUom,omitempty"`
		AvgMeasure  string  `json:"avgMeasure,omitempty"`
		NetContents string  `json:"netContents,omitempty"`
	} `json:"qtyContents"`
	ProductCharacteristics *struct {
		IsFood  bool `json:"isFood"`
		IsDrink bool `json:"isDrink"`
	}
}
