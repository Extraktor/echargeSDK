package echargeSDK

type RequestCharge struct {
	Customer    string  `json:"customer,omitempty" bson:"costumer,omitempty"`
	BillingType string  `json:"billingType" bson:"billing_type"`
	DueDate     string  `json:"dueDate" bson:"due_date"`
	Description string  `json:"description" bson:"description"`
	Value       float64 `json:"value" bson:"value"`
	Split       []Slipt `json:"split" bson:"split"`
}

type Slipt struct {
	WalletID          string  `json:"walletid" bson:"wallet_id"`
	ExternalReference string  `json:"external_reference" bson:"external_reference"`
	Description       string  `json:"description" bson:"description"`
	FixedValue        float64 `json:"fixedValue" bson:"fixed_value"`
}

type ResponseCharge struct {
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
