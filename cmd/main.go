package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/eggsbenjamin/groceries-tesco/domain"
	pb "github.com/eggsbenjamin/groceries-tesco/grpc"
	"github.com/eggsbenjamin/groceries-tesco/server"
	"github.com/eggsbenjamin/groceries-tesco/service/tesco"
	"google.golang.org/grpc"

	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("tesco")
}

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", viper.GetString("service_port")))
	if err != nil {
		log.Fatal("error listening on port - %v", err)
	}

	c := &http.Client{}
	tS := tesco.NewTescoService(c)
	pG := domain.NewProductHandler(tS)
	pGH := server.NewGetProductsHandler(pG)
	s := grpc.NewServer()
	pb.RegisterTescoServiceServer(s, pGH)
	s.Serve(l)
}
