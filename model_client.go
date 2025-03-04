package echargeSDK

import (
	"time"
)

type RequestClient struct {
	Name     string     `json:"name" bson:"name"`
	Document string     `json:"cpfCnpj" bson:"document"`
	Email    string     `json:"email" bson:"email"`
	CreateAt *time.Time `json:"-" bson:"create_at"`
}

type ResponseClient struct {
	Costumer string `json:"id" bson:"id"`
	Object   string `json:"object" bson:"object"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Document string `json:"cpfCnpj,omitempty" bson:"document,omitempty"`
	Deleted  bool   `json:"deleted,omitempty" bson:"deleted,omitempty"`
}
