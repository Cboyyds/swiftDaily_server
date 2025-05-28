package request

type Register struct {
	Email            string `json:"email" binding:"required"`
	Password         string `json:"password" binding:"required"`
	Username         string `json:"user_name" binding:"required"`
	CompanyID        uint   `json:"company_id" binding:"required"`
	VerificationCode string `json:"verification_code" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
