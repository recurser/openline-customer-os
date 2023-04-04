package server

import (
	contactGrpcService "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/proto/contact"
	phoneNumberGrpcService "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/proto/phone_number"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/constants"
	contactService "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contact/service"
	phoneNumberService "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/phone_number/service"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
)

const (
	maxConnectionIdle = 5
	gRPCTimeout       = 15
	maxConnectionAge  = 5
	gRPCTime          = 10
)

func (server *server) newEventProcessorGrpcServer() (func() error, *grpc.Server, error) {
	l, err := net.Listen(constants.Tcp, server.cfg.GRPC.Port)
	if err != nil {
		return nil, nil, errors.Wrap(err, "net.Listen")
	}

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: maxConnectionIdle * time.Minute,
			Timeout:           gRPCTimeout * time.Second,
			MaxConnectionAge:  maxConnectionAge * time.Minute,
			Time:              gRPCTime * time.Minute,
		}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_recovery.UnaryServerInterceptor(),
		),
		),
	)

	contactService := contactService.NewContactService(server.log, server.commands.ContactCommandsService)
	contactGrpcService.RegisterContactGrpcServiceServer(grpcServer, contactService)

	phoneNumberService := phoneNumberService.NewPhoneNumberService(server.log, server.repositories, server.commands.PhoneNumberCommands)
	phoneNumberGrpcService.RegisterPhoneNumberGrpcServiceServer(grpcServer, phoneNumberService)

	go func() {
		server.log.Infof("%s gRPC server is listening on port: {%s}", GetMicroserviceName(server.cfg), server.cfg.GRPC.Port)
		server.log.Error(grpcServer.Serve(l))
	}()

	return l.Close, grpcServer, nil
}