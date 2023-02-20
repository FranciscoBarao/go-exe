package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"

	pb "bookshop/server/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedInventoryServer
}

func (s *server) GetBookList(ctx context.Context, in *pb.GetBookListRequest) (*pb.GetBookListResponse, error) {
	return &pb.GetBookListResponse{
		Books: getSampleBooks(),
	}, nil
}

// Receives a book and returns all the fields via a Stream
func (s *server) GetBookFieldsStream(book *pb.Book, stream pb.Inventory_GetBookFieldsStreamServer) error {

	if err := stream.Send(&pb.BookField{Key: "author", Value: book.GetAuthor()}); err != nil {
		return err
	}
	if err := stream.Send(&pb.BookField{Key: "language", Value: book.GetLanguage()}); err != nil {
		return err
	}
	if err := stream.Send(&pb.BookField{Key: "page_count", Value: strconv.Itoa(int(book.GetPageCount()))}); err != nil {
		return err
	}
	if err := stream.Send(&pb.BookField{Key: "title", Value: book.GetTitle()}); err != nil {
		return err
	}
	return nil
}

// Receives fields via stream and returns a book with those fields
func (s *server) CreateBookStream(stream pb.Inventory_CreateBookStreamServer) error {

	var bookMap = make(map[string]string)

	for {
		field, err := stream.Recv()
		if err == io.EOF {
			pageCount, _ := strconv.ParseInt(bookMap["page_count"], 10, 64)
			return stream.SendAndClose(&pb.Book{
				Title:     bookMap["title"],
				Author:    bookMap["author"],
				PageCount: int32(pageCount),
				Language:  StringPtr(bookMap["language"]),
			})
		}
		bookMap[field.Key] = field.Value
	}
}

// Receives books via stream and returns all fields also via stream
func (s *server) BiGetBooksToFields(stream pb.Inventory_BiGetBooksToFieldsServer) error {
	for {
		book, err := stream.Recv()
		if err == io.EOF {
			continue
		}
		if err != nil {
			return err
		}

		log.Println("New Message from client", book)

		if err := stream.Send(&pb.BookField{Key: "author", Value: book.GetAuthor()}); err != nil {
			return err
		}
		if err := stream.Send(&pb.BookField{Key: "language", Value: book.GetLanguage()}); err != nil {
			return err
		}
		if err := stream.Send(&pb.BookField{Key: "page_count", Value: strconv.Itoa(int(book.GetPageCount()))}); err != nil {
			return err
		}
		if err := stream.Send(&pb.BookField{Key: "title", Value: book.GetTitle()}); err != nil {
			return err
		}
	}

}

func NewServer() {
	log.Println("Starting Server..")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterInventoryServer(grpcServer, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

// >>>>>>> Util functions <<<<<<<<<

func getSampleBooks() []*pb.Book {
	sampleBooks := []*pb.Book{
		{
			Title:     "The Hitchhiker's Guide to the Galaxy",
			Author:    "Douglas Adams",
			PageCount: 42,
		},
		{
			Title:     "The Lord of the Rings",
			Author:    "J.R.R. Tolkien",
			PageCount: 1234,
		},
	}
	return sampleBooks
}

// StringPtr returns a pointer to the passed string.
func StringPtr(s string) *string {
	return &s
}
