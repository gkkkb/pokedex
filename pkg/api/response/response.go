// Package response is used to write response based on Bukalapak API V4 standard
package response

import (
	"encoding/json"
	"net/http"
	"strings"
)

var (
	ErrTetapTenangTetapSemangat = CustomError{
		Message:  "Tetap Tenang Tetap Semangat",
		Code:     999,
		HTTPCode: http.StatusInternalServerError,
	}

	// UnexpectedServerError represents Internal server error
	UnexpectedServerError = CustomError{
		Message:  "Unexpected server error",
		Code:     10000,
		HTTPCode: http.StatusInternalServerError,
	}

	// RecordConflictError represents Duplicate entry for unique field error
	RecordConflictError = CustomError{
		Message:  "Record conflict",
		Code:     10202,
		HTTPCode: http.StatusConflict,
	}
	// ProposalStatusConflictError represents new status for proposal not acceptable
	ProposalNewStatusError = CustomError{
		Message:  "Status not acceptable",
		Code:     10205,
		HTTPCode: http.StatusNotAcceptable,
	}
	// DataNotUpdatableError represents Data not updatable error
	DataNotUpdatableError = CustomError{
		Message:  "Data not updatable",
		Code:     10210,
		HTTPCode: http.StatusNotAcceptable,
	}
	// ProposalNotExistsError represents Proposal not found error
	ProposalNotExistsError = CustomError{
		Message:  "Proposal not exists or has been deleted",
		Code:     10211,
		HTTPCode: http.StatusNotFound,
	}
	// BrandNotExistsError represents Brand not found error
	BrandNotExistsError = CustomError{
		Message:  "Brand not exists or has been deleted",
		Code:     10212,
		HTTPCode: http.StatusNotFound,
	}
	// ModelNotExistsError represents Model not found error
	ModelNotExistsError = CustomError{
		Message:  "Model not exists or has been deleted",
		Code:     10213,
		HTTPCode: http.StatusNotFound,
	}
	// VariantNotExistsError represents Variant not found error
	VariantNotExistsError = CustomError{
		Message:  "Variant not exists or has been deleted",
		Code:     10214,
		HTTPCode: http.StatusNotFound,
	}
	// VehicleImageNotExistsError represents Vehicle Image not found error
	VehicleImageNotExistsError = CustomError{
		Message:  "Vehicle Image not exists or has been deleted",
		Code:     10215,
		HTTPCode: http.StatusNotFound,
	}
	// VehicleColorNotExistsError represents Vehicle Color not found error
	VehicleColorNotExistsError = CustomError{
		Message:  "Vehicle Color not exists or has been deleted",
		Code:     10216,
		HTTPCode: http.StatusNotFound,
	}
	// BannerNotExistsError represents Banner not found error
	BannerNotExistsError = CustomError{
		Message:  "Banner not exists or has been deleted",
		Code:     10217,
		HTTPCode: http.StatusNotFound,
	}
	// LocationNotExistsError represents Location not found error
	LocationNotExistsError = CustomError{
		Message:  "Location does not exists or has been deleted",
		Code:     10218,
		HTTPCode: http.StatusNotFound,
	}
	// VariantOnLocationNotExistsError represents Variant not found on a location error
	VariantOnLocationNotExistsError = CustomError{
		Message:  "Variant does not exists or has been deleted in requested location",
		Code:     10219,
		HTTPCode: http.StatusNotFound,
	}
	// CityNotExistsError represents City not found error
	CityNotExistsError = CustomError{
		Message:  "City does not exists or has been deleted",
		Code:     10220,
		HTTPCode: http.StatusNotFound,
	}

	// InvalidTokenError represents Invalid token error
	InvalidTokenError = CustomError{
		Message:  "Invalid token",
		Code:     10102,
		HTTPCode: http.StatusUnauthorized,
	}
	// InvalidParameterError represents Invalid parameter error
	InvalidParameterError = CustomError{
		Message:  "Invalid parameter",
		Code:     10111,
		HTTPCode: http.StatusUnprocessableEntity,
	}

	// UserUnauthorizedError represents User unauthorized error
	UserUnauthorizedError = CustomError{
		Message:  "User unauthorized",
		Code:     10003,
		HTTPCode: http.StatusForbidden,
	}
	// ProposalExistsError represents User already have active proposal exists
	ProposalExistsError = CustomError{
		Message:  "Proposal exists",
		Code:     10004,
		HTTPCode: http.StatusForbidden,
	}

	// BadRequestError represents bad request error
	BadRequestError = CustomError{
		Message:  "Bad Request",
		Code:     10005,
		HTTPCode: http.StatusBadRequest,
	}

	// InvalidFileTypeError represents Uploaded file type not supported
	InvalidFileTypeError = CustomError{
		Message:  "File type not supported",
		Code:     71001,
		HTTPCode: http.StatusUnsupportedMediaType,
	}
	// NoMediaError represents No attached media / no media type error
	NoMediaError = CustomError{
		Message:  "No attached media / no media type",
		Code:     71002,
		HTTPCode: http.StatusNotAcceptable,
	}

	//OfflineProposalCsvError represents error on offline proposals csv
	OfflineProposalCsvError = CustomError{
		Message:  "Invalid Offline Proposal CSV File",
		Code:     71003,
		HTTPCode: http.StatusBadRequest,
	}
)

