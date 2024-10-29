package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {	
	method,headers,data,authType,token,url:=parseFlags()

	reqBody,err:=createRequestBody(method, data)
	errCheck(err, "Error creating request body")
	
    req,err:=http.NewRequest(method, url, reqBody)
	errCheck(err, "Error creating request")

	addHeaders(req, headers)
	authenticateRequest(req, authType, token)

	client:=&http.Client{}
	res,err:=client.Do(req)
    errCheck(err, "Error sending request")
	defer res.Body.Close()
	
	body, err := io.ReadAll(res.Body)
	errCheck(err, "Error reading response body")

	fmt.Printf("Response Status: %s\n", res.Status)
	fmt.Println("Response Body:", string(body))
}

func parseFlags()(string,string,string,string,string,string) {
	method := flag.String("method", "GET", "HTTP method")
	headers := flag.String("header", "", "Request Headers")
	data := flag.String("data", "", "Payload")
	authType := flag.String("auth-type", "", "Authentication type (bearer/basic)")
	token := flag.String("token", "", "Token or username:password for authentication")
	flag.Parse()
	url := flag.Arg(0)

	if url==""{
		fmt.Println("No URL: Required parameter missing")
		os.Exit(1)
	}
	
	return *method, *headers, *data, *authType, *token, url
}

func createRequestBody(method string, data string)(io.Reader, error) {
	if method=="POST" || method=="PUT" {
		if data==""{
			return nil, fmt.Errorf("data is required for %s requests", method)
		}
		return bytes.NewBuffer([]byte(data)),nil
	}
	return nil,nil //no request body if request not POST or PUT
}

func addHeaders(req *http.Request, headers string) {
	if headers == "" {
        return 
    }
    headerPairs := strings.Split(headers, ",")
    for _, h := range headerPairs {
        keyVal := strings.SplitN(h, ":", 2)
        if len(keyVal) != 2 {
            fmt.Printf("Invalid header format for: %q. Expected format 'Key: Value'\n", h)
            continue
        }
        key := strings.TrimSpace(keyVal[0])
        value := strings.TrimSpace(keyVal[1])
        req.Header.Set(key, value)
    }
}

func authenticateRequest (req *http.Request, authType, token string) {
	switch authType {
	case "bearer":
		req.Header.Set("Authorization", "Bearer "+token)
	case "basic":
		parts:=strings.SplitN(token, ":", 2)
		if len(parts)==2{
			req.SetBasicAuth(parts[0],parts[1])
		} else {
			fmt.Println("Invalid format for basic authentication. Expected 'username:password'")
			os.Exit(1)
		}
	}
}

func errCheck(err error, msg string) {
	if err!=nil{
		fmt.Printf("%s: %v\n", msg, err)
		os.Exit(1)
	}
} 

