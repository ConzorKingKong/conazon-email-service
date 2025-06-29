package main

import "time"

type User struct {
	ID       int    `json:"-"`
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	Picture  string `json:"picture"`
	GoogleID int    `json:"-"`
}

type Product struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	MainImage   string    `json:"mainImage"`
	Category    string    `json:"category"`
	Price       float32   `json:"price"`
	Quantity    int       `json:"quantity,omitempty"`
	Author      string    `json:"author,omitempty"`
}

type Cart struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	ProductID int    `json:"productId"`
	Quantity  int    `json:"quantity"`
	Status    string `json:"status"`
}

type IdTokenPayload struct {
	Iss            string `json:"iss"`
	Azp            string `json:"azp"`
	Aud            string `json:"aud"`
	Sub            string `json:"sub"`
	Hd             string `json:"hd"`
	Email          string `json:"email"`
	Email_verified bool   `json:"email_verified"`
	At_hash        string `json:"at_hash"`
	Name           string `json:"name"`
	Picture        string `json:"picture"`
	Given_name     string `json:"given_name"`
	Family_name    string `json:"family_name"`
	Iat            int    `json:"iat"`
	Exp            int    `json:"exp"`
}

type Checkout struct {
	Id             int       `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	UserId         int       `json:"userId"`
	TotalPrice     string    `json:"totalPrice"`
	BillingStatus  string    `json:"billingStatus"`
	ShippingStatus string    `json:"shippingStatus"`
	TrackingNumber string    `json:"trackingNumber"`
}

type MyJWT struct {
	Id int `json:"id"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type CheckoutResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    Checkout `json:"data"`
}

type CheckoutsResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []Checkout `json:"data"`
}

type Email struct {
	Checkout Checkout `json:"checkout"`
	User     User     `json:"user"`
}
