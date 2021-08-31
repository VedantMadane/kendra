package data

type AdminLogin struct {
	AdminUserName string `form:"amdinuser" bson:"adminuser" json:"adminuser"`
	AdminPassword string `form:"adminpass" bson:"adminpass" json:"adminpass"`
}
