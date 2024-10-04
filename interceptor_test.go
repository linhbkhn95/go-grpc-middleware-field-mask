// Copyright (c) The go-grpc-middleware Authors.
// Licensed under the Apache License 2.0.

package interceptor

import (
	"context"
	"testing"

	discoveryv1 "github.com/linhbkhn95/go-grpc-middleware-field-mask/pb/go/discovery/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestFieldMaskSuite(t *testing.T) {
	s := &FieldMaskSuite{
		InterceptorTestSuite: &discoveryv1.InterceptorTestSuite{
			DiscoveryService: &discoveryv1.DiscoveryService{},
			ServerOpts: []grpc.ServerOption{
				grpc.UnaryInterceptor(
					UnaryServerInterceptor(DefaultFilterFunc),
				),
			},
		},
	}
	suite.Run(t, s)
}

type FieldMaskSuite struct {
	*discoveryv1.InterceptorTestSuite
}

func (s *FieldMaskSuite) TestUnary_ReturnAllResponseWhenDisableFieldMask() {
	resp, err := s.Client.ListProducts(context.Background(), &discoveryv1.ListProductsRequest{Id: "1"})
	assert.Equal(s.T(), nil, err)
	expected := &discoveryv1.ListProductsResponse{
		Result: &discoveryv1.ListProductsResult{
			Products: []*discoveryv1.Product{
				{
					Id:    "1",
					Name:  "Product 1",
					Img:   "Image 1",
					Price: 1,
					Shop: &discoveryv1.Shop{
						Id:   "1",
						Name: "Shop 1",
					},
				},
				{
					Id:    "2",
					Name:  "Product 2",
					Img:   "Image 2",
					Price: 1,
					Shop: &discoveryv1.Shop{
						Id:   "2",
						Name: "Shop 2",
					},
				},
				{
					Id:    "3",
					Name:  "Product 3",
					Img:   "Image 3",
					Price: 1,
					Shop: &discoveryv1.Shop{
						Id:   "3",
						Name: "Shop 3",
					},
				},
				{
					Id:    "4",
					Name:  "Product 4",
					Img:   "Image 4",
					Price: 1,
					Shop: &discoveryv1.Shop{
						Id:   "4",
						Name: "Shop 4",
					},
				},
			},
		},
	}
	assert.ElementsMatch(s.T(), expected.Result.Products, resp.Result.Products)
}

func (s *FieldMaskSuite) TestUnary_FilterResponseWhenApplyingFieldMask() {
	resp, err := s.Client.ListProducts(
		context.Background(),
		&discoveryv1.ListProductsRequest{
			Id: "1", FieldMask: &fieldmaskpb.FieldMask{
				Paths: []string{"result.products.id", "result.products.price"},
			}})
	assert.Equal(s.T(), nil, err)
	expected := &discoveryv1.ListProductsResponse{
		Result: &discoveryv1.ListProductsResult{
			Products: []*discoveryv1.Product{
				{
					Id:    "1",
					Price: 1,
				},
				{
					Id:    "2",
					Price: 1,
				},
				{
					Id:    "3",
					Price: 1,
				},
				{
					Id:    "4",
					Price: 1,
				},
			},
		},
	}
	assert.ElementsMatch(s.T(), expected.Result.Products, resp.Result.Products)
}
