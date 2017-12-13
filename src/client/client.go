package main

import (
    "log"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "calc"
)

const (
    address     = "localhost:50051"
)

func main() {
    // Set up a connection to the server.
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := calc.NewCalcClient(conn)

    // Contact the server and print out its response.
    // Add test with 12 13 and 25 
    var num1, num2 int64
    num1 = 12
    num2 = 13
    addRulst,err :=  c.Add(context.Background(),&calc.AddRequest{Num1:num1,Num2:num2})
    if err != nil{
        log.Fatalf("grpc Add fail: %v",err)
    }else{
        log.Printf("grpc Add: %v + %v = %v",num1,num2,addRulst.Result)
    }
    // Sub test with 8,2 and 6
    num1 = 8
    num2 = 2
    subRulst,err := c.Sub(context.Background(), &calc.SubRequest{Num1:num1,Num2:num2})
    if err != nil{
        log.Fatalf("grpc Sub fail: %v",err)
    }else{
        log.Printf("grpc Sub: %v - %v = %v",num1,num2,subRulst.Result)
    }
   //Mult test with 5,2 and 10 
    num1 = 5 
    num2 = 2
    multRulst,err := c.Mult(context.Background(), &calc.MultRequest{Num1:num1,Num2:num2})
    if err != nil{
	log.Fatalf("grpc Mult fail: %v",err)
    }else{
	log.Printf("grpc Mult: %v * %v = %v",num1,num2,multRulst.Result)
    }
    //Div test with 24 2 and 12
    num1 = 24
    num2 = 2
    divRulst,err := c.Div(context.Background(), &calc.DivRequest{Num1:num1,Num2:num2})
    if err != nil{
        log.Fatalf("grpc Div fail: %v",err)
    }else{
        log.Printf("grpc div: %v / %v = %v",num1,num2,divRulst.Result)
    }
}
