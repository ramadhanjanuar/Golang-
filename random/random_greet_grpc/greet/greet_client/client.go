package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"grpc-course/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("I'm client")

	interceptor := NewAuthInterceptor(getToken())
	con, err := grpc.Dial(
		"localhost:50010",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
	)

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer con.Close()

	greetService := greetpb.NewGreetServiceClient(con)
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ramdhan",
			LastName:  "Januar",
		},
	}
	res, err := greetService.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call Greet RPC: %v", err)
	}
	log.Println(res.Result)
}

func getToken() string {
	// internalServiceToken := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIwMDAwMDAiLCJleHAiOjAsImp0aSI6ImRmYTtsa2Fqc2ZranNhZGtsZmpkc2FrbCIsImlhdCI6MTU4NzY2NzAyNSwiaXNzIjoiaW50ZXJuYWxzZXJ2aWNlY29yZV8zLjAiLCJuYmYiOjE1ODc2NjcwMjUsInN1YiI6ImludGVybmFsc2VydmljZWNvcmVfMy4wIn0.M9aVKQVMFgINWTFp1HtCu0CwEBrWGYaD763AG7Nh902FcaUw_xSBoOS6C84iq85gYI09TKv9XTpVtKtbjLX1NnoFs0yA7H8A-9Hw123sGM51eAnIY3-9zk9zgtTFRqV4GqJ6wfR6gTJQCff9q0H8SFyzdXkafvJ8t92lco8LzO4`
	url := "http://localhost:9096/oauth2/token"
	method := "POST"

	payload := strings.NewReader("grant_type=password&username=istiq&password=qwerty&client_id=000000&client_secret=999999")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var data map[string]interface{}
	json.Unmarshal(body, &data)

	//token := fmt.Sprintf("%v+adklsfjdsk", data["access_token"])

	return `fmk9i0UkN7KeY3oFVlUs1uA5JX4OTJQLAkA61trHVbaNjfTvwX6nIyrZiA5VNR`
}