type ResponseBody struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Errors  []ErrorInfo `json:"errors,omitempty"`
	Meta    MetaInfo    `json:"meta"`
}

// MetaInfo holds meta data
type MetaInfo struct {
	HTTPStatus int         `json:"http_status"`
	Offset     int         `json:"offset,omitempty"`
	Limit      int         `json:"limit,omitempty"`
	Total      int64       `json:"total,omitempty"`
	Sort       string      `json:"sort,omitempty"`
	Facets     interface{} `json:"facets,omitempty"`
}

// ErrorBody holds data for error response
type ErrorBody struct {
	Errors []ErrorInfo `json:"errors"`
	Meta   interface{} `json:"meta"`
}

// ErrorInfo holds error detail
type ErrorInfo struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Field   string `json:"field,omitempty"`
}

// CustomError holds data for customized error
type CustomError struct {
	Message  string
	Field    string
	Code     int
	HTTPCode int
}

// Error is a function to convert error to string.
// It exists to satisfy error interface
func (c CustomError) Error() string {
	return c.Message
}

// BuildSuccess is a function to create ResponseBody
func BuildSuccess(data interface{}, meta MetaInfo) ResponseBody {
	return ResponseBody{
		Data: data,
		Meta: meta,
	}
}

// BuildError is a function to create ErrorBody
func BuildError(errors []error) ErrorBody {
	var (
		ce CustomError
		ok bool
	)

	if len(errors) == 0 {
		ce = ErrTetapTenangTetapSemangat
	} else {
		err := errors[0]
		ce, ok = err.(CustomError)
		if !ok {
			ce = ErrTetapTenangTetapSemangat
		}
	}

	return ErrorBody{
		Errors: []ErrorInfo{
			{
				Message: ce.Message,
				Code:    ce.Code,
				Field:   ce.Field,
			},
		},
		Meta: MetaInfo{
			HTTPStatus: ce.HTTPCode,
		},
	}
}

// BuildErrors is a function to create ErrorBody
func BuildErrors(errors []error) ErrorBody {
	var (
		ce         CustomError
		ok         bool
		errorInfos []ErrorInfo
	)

	for _, err := range errors {
		ce, ok = err.(CustomError)
		if !ok {
			ce = ErrTetapTenangTetapSemangat
		}

		errorInfo := ErrorInfo{
			Code:    ce.Code,
			Field:   ce.Field,
			Message: ce.Message,
		}

		errorInfos = append(errorInfos, errorInfo)
	}

	return ErrorBody{
		Errors: errorInfos,
		Meta: MetaInfo{
			HTTPStatus: ce.HTTPCode,
		},
	}
}

