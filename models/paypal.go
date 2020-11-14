package models

type PaypalToken struct {
	AccessToken string `json:"access_token" form:"access_token" bson:"access_token"`
	AppId       string `json:"app_id" form:"app_id" bson:"app_id"`
	ExpiresIn   int    `json:"expires_in" form:"expires_in" bson:"expires_in"`
}

type ProductList struct {
	Products []Product `json:"products" form:"products" bson:"products"`
}

type Product struct {
	Id   string `json:"id" form:"id" bson:"id"`
	Name string `json:"name" form:"name" bson:"name"`
}

type PlanList struct {
	Plans []Plan `json:"plans" form:"plans" bson:"plans"`
}

type Plan struct {
	Id            string        `json:"id" form:"id" bson:"id"`
	ProductId     string        `json:"product_id" form:"product_id" bson:"product_id"`
	Name          string        `json:"name" form:"name" bson:"name"`
	Status        string        `json:"status" form:"status" bson:"status"`
	BillingCycles []PlanDetails `json:"billing_cycles" form:"billing_cycles" bson:"billing_cycles"`
}

type SubscriptionDetails struct {
	Id     string `json:"id" form:"id" bson:"id"`
	PlanId string `json:"plan_id" form:"plan_id" bson:"plan_id"`
	Status string `json:"status" form:"status" bson:"status"`
}

type PlanDetails struct {
} //TODO only if needed
