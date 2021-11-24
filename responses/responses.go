package response

import (
	"net/http"
)

// function response false param
func FalseParamResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "False Param",
	}
	return result
}

// function response bad request
func BadRequestResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Bad Request",
	}
	return result
}

// function response access forbidden
func AccessForbiddenResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Access Forbidden",
	}
	return result
}

// function response success dengan paramater
func SuccessResponseData(data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation",
		"data":    data,
	}
	return result
}

// function response success tanpa parameter
func SuccessResponseNonData() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation",
	}
	return result
}

// function response login failure
func LoginFailedResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Login Failed",
	}
	return result
}

// function response login success
func LoginSuccessResponse(data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Login Success",
		"data":    data,
	}
	return result
}
