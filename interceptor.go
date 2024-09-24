package interceptor

import (
	"context"

	"github.com/mennanov/fmutils"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// UnaryServerInterceptor returns a new unary server interceptor that will decide whether to which fields should return to clients.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			return
		}
		reqWithFieldMask, ok := req.(interface {
			GetFieldMask() *fieldmaskpb.FieldMask
		})
		if !ok {
			return
		}
		if len(reqWithFieldMask.GetFieldMask().GetPaths()) > 0 {
			protoResp, ok := resp.(proto.Message)
			if !ok {
				return
			}
			fmutils.Filter(protoResp, reqWithFieldMask.GetFieldMask().GetPaths())
		}
		return
	}
}
