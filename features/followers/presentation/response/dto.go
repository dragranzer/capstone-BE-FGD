package response

import "github.com/dragranzer/capstone-BE-FGD/features/followers"

type Followed struct {
	UserID int
	Name   string
}

func FromCore(res followers.Core) Followed {
	return Followed{
		UserID: res.FollowedID,
		Name:   res.NameFollowed,
	}
}

func FromCoreSlice(core []followers.Core) []Followed {
	var FolArray []Followed
	for key := range core {
		FolArray = append(FolArray, FromCore(core[key]))
	}
	return FolArray
}
