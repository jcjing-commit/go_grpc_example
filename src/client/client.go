package main

import (
	"calc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
)

func add(client calc.CalcClient, num1 int64, num2 int64) int64 {
	// Contact the server and print out its response.
	// Add test with 12 13 and 25
	addRulst, err := client.Add(context.Background(), &calc.AddRequest{Num1: num1, Num2: num2})
	if err != nil {
		log.Fatalf("grpc Add fail: %v", err)
	} else {
		log.Printf("grpc Add: %v + %v = %v", num1, num2, addRulst.Result)
	}
	return addRulst.Result
}

func sub(client calc.CalcClient, num1 int64, num2 int64) int64 {
	// Sub test with 8,2 and 6
	subRulst, err := client.Sub(context.Background(), &calc.SubRequest{Num1: num1, Num2: num2})
	if err != nil {
		log.Fatalf("grpc Sub fail: %v", err)
	} else {
		log.Printf("grpc Sub: %v - %v = %v", num1, num2, subRulst.Result)
	}
	return subRulst.Result
}

func mult(client calc.CalcClient, num1 int64, num2 int64) int64 {
	//Mult test with 5,2 and 10
	multRulst, err := client.Mult(context.Background(), &calc.MultRequest{Num1: num1, Num2: num2})
	if err != nil {
		log.Fatalf("grpc Mult fail: %v", err)
	} else {
		log.Printf("grpc Mult: %v * %v = %v", num1, num2, multRulst.Result)
	}
	return multRulst.Result
}

func div(client calc.CalcClient, num1 int64, num2 int64) int64 {
	//Div test with 24 2 and 12
	divRulst, err := client.Div(context.Background(), &calc.DivRequest{Num1: num1, Num2: num2})
	if err != nil {
		log.Fatalf("grpc Div fail: %v", err)
	} else {
		log.Printf("grpc div: %v / %v = %v", num1, num2, divRulst.Result)
	}
	return divRulst.Result
}

func login(client calc.CalcClient, name string, pwd string) {
	var users []*calc.LoginRequest
	users = append(users, &calc.LoginRequest{Name: name, Pwd: pwd})
	stream, err := client.Login(context.Background())
	if err != nil {
		log.Fatalf("%v.Login() is err: %v", client, err)
	}
	for _, user := range users {
		if err := stream.Send(user); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, user, err)
		} else {
			log.Fatalf("login oK")
		}
	}
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := calc.NewCalcClient(conn)

	// Contact the server and print out its response.

	add(c, 12, 13)

	sub(c, 8, 2)

	mult(c, 2, 5)

	div(c, 24, 2)

	login(c, "grpc", "123")
}
