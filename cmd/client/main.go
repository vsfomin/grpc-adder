package main

import (
	"context"
	"flag"
	"log"
	"strconv"

	"github.com/vsfomin/grpc-adder/api/proto"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("not enought arguments")
	}

	x, err := strconv.Atoi(flag.Arg((0)))
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(flag.Arg((1)))
	if err != nil {
		log.Fatal(err)
	}
	str := flag.Arg(2)

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := proto.NewAdderClient(conn)

	res, err := c.Add(context.Background(), &proto.AddRequest{X: int32(x), Y: int32(y)})
	if err != nil {
		log.Fatal(err)
	}
	mes, err := c.SayHello(context.Background(), &proto.Message{Body: str})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(mes.GetBody())
	log.Println(res.GetResult())

	//Book creating
	book := &proto.Book{
		Bid:         1,
		Title:       "War and Peace",
		Author:      "Lev Tolstoy",
		Description: "Great book for everyone",
	}

	crtBook, err := c.CreateBook(context.Background(), &proto.CreateBookRequest{Book: book})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("The book with ID: ", crtBook.GetBid(), " was added to the system!")
}
