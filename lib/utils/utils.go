package utils

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	err "com.youyu.api/lib/errors"
)

// 从go文件中定义的errors转换为grpc中自定义的errors
func CustomErrToGrpcCustomErr(grpcErr *rpc.Errors, errno *err.Errno) *rpc.Errors {
	grpcErr.Code = int32(errno.Code)
	grpcErr.Message = errno.Message
	grpcErr.HttpCode = int32(errno.HttpCode)
	return grpcErr
}

func GrpcCustomErrToCustomErr(grpcErr *rpc.Errors, errno *err.Errno) *err.Errno {
	errno.Code = int(grpcErr.Code)
	errno.HttpCode = int(grpcErr.HttpCode)
	errno.Message = grpcErr.Message
	return errno
}
