package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/openpsd/modelbank/providers"
	"github.com/openpsd/modelbank/usecases"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"

	"github.com/openpsd/modelbank/controllers"
	"github.com/openpsd/modelbank/entities"

	pb "github.com/openpsd/modelbank/providers"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "testdata/route_guide_db.json", "A json file containing a list of features")
	port       = flag.Int("port", 10000, "The server port")
)

type modelbankServer struct {
	Controller controllers.ModelBankController
}

func (s *modelbankServer) CreateConsent(ctx context.Context, consent *pb.Consent) (*pb.Consent, error) {
	var err error
	var resp usecases.Response
	req := usecases.NewRequest()
	cnsnt := entities.NewConsent()
	cnsnt.ID = consent.GetId()
	req.Container["consent"], err = cnsnt.Marshal()
	resp, err = s.Controller.Usecase.CreateConsent(req)
	// No feature was found, return an unnamed feature
	return &pb.Consent{}, nil
}

func newServer() *modelbankServer {
	repository := providers.NewDbProvider()
	controller := controllers.NewModelBankontroller(repository)
	s := &modelbankServer{Controller: controller}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterModelBankServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
