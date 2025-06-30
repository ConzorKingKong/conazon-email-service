package types

import (
	checkout "github.com/conzorkingkong/conazon-checkout/types"
	user "github.com/conzorkingkong/conazon-users-and-auth/types"
)

type Email struct {
	Checkout checkout.Checkout `json:"checkout"`
	User     user.User         `json:"user"`
}
