package entity

//User struct is to handle user data
type Score struct {
	User  string `json:"user" bson:"user" binding:"required"`
	Mode  string `json:"mode" bson:"mode" binding:"required"`
	Score float64 `json:"score" bson:"score" binding:"required"`
}
