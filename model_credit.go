package echargeSDK

type RequestCreditCard struct {
	RemoteIP             string               `json:"remoteIp" bson:"remote_ip"`
	Customer             string               `json:"customer" bson:"customer"`
	CreditCard           CreditCardInfo       `json:"creditCardInfo" bson:"credit_card_info"`
	CreditCardHolderInfo CreditCardHolderInfo `json:"creditCardHolderInfo" bson:"credit_card_holder_info"`
}

type CreditCardInfo struct {
	HolderName  string `json:"holderName" bson:"holder_name"`
	Number      string `json:"number" bson:"number"`
	ExpiryMonth string `json:"expiryMonth" bson:"expiry_month"`
	ExpiryYear  string `json:"expiryYear" bson:"expiry_year"`
	CCV         string `json:"ccv" bson:"ccv"`
}

type CreditCardHolderInfo struct {
	Name              string `json:"name" bson:"name"`
	Email             string `json:"email" bson:"email"`
	Document          string `json:"cpfCnpj" bson:"document"`
	PostalCode        string `json:"postalCode" bson:"postal_code"`
	AddressNumber     string `json:"addressNumber" bson:"address_number"`
	AddressComplement string `json:"addressComplement" bson:"address_complement"`
	Phone             string `json:"phone" bson:"phone"`
	MobilePhone       string `json:"mobilePhone" bson:"mobile_phone"`
}

type ResponseCreditCard struct {
	CreditCardNumber string `json:"creditCardNumber" bson:"creditCardNumber"`
	CreditCardBrand  string `json:"creditCardBrand" bson:"credit_card_brand"`
	CreditCardToken  string `json:"creditCardToken" bson:"credit_card_token"`
}
