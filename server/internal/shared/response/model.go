package response

type DocBaseResponse struct {
	Message string `json:"message" example:"Operation successful"`
}

// Response is the standard success response wrapper
type Response[T any] struct {
	Message string `json:"message"`
	Data    *T     `json:"data,omitempty"`
}

type PaginationMeta struct {
	NextCursor *string `json:"next_cursor,omitempty"`
	HasNext    bool    `json:"has_next"`
	Count      int     `json:"count"`
}

// PaginatedResponse is the standard paginated response wrapper
type PaginatedResponse[T any] struct {
	Message string         `json:"message"`
	Data    []T            `json:"data"`
	Meta    PaginationMeta `json:"meta"`
}

// ErrorResponse is the standard error response wrapper
type ErrorResponse struct {
	Message string `json:"message"`
}

func NewResponse[T any](message string, data T) Response[T] {
	if message == "" {
		message = "Successful operation"
	}

	return Response[T]{
		Message: message,
		Data:    &data,
	}
}

func NewMessageResponse(message string) Response[any] {
	if message == "" {
		message = "Successful operation"
	}

	return Response[any]{
		Message: message,
	}
}

func NewPaginatedResponse[T any](message string, data []T, meta PaginationMeta) PaginatedResponse[T] {
	if message == "" {
		message = "Successful operation"
	}

	return PaginatedResponse[T]{
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Message: message,
	}
}
