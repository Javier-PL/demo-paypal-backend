package services

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"paypal-starter-backend2/m/models"
	"strconv"
	"strings"
)

func RequestToken(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var productId string
		keys, _ := r.URL.Query()["product_id"]
		if len(keys) > 0 {
			productId = keys[0]
		}

		var planId string
		keys, _ = r.URL.Query()["plan_id"]
		if len(keys) > 0 {
			planId = keys[0]
		}

		var subscriptionID string
		keys, _ = r.URL.Query()["subscriptionID"]
		if len(keys) > 0 {
			subscriptionID = keys[0]
		}

		form := url.Values{}
		form.Set("grant_type", "client_credentials")

		client := &http.Client{}

		r, err := http.NewRequest("POST", "https://api.sandbox.paypal.com/v1/oauth2/token", strings.NewReader(form.Encode()))
		if err != nil {
			Log.Error(err)
		}

		r.SetBasicAuth(os.Getenv("CLIENT_ID"), os.Getenv("SECRET"))
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
		r.Header.Add("Authorization", strconv.Itoa(len(form.Encode())))

		res, err := client.Do(r)
		if err != nil {
			Log.Error(err)
		}

		defer res.Body.Close()

		var paypalToken models.PaypalToken

		json.NewDecoder(res.Body).Decode(&paypalToken)

		if err == nil {

			contextValues := map[string]string{
				"access_token":   paypalToken.AccessToken,
				"product_id":     productId,
				"subscriptionID": subscriptionID,
				"plan_id":        planId,
			}
			ctx := context.WithValue(r.Context(), "contextValues", contextValues)
			next(w, r.WithContext(ctx))

		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		return

	})
}

func GetProducts(w http.ResponseWriter, r *http.Request) {

	contextValues := r.Context().Value("contextValues").(map[string]string) //.(string)
	accessToken := contextValues["access_token"]

	client := &http.Client{}

	r, err := http.NewRequest("GET", "https://api.sandbox.paypal.com/v1/catalogs/products?page_size=20&page=1&total_required=true", nil)
	if err != nil {
		Log.Error(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+accessToken)

	res, err := client.Do(r)
	if err != nil {
		Log.Error(err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var productList models.ProductList

	err = json.Unmarshal(data, &productList)
	if err != nil {
		Log.Error(err)
		http.Error(w, err.Error(), 500)
	}

	respBody, err := json.Marshal(productList)
	if err != nil {
		Log.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(respBody)
}

func GetPlans(w http.ResponseWriter, r *http.Request) {

	contextValues := r.Context().Value("contextValues").(map[string]string) //.(string)
	accessToken := contextValues["access_token"]
	productId := contextValues["product_id"]

	client := &http.Client{}

	r, err := http.NewRequest("GET", "https://api.sandbox.paypal.com/v1/billing/plans?product_id="+productId+"&page_size=20&page=1&total_required=true", nil)
	if err != nil {
		Log.Error(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+accessToken)

	res, err := client.Do(r)
	if err != nil {
		Log.Error(err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var planList models.PlanList

	err = json.Unmarshal(data, &planList)
	if err != nil {
		Log.Error(err)
		http.Error(w, err.Error(), 500)
	}

	respBody, err := json.Marshal(planList)
	if err != nil {
		Log.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(respBody)
}

func CancelSubscription(w http.ResponseWriter, r *http.Request) {

	contextValues := r.Context().Value("contextValues").(map[string]string) //.(string)
	accessToken := contextValues["access_token"]
	subscriptionID := contextValues["subscriptionID"]

	client := &http.Client{}

	r, err := http.NewRequest("POST", "https://api.sandbox.paypal.com/v1/billing/subscriptions/"+subscriptionID+"/cancel", nil)
	if err != nil {
		Log.Error(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+accessToken)

	res, err := client.Do(r)
	if err != nil {
		Log.Error(err)
	}
	defer res.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(res.Body)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(string(bodyBytes)))
}

func GetSubscriptionDetails(w http.ResponseWriter, r *http.Request) {

	contextValues := r.Context().Value("contextValues").(map[string]string) //.(string)
	accessToken := contextValues["access_token"]
	subscriptionID := contextValues["subscriptionID"]

	client := &http.Client{}

	r, err := http.NewRequest("GET", "https://api.sandbox.paypal.com/v1/billing/subscriptions/"+subscriptionID, nil)
	if err != nil {
		Log.Error(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+accessToken)

	res, err := client.Do(r)
	if err != nil {
		Log.Error(err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var subscriptionDetails models.SubscriptionDetails

	err = json.Unmarshal(data, &subscriptionDetails)
	if err != nil {
		Log.Error(err)
		http.Error(w, err.Error(), 500)
	}

	respBody, err := json.Marshal(subscriptionDetails)
	if err != nil {
		Log.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(respBody)
}

func GetPlanDetails(w http.ResponseWriter, r *http.Request) {

	contextValues := r.Context().Value("contextValues").(map[string]string) //.(string)
	accessToken := contextValues["access_token"]
	planID := contextValues["plan_id"]

	client := &http.Client{}

	r, err := http.NewRequest("GET", "https://api.sandbox.paypal.com/v1/billing/plans/"+planID, nil)
	if err != nil {
		Log.Error(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+accessToken)

	res, err := client.Do(r)
	if err != nil {
		Log.Error(err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var plan models.Plan

	err = json.Unmarshal(data, &plan)
	if err != nil {
		Log.Error(err)
		http.Error(w, err.Error(), 500)
	}

	respBody, err := json.Marshal(plan)
	if err != nil {
		Log.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(respBody)
}
