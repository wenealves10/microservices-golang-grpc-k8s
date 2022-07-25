package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/wenealves10/microservices-golang-grpc-k8s/internal/service"
	greeting_v1 "github.com/wenealves10/microservices-golang-grpc-k8s/pkg/pb/greeting/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	listener net.Listener
	server   *grpc.Server
	logger   *zap.Logger
)

func main() {
	// Inicializa o Logger.
	logger, _ = zap.NewProduction()
	defer logger.Sync()

	// Inicializa o Listener.
	initListener()

	// Instancia o servidor gRPC.
	server = grpc.NewServer()
	reflection.Register(server)

	// Registra os handlers.
	greeting_v1.RegisterGreetingServiceServer(server, &service.GreeterService{})
	logger.Info("Handlers registered")

	// Executa um listener da aplicaçnao em uma goroutine apartada.
	go signalsListener(server)

	// Inicializa o servidor.
	logger.Info("Starting gRPC server...")
	if err := server.Serve(listener); err != nil {
		logger.Panic("Failed to start gRPC server", zap.Error(err))
	}
}

// initListener instancia um novo listener TCP.
// Em caso de erro, lança uma exceção.
func initListener() {
	var err error
	addr := "localhost:50051"

	listener, err = net.Listen("tcp", addr)
	if err != nil {
		logger.Panic("Failed to listen",
			zap.String("address", addr),
			zap.Error(err),
		)
	}

	logger.Info("Started listening...", zap.String("address", addr))
	return
}

// signalsListener cria um channel em Go para ser notificado
// caso nossa aplicação seja interrompida por algum motivo, o que
// nos permite fechar conexões e parar nosso servidor em modo "graceful".
func signalsListener(server *grpc.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	_ = <-sigs

	logger.Info("Gracefully stopping server...")
	server.GracefulStop()
}
