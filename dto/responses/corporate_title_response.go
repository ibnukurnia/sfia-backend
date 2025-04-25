package responses

type CorporateTitleResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CorporateTitleListResponse struct {
	CorporateTitleID   string `json:"corporate_title_id" gorm:"column:uuid"`
	CorporateTitleName string `json:"corporate_title_name" gorm:"column:name"`
}
