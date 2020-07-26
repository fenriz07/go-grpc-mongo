package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/fenriz07/go-grpc-mongo/server/bd"
	"github.com/fenriz07/go-grpc-mongo/server/handler"
	"github.com/fenriz07/go-grpc-mongo/server/product"
	"google.golang.org/grpc"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("Iniciando server go")

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la bd")
		return
	}

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatal(err.Error())
	}

	s := grpc.NewServer()

	product.RegisterProductServiceServer(s, &handler.Server{})

	go func() {
		log.Println("Iniciando server grpc")
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