// BuildErrorAndStatus is a function to Differentiate Error and create Error Body and Response Status Code
func BuildErrorAndStatus(err error, fieldName string) (ErrorBody, int) {

	if strings.Contains(err.Error(), "strconv.ParseInt: parsing") ||
		strings.Contains(err.Error(), "strconv.Atoi: parsing") ||
		strings.Contains(err.Error(), "invalid character") ||
		strings.Contains(err.Error(), "json: cannot unmarshal") {
		return BuildError([]error{InvalidParameterError}), InvalidParameterError.HTTPCode
	} else if strings.Contains(err.Error(), "is not a valid Vehicle Type") {
		pe := InvalidParameterError
		pe.Field = fieldName
		return BuildError([]error{pe}), InvalidParameterError.HTTPCode
	} else if strings.Contains(err.Error(), "Proposal not found") {
		return BuildError([]error{ProposalNotExistsError}), ProposalNotExistsError.HTTPCode
	} else if strings.Contains(err.Error(), "Proposal exists") {
		return BuildError([]error{ProposalExistsError}), ProposalExistsError.HTTPCode
	} else if strings.Contains(err.Error(), "File type not supported") {
		return BuildError([]error{InvalidFileTypeError}), InvalidFileTypeError.HTTPCode
	} else if strings.Contains(err.Error(), "sql: Scan error") {
		return ErrorBody{
				Errors: []ErrorInfo{
					{
						Message: http.StatusText(http.StatusInternalServerError),
						Code:    http.StatusInternalServerError,
					},
				},
				Meta: MetaInfo{
					HTTPStatus: http.StatusInternalServerError,
				}},
			http.StatusInternalServerError
	} else if strings.Contains(err.Error(), "Duplicate entry") {
		ce := RecordConflictError
		ce.Field = fieldName

		return BuildError([]error{ce}), InvalidFileTypeError.HTTPCode
	} else if strings.Contains(err.Error(), "mime: no media type") {
		return BuildError([]error{NoMediaError}), NoMediaError.HTTPCode
	} else if strings.Contains(err.Error(), "Brand Not Found") {
		return BuildError([]error{BrandNotExistsError}), BrandNotExistsError.HTTPCode
	} else if strings.Contains(err.Error(), "Model Not Found") {
		return BuildError([]error{ModelNotExistsError}), ModelNotExistsError.HTTPCode
	} else if strings.Contains(err.Error(), "Variant Not Found") {
		return BuildError([]error{VariantNotExistsError}), VariantNotExistsError.HTTPCode
	} else if strings.Contains(err.Error(), "New status is not acceptable") {
		return BuildError([]error{ProposalNewStatusError}), ProposalNewStatusError.HTTPCode
	} else if strings.Contains(err.Error(), "Unknown column") {
		return BuildError([]error{UnexpectedServerError}), UnexpectedServerError.HTTPCode
	} else if strings.Contains(err.Error(), "User not authorized") {
		return BuildError([]error{UserUnauthorizedError}), UserUnauthorizedError.HTTPCode
	} else if strings.Contains(err.Error(), "Data not updatable") {
		return BuildError([]error{DataNotUpdatableError}), DataNotUpdatableError.HTTPCode
	} else if strings.Contains(err.Error(), LocationNotExistsError.Message) {
		return BuildError([]error{LocationNotExistsError}), LocationNotExistsError.HTTPCode
	} else if strings.Contains(err.Error(), VariantOnLocationNotExistsError.Message) {
		return BuildError([]error{VariantOnLocationNotExistsError}), VariantOnLocationNotExistsError.HTTPCode
	} else if strings.Contains(err.Error(), CityNotExistsError.Message) {
		return BuildError([]error{CityNotExistsError}), CityNotExistsError.HTTPCode
	}

	return BuildError([]error{ErrTetapTenangTetapSemangat}), ErrTetapTenangTetapSemangat.HTTPCode
}

// Write is a function to write data in json format
func Write(w http.ResponseWriter, result interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(result)
}
