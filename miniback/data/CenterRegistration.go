package data

type RegistrationData struct {
	Centername        string `form:"centername" bson:"centername" json:"centername"`
	Ownername         string `form:"ownername" bson:"ownername" json:"ownername"`
	Ownermobilenumber int    `form:"ownermobilenumber" bson:"ownermobilenumber" json:"ownermobilenumber" validate:"numeric,len=10"`
	Owneremailid      string `form:"owneremailidlf" bson:"owneremailidlf" json:"owneremailidlf"`
	Lf                string `form:"lf" bson:"lf" json:"lf"`
	Lfemailid         string `form:"lfemailid" bson:"lfemailid" json:"lfemailid" validate:"email"`
	Lfmobilenumber    int    `form:"lfmobilenumber" bson:"lfmobilenumber" json:"lfmobilenumber" validate:"required,numeric,len=10"`
	City              string `form:"city" bson:"city" json:"city"`
	Pincode           int    `form:"pincode" bson:"pincode" json:"pincode"`
	Address           string `form:"address" bson:"address" json:"address"`
	Password          string `form:"password" bson:"password" json:"password" validate:"alphanum,required,gte=8,lte=20"`
	Confirmpassword   string `form:"confirmpassword" bson:"confirmpassword" json:"confirmpassword" validate:"alphanum,eqfield=Password"`
}
