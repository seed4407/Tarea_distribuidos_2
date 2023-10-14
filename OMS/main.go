/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	// "net" 
    "os"
	"bufio"
	// "math/rand"
	// "time"
	"strconv"
	// "google.golang.org/grpc"
	// "github.com/streadway/amqp"
	pb "github.com/seed4407/Tarea_distribuidos_2/proto"
)

var (
	port = flag.Int("port", 80, "The server port")
)

var id_datos int
var datos_cupos int 
var err error
var datos_rechazados int 
// var valor_inicial string
var valor_modificado int
var numeroAleatorio int
var limite_inferior int
var limite_superior int
var aux string
// var ch *amqp.Channel
// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedNameNodeServer
}

func (s *server) Recepcion_Info(ctx context.Context, in *pb.Datos) (*pb.Recepcion, error) {
	var datos_persona = in.GetApellido()
	var nodo int
	filePath := "/app/Data.txt"

	if datos_persona[0] <= 77 {
		log.Printf("enviar a datanode1")
		nodo = 1
	} else{
		log.Printf("enviar a datanode2")
		nodo = 2
	}

	dato_escribir := strconv.Itoa(id_datos) + " " + strconv.Itoa(nodo) + " " +  in.GetEstado()
	id_datos = id_datos + 1

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	//Se escribe en archivo
	_, err = file.WriteString(dato_escribir + "\\n")
	if err != nil {
		return nil, err
	}
	
	err = file.Sync()
	if err != nil {
		return nil, err
	}
	file.Close()

	return &pb.Recepcion{Ok:aux}, nil
}

func main() {
    // Abrir el archivo en modo lectura
	filePath := "/app/Data.txt"

	// var file *os.File
	//Se verifica que archivo este creado
	// if _, err = os.Stat(filePath); os.IsNotExist(err) {
	// 	file, err = os.Create(filePath)
	// 	if err != nil {
	// 		return 
	// 	}
	// 	file.Close()
	// }

	//Se abre archivo
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	file.Close()

	//Leer linea por linea
	file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

	if err = scanner.Err(); err != nil {
        log.Fatal(err)
    }

	file.Close()

	var nodo int 
	nodo = 1
	id_datos = 0
	dato_escribir := strconv.Itoa(id_datos) + " " + strconv.Itoa(nodo) + " " +  "infectado"
	id_datos = id_datos + 1
	fmt.Printf(dato_escribir)

    // // Lee el contenido del archivo
    // contenido, err := os.ReadFile(filePath)
    // if err != nil {
    //     fmt.Printf("Error al leer el archivo: %v\n", err)
    //     return
    // }

	// valor_inicial = string(contenido)
	// fmt.Printf(valor_inicial)

	// if valor_inicial >= 0{
	// 	log.Printf("Inicio exitoso")
	// }

	// lis, err := net.Listen("tcp", fmt.Sprintf(":%d",*port))
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// s := grpc.NewServer()
	// pb.RegisterServidorRegionalServer(s, &server{})
	// log.Printf("server listening at %v", lis.Addr())
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
}
