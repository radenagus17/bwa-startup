package transaction

import "bwastartup/user"

type GetTransactionsCampaignInput struct { 
	ID 				int 			`uri:"id" binding:"required"`
	User 			user.User
}