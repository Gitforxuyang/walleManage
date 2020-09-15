package dto

type PostApi struct {
	Url string `json:"url"`
	Method string `json:"method"`
	GrpcService string `json:"grpcService"`
	GrpcMethod string `json:"grpcMethod"`
	Plugins []Plugin `json:"plugins"`
	Desc string `json:"desc"`
	Status int32 `json:"status"`
}

type Plugin struct {
	Name  string `json:"name"`
	Param string `json:"param"`
}