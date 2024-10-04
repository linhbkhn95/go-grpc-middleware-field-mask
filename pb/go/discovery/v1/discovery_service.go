package discoveryv1

import "context"

// Interface implementation assert.
var _ DiscoveryServiceServer = &DiscoveryService{}

type DiscoveryService struct {
	UnimplementedDiscoveryServiceServer
}

func (s *DiscoveryService) ListProducts(ctx context.Context, req *ListProductsRequest) (*ListProductsResponse, error) {
	products := []*Product{
		{
			Id:    "1",
			Name:  "Product 1",
			Img:   "Image 1",
			Price: 1,
			Shop: &Shop{
				Id:   "1",
				Name: "Shop 1",
			},
		},
		{
			Id:    "2",
			Name:  "Product 2",
			Img:   "Image 2",
			Price: 1,
			Shop: &Shop{
				Id:   "2",
				Name: "Shop 2",
			},
		},
		{
			Id:    "3",
			Name:  "Product 3",
			Img:   "Image 3",
			Price: 1,
			Shop: &Shop{
				Id:   "3",
				Name: "Shop 3",
			},
		},
		{
			Id:    "4",
			Name:  "Product 4",
			Img:   "Image 4",
			Price: 1,
			Shop: &Shop{
				Id:   "4",
				Name: "Shop 4",
			},
		},
	}
	return &ListProductsResponse{
		Result: &ListProductsResult{
			Products: products,
		},
	}, nil
}
