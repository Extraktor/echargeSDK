package echargeSDK

type RequestPayment struct {
	CreditCardToken string `json:"creditCardToken" bson:"credit_card_token"`
}

type ResponsePayment struct {
	ChargeID    string  `json:"id" bson:"id"`
	Object      string  `json:"object" bson:"object"`
	Customer    string  `json:"customer" bson:"customer"`
	DateCreated string  `json:"dateCreated" bson:"date_created"`
	Description string  `json:"description" bson:"description"`
	Status      string  `json:"status" bson:"status"`
	BillingType string  `json:"billingType" bson:"billing_type"`
	Value       float64 `json:"value" bson:"value"`
	NetValue    float64 `json:"netValue" bson:"net_value"`
}
