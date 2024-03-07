package exception

var NumeratorLogMessage = struct {
	INVALID_BODY_ERROR    string
	BAD_REQUEST_ERROR     string
	INVALID_SDK_KEY_ERROR string
}{
	INVALID_BODY_ERROR:    "Fetching JSON was successful but the HTTP response content was invalid.",
	BAD_REQUEST_ERROR:     "Unexpected error occurred while trying to fetch JSON. It is most likely due to a local network issue. Please make sure your request is valid and try again.",
	INVALID_SDK_KEY_ERROR: "Your SDK Key seems to be wrong. You can find the valid SDK Key at https://web-dashboard.numerator.io/api-key",
}

func GetUnexpectedHttpResponse(responseMessage string) string {
	return "Unexpected HTTP response was received while trying to fetch JSON: " + responseMessage
}

func GetObjectDoesNotExist(responseMessage string) string {
	return "Cannot find the object of your API_KEY. Please make sure your API_KEY is correct. " + responseMessage
}
