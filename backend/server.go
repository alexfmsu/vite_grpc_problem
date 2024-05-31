package main

import (
	lotspb "backend/proto"

	// system
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"
	"time"

	// external
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	getopt "github.com/pborman/getopt/v2"
	"google.golang.org/grpc"
)

type RpcServer struct {
	Grpc        *grpc.Server
	WrappedGrpc *grpcweb.WrappedGrpcServer
}

type CategoryServer struct{}

func (s *CategoryServer) ActiveLots(req *lotspb.LotsRequest, resp lotspb.LotsService_ActiveLotsServer) error {
	// startPrice := rand.Intn(100)

	var i int = 0
	for {
		res := lotspb.LotsResponse{
			Lot: &lotspb.Lot{
				TotalMoneyRU: float64(i),
				FreeMoneyRU:  float64(i),
				Stocks: []*lotspb.Stock{
					&lotspb.Stock{
						Count:                     int64(rand.IntN(100)),
						Lots:                      20,
						Ticker:                    "ok1",
						Name:                      "ok2",
						CurrentPrice:              "ok3",
						BuyPrice:                  "ok4",
						Diff:                      "ok5",
						DiffPercent:               "ok6",
						DiffWithCommission:        "ok7",
						DiffPercentWithCommission: "ok8",
						Currency:                  "ok9",
						TotalCost:                 "ok10",
					},
				},
				// Price:        float64(startPrice * i),
			},
		}
		resp.Send(&res)
		time.Sleep(time.Second * 1)
		// if req.Limit < int64(i)+1 {
		// 	break
		// }

		i++
	}
	return nil
}

func NewServer() *RpcServer {
	var opts []grpc.ServerOption
	//opts = append(opts, ServerInterceptor())
	opts = append(opts)

	// It's increase to 5MB the maximum size allowed for requests and responses
	opts = append(opts, grpc.MaxSendMsgSize(5*1024*1024*1024*1024))
	opts = append(opts, grpc.MaxRecvMsgSize(5*1024*1024*1024*1024))
	gs := grpc.NewServer(opts...)
	return &RpcServer{
		Grpc:        gs,
		WrappedGrpc: grpcweb.WrapServer(gs),
	}
}

var (
	port int = 0
)

func main() {
	getopt.FlagLong(&port, "port", 'p', "port")
	getopt.Parse()

	if port == 0 {
		log.Fatal("Port is 0")
		port = 9002
		// return
	}

	fmt.Println("port: ", port)

	rpcServer := NewServer()

	categoryServer := CategoryServer{}

	lotspb.RegisterLotsServiceServer(rpcServer.Grpc, &categoryServer)

	Start(":"+strconv.Itoa(port), rpcServer)
}

func Start(httpPort string, rpcServer *RpcServer) {
	grpc := rpcServer.WrappedGrpc

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		allowCors(resp, req)
		if grpc.IsGrpcWebRequest(req) || grpc.IsAcceptableGrpcCorsRequest(req) {
			grpc.ServeHTTP(resp, req)
		}
	})

	fmt.Println("HTTP server listening on", httpPort)
	err := http.ListenAndServe(httpPort, nil)
	if err != nil {
		log.Fatal("Failed to start a HTTP server:", err)
	}
}

func allowCors(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	resp.Header().Set("Access-Control-Expose-Headers", "grpc-status, grpc-message")
	resp.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, XMLHttpRequest, x-user-agent, x-grpc-web, grpc-status, grpc-message")
}
