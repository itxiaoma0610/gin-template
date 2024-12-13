package request

type LoginRequest struct {
	Mobile string `json:"mobile" binding:"required"`
	Code   string `json:"code" binding:"required"`
}

type VerificationMobile struct {
	Mobile string `json:"mobile" binding:"required"`
}
