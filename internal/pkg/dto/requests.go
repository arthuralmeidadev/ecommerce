package dto

type OrderRequestDTO struct {
	Id                 string
	StoreId            string
	CompanyId          string
	Subtotal           float64
	ShippingFee        float64
	Tax                float64
	Discount           float64
	DiscountPercentage float64
	Total              float64
	Currency           string
	Items              struct {
		Products []struct {
			Id           string
			SKU          string
			Description  string
			Kind         string
			Quantity     int16
			PricePerUnit float64
		}
		Coupons []struct {
			Id                 string
			Description        string
			Discount           float64
			DiscountPercentage float64
		}
	}
	Customer struct {
		Id         string
		EntityKind string
		FullName   string
		FirstName  string
		LastName   string
		Address    struct {
			Street  string
			Number  int8
			City    string
			State   string
			ZipCode string
			Country string
		}
	}
	Payments []struct {
		Total      float64
		Type       string
		Save       bool
		CreditCard struct {
			CardBrand           string
			CardNumber          string
			CardHolderName      string
			Expiry              string
			CVV                 string
			Installments        int8
			TotalPerInstallment float64
		}
		DebitCard struct {
			CardBrand      string
			CardNumber     string
			CardHolderName string
			Expiry         string
			CVV            string
		}
		BankTransfer struct {
			BankName      string
			AccountNumber string
			RoutingNumber string
			SwiftCode     string
		}
		DigitalWallet struct {
			Provider            string
			WalletId            string
			Installments        int8
			TotalPerInstallment float64
		}
	}
	ShippingAddress struct {
		Street  string
		Number  int8
		City    string
		State   string
		ZipCode string
		Country string
	}
}
