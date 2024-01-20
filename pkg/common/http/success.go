package common

func NewSuccessResponse() DefaultResponse {
	return DefaultResponse{
		Code:    200,
		Message: "Success",
	}
}

func NewCreatedSuccessResponse(insertedId int64) CreatedSuccessResponse {
	return CreatedSuccessResponse{
		Code:    200,
		Message: "Success",
		Payload: insertedId,
	}
}
