package services

import (
	"encoding/json"
	"net/http"
	"paypal-starter-backend2/m/models"

	"go.mongodb.org/mongo-driver/bson"
)

var subscriptions_database = "demopaypal"
var subscriptions_collection = "subscriptions"

func PostSubscription(w http.ResponseWriter, r *http.Request) {

	var item models.Subscription

	_ = json.NewDecoder(r.Body).Decode(&item)

	itemmodel := mymodel{I: item, DatabaseName: subscriptions_database, Collection: subscriptions_collection}
	_, err := itemmodel.postInterface(w, r)
	if err != nil {
		Log.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}

}

func GetSubscriptions(w http.ResponseWriter, r *http.Request) {

	var item models.Subscription

	filter := bson.M{}

	itemmodel := mymodel{I: &item, F: filter, DatabaseName: subscriptions_database, Collection: subscriptions_collection}

	itemmodel.getInterfaces(w, r)
}

func DeleteSubscription(w http.ResponseWriter, r *http.Request) {

	keys, _ := r.URL.Query()["subscriptionID"]
	key := keys[0]

	var filter interface{}
	if key != "" {

		filter = bson.M{"subscriptionID": key}
	} else {
		filter = bson.M{}

	}

	itemmodel := mymodel{F: filter, DatabaseName: subscriptions_database, Collection: subscriptions_collection}
	_, err := itemmodel.deleteInterface(w, r)
	if err != nil {
		Log.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}

}
