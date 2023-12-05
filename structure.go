package webhook

type Reply struct {
	Message string `bson:"messsage"`
}

type Logindata struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
