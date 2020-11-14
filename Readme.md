# Paypal demo implementation with Ionic

Navigate to: https://demo-paypal-frontend.herokuapp.com <br>
Login test data: <br>
Email: javier.perez.losada-buyer@gmail.com <br>
Password: abcd1234

# Project description.

Implementation of the Paypal gateway to work in a web as as well as in Android/iOS devices. Single payments as well as Subscriptions have been implemented. 

# Technical details

In order to implement the paypal gateway, a Sandbox test account has been used. As the communication with the Paypal API successes using an access token that enables full control, I wrote a small API to act as a proxy in order to hide the access token and other sensitive information from the frontend logic.

# Basic Functionalities

- Single payments using Paypal
- Subscribe to subscription plans using Paypal
- inApp subscription cancelation


# Project structure

- Proxy API and Web frontend hosted in Heroku.
- MongoDB hosted in MongoDB Atlas



