package deuterium 

import (
	"encoding/json"
	"net/http"
)

type response struct {
	w http.ResponseWriter
}

type httpStatus int

func (res *response) Header(key, v string) {
	res.w.Header().Set(key, v)
}

func (res *response) RemoveHeader(key string) {
	res.w.Header().Del(key)
}

func (res *response) Cookie(c *http.Cookie) {
	http.SetCookie(res.w, c)
}

func (res *response) WriteString(s string) {
	res.w.WriteHeader(200)
	res.w.Write([]byte(s))
}

func (res *response) WriteJSON(v any) {
	res.w.WriteHeader(200)
	res.w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res.w).Encode(v)
}

func (res *response) WriteStatusMessage(status httpStatus, msg string) {
	res.w.WriteHeader(int(status))
	res.w.Header().Set("Content-Type", "application/json")
	data := struct {
		Status  httpStatus `json:"status"`
		Message string     `json:"message"`
	}{
		Status:  status,
		Message: msg,
	}
	json.NewEncoder(res.w).Encode(&data)
	return
}

func (res *response) BadRequestError(msg string) {
	http.Error(res.w, msg, http.StatusBadRequest)
	return
}

func (res *response) UnauthorizedError(msg string) {
	http.Error(res.w, msg, http.StatusUnauthorized)
	return
}

func (res *response) PaymentRequiredError(msg string) {
	http.Error(res.w, msg, http.StatusPaymentRequired)
	return
}

func (res *response) ForbiddenError(msg string) {
	http.Error(res.w, msg, http.StatusForbidden)
	return
}

func (res *response) NotFoundError(msg string) {
	http.Error(res.w, msg, http.StatusNotFound)
	return
}

func (res *response) MethodNotAllowedError(msg string) {
	http.Error(res.w, msg, http.StatusMethodNotAllowed)
	return
}

func (res *response) NotAcceptableError(msg string) {
	http.Error(res.w, msg, http.StatusNotAcceptable)
	return
}

func (res *response) ProxyAuthenticationRequiredError(msg string) {
	http.Error(res.w, msg, http.StatusProxyAuthRequired)
	return
}

func (res *response) TimeoutError(msg string) {
	http.Error(res.w, msg, http.StatusRequestTimeout)
	return
}

func (res *response) ConflictError(msg string) {
	http.Error(res.w, msg, http.StatusConflict)
	return
}

func (res *response) GoneError(msg string) {
	http.Error(res.w, msg, http.StatusGone)
	return
}

func (res *response) LengthRequiredError(msg string) {
	http.Error(res.w, msg, http.StatusLengthRequired)
	return
}

func (res *response) PreconditionFailedError(msg string) {
	http.Error(res.w, msg, http.StatusPreconditionFailed)
	return
}

func (res *response) EntityTooLargeError(msg string) {
	http.Error(res.w, msg, http.StatusRequestEntityTooLarge)
	return
}

func (res *response) URITooLongError(msg string) {
	http.Error(res.w, msg, http.StatusRequestURITooLong)
	return
}

func (res *response) UnsupportedMediaTypeError(msg string) {
	http.Error(res.w, msg, http.StatusUnsupportedMediaType)
	return
}

func (res *response) RangeNotSatisfiableError(msg string) {
	http.Error(res.w, msg, http.StatusRequestedRangeNotSatisfiable)
	return
}

func (res *response) ExpectationFailedError(msg string) {
	http.Error(res.w, msg, http.StatusExpectationFailed)
	return
}

func (res *response) TeapotError(msg string) {
	http.Error(res.w, msg, http.StatusTeapot)
	return
}

func (res *response) MisdirectedRequestError(msg string) {
	http.Error(res.w, msg, http.StatusMisdirectedRequest)
	return
}

func (res *response) UnprocessableEntityError(msg string) {
	http.Error(res.w, msg, http.StatusUnprocessableEntity)
	return
}

func (res *response) LockedError(msg string) {
	http.Error(res.w, msg, http.StatusLocked)
	return
}

func (res *response) FailedDependencyError(msg string) {
	http.Error(res.w, msg, http.StatusFailedDependency)
	return
}

func (res *response) TooEarlyError(msg string) {
	http.Error(res.w, msg, http.StatusTooEarly)
	return
}

func (res *response) UpgradeRequiredError(msg string) {
	http.Error(res.w, msg, http.StatusUpgradeRequired)
	return
}

func (res *response) PreconditionRequiredError(msg string) {
	http.Error(res.w, msg, http.StatusPreconditionRequired)
	return
}

func (res *response) TooManyRequestsError(msg string) {
	http.Error(res.w, msg, http.StatusTooManyRequests)
	return
}

func (res *response) HeaderFieldsTooLargeError(msg string) {
	http.Error(res.w, msg, http.StatusRequestHeaderFieldsTooLarge)
	return
}

func (res *response) UnavailableForLegalReasonsError(msg string) {
	http.Error(res.w, msg, http.StatusUnavailableForLegalReasons)
	return
}

func (res *response) InternalServerError(msg string) {
	http.Error(res.w, msg, http.StatusInternalServerError)
	return
}

func (res *response) NotImplementedError(msg string) {
	http.Error(res.w, msg, http.StatusNotImplemented)
	return
}

func (res *response) BadGatewayError(msg string) {
	http.Error(res.w, msg, http.StatusBadGateway)
	return
}

func (res *response) ServiceUnavailableError(msg string) {
	http.Error(res.w, msg, http.StatusServiceUnavailable)
	return
}

func (res *response) GatewayTimeoutError(msg string) {
	http.Error(res.w, msg, http.StatusGatewayTimeout)
	return
}

func (res *response) HTTPVersionNotSupportedError(msg string) {
	http.Error(res.w, msg, http.StatusHTTPVersionNotSupported)
	return
}

func (res *response) VariantAlsoNegotiatesError(msg string) {
	http.Error(res.w, msg, http.StatusVariantAlsoNegotiates)
	return
}

func (res *response) InsufficientStorageError(msg string) {
	http.Error(res.w, msg, http.StatusInsufficientStorage)
	return
}

func (res *response) LoopDetectedError(msg string) {
	http.Error(res.w, msg, http.StatusLoopDetected)
	return
}

func (res *response) NotExtendedError(msg string) {
	http.Error(res.w, msg, http.StatusNotExtended)
	return
}

func (res *response) NetworkAuthenticationRequiredError(msg string) {
	http.Error(res.w, msg, http.StatusNetworkAuthenticationRequired)
	return
}
