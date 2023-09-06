package model

type PassengerInfo struct {
	PassengerID string  `json:"passenger_id" gorm:"primaryKey"`
	Survived    int     `json:"survived"`
	Pclass      int     `json:"pClass"`
	Name        string  `json:"name"`
	Sex         string  `json:"sex"`
	Age         float32 `json:"age"`
	SibSb       int     `json:"sib_sb"`
	Parch       int     `json:"parch"`
	Ticket      string  `json:"ticket"`
	Fare        float64 `json:"fare"`
	Cabin       string  `json:"cabin"`
	Embarked    string  `json:"embarked"`
}

type PassengerInfoDTO struct {
	PassengerID string  `json:"passengerId"`
	Survived    int     `json:"survived"`
	Pclass      int     `json:"pClass"`
	Name        string  `json:"name"`
	Sex         string  `json:"sex"`
	Age         float32 `json:"age"`
	SibSb       int     `json:"sibSb"`
	Parch       int     `json:"parch"`
	Ticket      string  `json:"ticket"`
	Fare        float64 `json:"fare"`
	Cabin       string  `json:"cabin"`
	Embarked    string  `json:"embarked"`
}
