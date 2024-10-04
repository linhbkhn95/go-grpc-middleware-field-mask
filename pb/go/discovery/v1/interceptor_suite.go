package discoveryv1

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// InterceptorTestSuite is a testify/Suite that starts a gRPC PingService server and a client.
type InterceptorTestSuite struct {
	suite.Suite

	DiscoveryService *DiscoveryService
	ServerOpts       []grpc.ServerOption
	ClientOpts       []grpc.DialOption

	serverAddr     string
	ServerListener net.Listener
	Server         *grpc.Server
	clientConn     *grpc.ClientConn
	Client         DiscoveryServiceClient

	restartServerWithDelayedStart chan time.Duration
	serverRunning                 chan bool

	cancels []context.CancelFunc
}

func (s *InterceptorTestSuite) SetupSuite() {
	s.restartServerWithDelayedStart = make(chan time.Duration)
	s.serverRunning = make(chan bool)

	s.serverAddr = "127.0.0.1:0"
	go func() {
		for {
			var err error
			s.ServerListener, err = net.Listen("tcp", s.serverAddr)
			s.serverAddr = s.ServerListener.Addr().String()
			require.NoError(s.T(), err, "must be able to allocate a port for serverListener")
			s.Server = grpc.NewServer(s.ServerOpts...)
			if s.DiscoveryService == nil {
				s.DiscoveryService = &DiscoveryService{}
			}
			RegisterDiscoveryServiceServer(s.Server, s.DiscoveryService)

			w := sync.WaitGroup{}
			w.Add(1)
			go func() {
				_ = s.Server.Serve(s.ServerListener)
				w.Done()
			}()
			if s.Client == nil {
				s.Client = s.NewClient(s.ClientOpts...)
			}

			s.serverRunning <- true

			d := <-s.restartServerWithDelayedStart
			s.Server.Stop()
			time.Sleep(d)
			w.Wait()
		}
	}()

	select {
	case <-s.serverRunning:
	case <-time.After(2 * time.Second):
		s.T().Fatal("server failed to start before deadline")
	}
}

func (s *InterceptorTestSuite) RestartServer(delayedStart time.Duration) <-chan bool {
	s.restartServerWithDelayedStart <- delayedStart
	time.Sleep(10 * time.Millisecond)
	return s.serverRunning
}

func (s *InterceptorTestSuite) NewClient(dialOpts ...grpc.DialOption) DiscoveryServiceClient {
	var err error
	dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.clientConn, err = grpc.NewClient(s.ServerAddr(), dialOpts...)
	require.NoError(s.T(), err, "must not error on client Dial")
	return NewDiscoveryServiceClient(s.clientConn)
}

func (s *InterceptorTestSuite) ServerAddr() string {
	return s.serverAddr
}

func (s *InterceptorTestSuite) TearDownSuite() {
	time.Sleep(10 * time.Millisecond)
	if s.ServerListener != nil {
		s.Server.GracefulStop()
		s.T().Logf("stopped grpc.Server at: %v", s.ServerAddr())
		_ = s.ServerListener.Close()
	}
	if s.clientConn != nil {
		_ = s.clientConn.Close()
	}
	for _, c := range s.cancels {
		c()
	}
}
