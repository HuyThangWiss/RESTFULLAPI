package InformationAPI

type Books struct {
	Id   int64  `json:"Id" bson:"Id" binding:"required`
	Name string `json:"Name" bson:"Name" binding:"required `
	Year int64 `json:"Year" bson:"Year" binding:"required `
}




