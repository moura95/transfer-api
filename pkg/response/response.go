package httpresponse

type Response struct {
	Error    interface{} `json:"error"`
	Data     interface{} `json:"data"`
	PageInfo PageInfo    `json:"pageInfo"`
}

func SuccessResponseWithPageInfo(data interface{}, pageInfo PageInfo) Response {
	return Response{
		Data:     data,
		Error:    "",
		PageInfo: pageInfo,
	}
}

func SuccessResponse(data interface{}) Response {
	return Response{
		Data:  data,
		Error: "",
	}
}

func ErrorResponse(error string) Response {
	return Response{
		Data:  "",
		Error: error,
	}
}

type PageInfo struct {
	TotalRecords int
	CurrentPage  int
	TotalPages   int
	Limit        int
}

func NewPageInfo(limit, totalRecords, currentPage, totalPages int) PageInfo {
	return PageInfo{
		TotalRecords: totalRecords,
		CurrentPage:  currentPage,
		TotalPages:   totalPages,
		Limit:        limit,
	}
}
