package users

type Core struct {
	ID             int
	Email          string
	Password       string
	Username       string
	Follower       int
	SumLike        int
	SumComment     int
	ProfilePicture string
}

type Bussiness interface {
	Register(data Core) (err error)
	Login(data Core) (userData Core, token string, isAuth bool, err error)
	GetProfileData(data Core) (resp Core, err error)
}

type Data interface {
	CreateUser(data Core) (err error)
	CheckEmailPass(email string, pass string) (isAuth bool, user Core, err error)
	SelectDatabyEmail(data Core) (resp Core, err error)
	SelectDatabyID(data Core) (resp Core, err error)
}
