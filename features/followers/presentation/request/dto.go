package request

import "github.com/dragranzer/capstone-BE-FGD/features/followers"

type Follow struct {
	FollowedID int `json:"followed_id" form:"followed_id"`
}

func ToCore(req Follow) followers.Core {

	return followers.Core{
		FollowedID: req.FollowedID,
	}
}
