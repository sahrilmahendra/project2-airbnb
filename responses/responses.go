package response

import (
	"net/http"
)

// function response false param
func FalseParamResponse() map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusBadRequest,
		"Message": "False Param",
	}
	return result
}

// function response bad request
func BadRequestResponse() map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusBadRequest,
		"Message": "Bad Request",
	}
	return result
}

// function response access forbidden
func AccessForbiddenResponse() map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusBadRequest,
		"Message": "Access Forbidden",
	}
	return result
}

// function response success dengan paramater
func SuccessResponseData(data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusOK,
		"Message": "Successful Operation",
		"Data":    data,
	}
	return result
}

// function response success tanpa parameter
func SuccessResponseNonData() map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusOK,
		"Message": "Successful Operation",
	}
	return result
}

// function response login failure
func LoginFailedResponse() map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusBadRequest,
		"Message": "Login Failed",
	}
	return result
}

// function response login success
func LoginSuccessResponse(data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusOK,
		"Message": "Login Success",
		"Data":    data,
	}
	return result
}

// function response available homestay
func AvailableResponse(data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusOK,
		"Message": data,
	}
	return result
}
