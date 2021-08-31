package data

type CenterLogin struct {
	CenterUserName string `form:"centeruser" bson:"centeruser" json:"centeruser"`
	CenterPassword string `form:"centerpass" bson:"centerpass" json:"centerpass"`
}
