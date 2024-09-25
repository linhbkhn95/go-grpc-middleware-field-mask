# go-grpc-middleware-field-mask
Go grpc middleware for field mask

# How to use
1. install via go get 

 `go get github.com/linhbkhn95/go-grpc-middleware-field-mask`
`

2. Import and inject into grpc interceptor
The code in your application should be like that:
``` Go
import(
        // ...
        "google.golang.org/grpc"
    	fieldmaskpkg "github.com/linhbkhn95/go-grpc-middleware-field-mask"
        "github.com/mennanov/fmutils"


)
// ...

func main(){
    var unaryOpts []grpc.UnaryServerInterceptor{
		fieldmaskpkg.UnaryServerInterceptor(fieldmaskpkg.DefaultFilterFunc),
    }
    // Should append others interceptors
}
```