package myml

type User struct {
	ID               int64    `json:"id"`
	Nickname         string `json:"nickname"`
	RegistrationDate string `json:"registration_date"`
	CountryID        string `json:"country_id"`
	SiteID           string      `json:"site_id"`
	Permalink        string      `json:"permalink"`
	SellerReputation struct {
		LevelID           interface{} `json:"level_id"`
		PowerSellerStatus interface{} `json:"power_seller_status"`
		Transactions      struct {
			Canceled  int    `json:"canceled"`
			Completed int    `json:"completed"`
			Period    string `json:"period"`
			Ratings   struct {
				Negative int `json:"negative"`
				Neutral  int `json:"neutral"`
				Positive int `json:"positive"`
			} `json:"ratings"`
			Total int `json:"total"`
		} `json:"transactions"`
	} `json:"seller_reputation"`
	BuyerReputation struct {
		Tags []interface{} `json:"tags"`
	} `json:"buyer_reputation"`
	Status struct {
		SiteStatus string `json:"site_status"`
	} `json:"status"`
	SiteInfo Site
	Categories Category
}


type Site struct {
	ID                 string        `json:"id"`
	Name               string        `json:"name"`
	CountryID          string        `json:"country_id"`
	SaleFeesMode       string        `json:"sale_fees_mode"`
	MercadopagoVersion int           `json:"mercadopago_version"`
	DefaultCurrencyID  string        `json:"default_currency_id"`
	ImmediatePayment   string        `json:"immediate_payment"`
	PaymentMethodIds   []interface{} `json:"payment_method_ids"`
	Settings           struct {
	} `json:"settings"`
}

type Category []struct {
	ID string `json:"id"`
	Name string `json:"name"`

}