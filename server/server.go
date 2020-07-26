package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/fenriz07/go-grpc-mongo/product"
	"google.golang.org/grpc"
)

type server struct {
	product.UnimplementedProductServiceServer
}

func main() {
	fmt.Println("Iniciando el server go")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatal(err.Error())
	}

	s := grpc.NewServer()

	product.RegisterProductServiceServer(s, &server{})

	go func() {
		fmt.Println("Iniciando server grpc")
		if err := s.Serve(lis); err != nil {
			log.Fatal("Fallo en el inicio del server " + err.Error())
		}
	}()

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt)

	//Bloqueando hasta que resivamos la señal

	<-ch
	fmt.Println("Pausando el servidor")
	s.Stop()
	fmt.Println("Cerrando el oyente")
	lis.Close()
	fmt.Println("Chao :)")
}
