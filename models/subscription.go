package models

type Subscription struct {
	SubscriptionId string `json:"subscriptionID" form:"subscriptionID" bson:"subscriptionID"`
	PlanId         string `json:"planID" form:"planID" bson:"planID"`
}
