# Go-Practice

This repository will focus on learing Go.

## Progress

### 1> Hello World !

> Go packages and directory strucure for compilation
```sh
> /nameofgocomppath
> /nameofgocomppath/src
> /nameofgocomppath/src/helloworld/main.go
> /nameofgocomppath/bin
> /nameofgocomppath/pkg
```
```sh
$ export GOPATH="/name/of/go/comp/path"
$ go install package1
$ go install package2
```
> Go packages and directory strucure after compilation
```sh
> /nameofgocomppath
> /nameofgocomppath/src
> /nameofgocomppath/src/helloworld/main.go
> /nameofgocomppath/bin
> /nameofgocomppath/bin/helloworld
> /nameofgocomppath/pkg
```

### 2> Strings

Different ways of formatting and printing stuff.


### 3> Console Input

Taking in User Input from the console with bufio and fmt's Scanln.
Apart from this we also see how to handle the input which we return from the console.


### 4> Variables
Types for variables can be set explicitly or implicitly.
"=" operator is for explicit and ":=" is for implicit.
As Go is a statically typed language these values cannot be changed during runtime.

Constants in Go use "=" for both implicit and explicit declaration.
"time" package allows to use different time operations.

#### Arithmatic Operations
+ , - . * , / , %
"&"  Bitwise AND
"|"  Bitwise OR
"^"  Bitwise XOR
"&^" Bit Clear
"<<" LeftShift
">>" Rightshift

Mathematics requires the same types for the operators to function properly

#### Types
bool
string 
uint8-64 (number of bits)
int8-64 (number of bits)
byte 
int
uintptr (large enough to hold the bit pattern of any pointer)
float32-64
complex64-128 (Complex Numbers)
Arrays
Slices
Maps
Structs
Functions
Interfaces
Channels
Pointers

### 5> Pointers

These are the same as they would be in C, C++.
Similarly "&" would be to get the address of the value and "*" would be to point to the value.

### 6> Arrays

#### Arrays

Arrays are declared in the following way
```sh
var arr [1]string;
arr[0] = "value"
```

#### Slices

Slices are no different from Arrays with the only difference being that their size is not fixed, they are declared differently and they make up arrays.
```go
var sli1 []string;
var arr [5]int = [3]int{1,2,3,4,5};
sli2 := append(arr[1:3]);
```

### 7> Memory Allocation

- The runtime operates in its own background threads. Memory is managed completely in the background and no allocation of deallocation of memory needs to be handled by the programmer. 
- make(), new() are used to create instances of Maps, Slices, Channels
```sh
new()  // Allocates but does not initalize memory, returns address
make() // Allocates and initalizes memory, returns address
```
- The Go Garbage Collector deallocated memory for objects which are out of scope or set to nil.

#### Maps

So maps can be declared in the following way
```go
mp := make(map[int]string) 
```
In the above example we are creating maps which would have the keys as integers and the value for the key as a string. 
Maps are similar to dictionaries in Python with the difference being that the Keys and Values can be of only one type.

#### Structs

Structs are the same as they are in C. They are declared as follows.
```go
type Dog struct {
    Breed string
    Weight int
}
```
They are meant to be used like classes, but Go doesn't have the advanced OOP concepts such as inheritance and such. The only OOP Concepts it has are Encapslation and Type Memeber functions.
Go thorugh this [link](https://www.ardanlabs.com/blog/2013/07/object-oriented-programming-in-go.html) for more information about Object Oriented Programming in Go.


### 8> Functions

Basic functions in Go are little different than the normal way we end up writing functions.
```go
func function(value int) int {
    integervalue := 100
    return integervalue
}

func function(value int, value2 string) (int, string){
    return 10, "word"
}

func function(value int, value string) (number int, sentence string){
    number = value+10
    sentence = value+"adddedForReturn"
    return
}
```
#### Functions for Custom Types

Custom Types would be structs which would act like Types which would be made by you.
```go
type Dog struct {
    Breed string
    Weight int
    Sound string
}
```

Functions for Custom Types
```go
func (d Dog) Speak(){
    fmt.Println(d.Sound)   
}
```

### 9> Interfaces

Interfaces are basically wrapper structs. 
* Interfaces should be declared along with the structs that need to be declared which can use these interfaces.
* If there are functions which are a part of the interface and would be inherited by the structs which we are taking about, the functions need to be declared for each of the structs.
* We can say that interfaces help to define common attributes or functions for similar objects or structs.

### 10> Error Handling

Go does not ahve the usual error handling stuff we ahve in other languages.
Instead we ahve the "errors" package with which we can make and erro object which allows us to basically form a template for handling errors properly
```go
import "errors"
import "fmt"

func main() {
    MyError := errors.New("Error !!")
    fmt.Println(MyError.Error()) // gives the same output as fmt.Println(MyError)
}
```
* We usually define a function to return an error along with any toher value it might return. 
* Error handling goes hand in hand with conditionals as there is no exception handling of sorts.


### 11> Really Simple Web App

We use the "net/http" package to initialize a http socket and write data to response using the following functions.

```go
func main(){
    http.HandleFunc("/path",handlerfunction)
    http.ListenandServe(":3000",nil) // We would be listening on the port 3000
}
func handlerfunction(w http.ResponseWriter, r *http.Request){
    Fprintf(w, "Writing to the Resopnse") 
}
```

### 12> Methods / Functions and their Receivers

Methods in Go have receivers which are basically the object/struct they work on. 
Revceivers are basically "this" is Java or "self" in Python.

#### Value Receivers
We consider using these in case we need to use the values which have been declared in a n instance of a particular custom type.
```go
func (c Car) function(){}
```

#### Pointer Receivers
We consider using these if we want to make changes to the already existing values in the instance of the custom type.
```go
func (c *Car) function(){}
```

### 13> Simple Web Scraping
Here we are going to start looking into how we can get data from the internet and parse it in a simple webscraping web app.
```go
import (
    "fmt"
    "net/http" // for making the get request
    "io/ioutil" // foir reading the response and getting the output in bytes
    "encoding/xml" // This library allows us to parse xml very nicely
)
```
> Important functions
```go
http.Get() // returns response and error
ioutil.ReadAll() //returns bytes and error
xml.Unmarshall() // based on reflection, need to be give the bytes and the struct where we will place our data
```
#### Struct tags

```go
type Urlset struct {
    Urls []Url `xml:"url"` // The third element here is a struct tag which would be used by reflection which we would look into further
}
type Url struct {
    Location string `xml:"loc`
}
func (u Url) String() string { // This is a GOhack as I would like to call it which would allow for returning a string instead of Url Type to the Urlset struct
    return u.Sprintf(u.Location)
}
```

### 14> HTML Templates

In go we can create web apps with the use of templates
```go
import "html/template"
```
> Important Functions
```go
t, _ := template.ParseFiles("template.html")
t.execute(w,data)
```
> html template example
```html
<h1>
    {{.Title}}
</h1>

{{range $key,$value := .News}}
    <h3>
        {{$key}}
    </h3>
    <p>
        {{$value.Url}}
    </p>
    <p>
        {{$value.Keyword}}
    </p>
{{end}}
```
This would help a lot in dynamically generating components for the webpage that one would be building. Apart from that, this this available in go as its own inbuilt package.


### 15> Goroutines Basics

Goroutines are nothing but lightweight threads of execution. In GO we use them to run different tasks concurrently. When we talk about concurrency we are not taling about parallelism. Parallelism means that multiple tasks run a the same time whereas concurrency means that multiple tasks are dealt with at the same time and based on resource availability they are provided the required resources.

To convert a particular task of execution such as a function we insert the "go" keyword at the start of the function call.

```go
function("doing something") //before
go function("doing something") //after
```
Once we do the above, what will happen is that the particular function would be running in the background and with it other processes will run concurrently.
* One thing to remeber is that if all tasks have been initialized as 
goroutines as the above and the program runs, it is possible that the program may exit as there is no such task which has to be be executed for sure. 
* For the above problem we would need to syncronize the threads properly so that we have a some kind of syncronized execution.

```go
import "sync" // for syncronization utilities

wg = sync.WaitGroup // WaitGroup struct
```

* Now we have initialized "wg" as a WaitGroup to which we can Add a process which would be getting executed and once we are Done with a certain process the counter in the WaitGroup will be decremented.

```go
func function(s string) {
    defer wg.Done()
    fmt.Println(s)
}

func main(){
    wg.Add(1) // increment wait counter
    go function("doing something1")
    wg.Add(2)
    go function("doing something2")
    go function("doing something3")
    wg.Wait() 
}
```

* In the above if you notice we have 'defer'ed the statement wg.Done() as we would want Done() to run for sure after the function executes.
* In other words if there is an error of some sort which happens before the Done() function then the waitgroup will reamin hanging.

### 16> Channels

Channels allow for connecting the concurrent parts of a Go Program.
```go
out chan<- int // send-only channel
in  <-chan int // receive-only channel
any chan int   // send or receive
```
#### Blocking

Unlike the time when we used the "sync" package to take care of the syncronization to block code to execute with the Wait, Add, Done functions, channels block their own goroutine till they get completely executed.


### 17> Adding Concurrency to a Web APP

To add concurrency to an app we would need to breakdown each of the tasks it is performing and how the data is moving between goroutines.
In the Example App which we have made for this which is a news aggregator has an important routine which takes the news data an daggregates it onto a web page. 
We have converted the whole thing into two different routines considering main and the news routine which would send over the data through a channel.

#### Converting a function to a goroutine
* Insert "go" at the start of the function(no return, value would be passed on through a channel) execution.
* Initialize Waitgroup wg.
* Insert Add(1) on the top of the goroutine which we have defined to increment waitgroup parameter.
* Append Done() to the goroutine function (to decrement WaitGroup) and defer it so that it runs at the end whatever the issue.
* At the end of defining all goroutines append Wait() to create Blocking to wait for all the goroutines to finish executing.
* Initialize the channel which we would be passing between the goroutines.
* Pass the channel to the goroutine function.
* Add value to the channel through the function.
* Close the channel after the WaitGroup Wait() exits.
* Now the channel is populated and the values can be used.

### 18> Streaming IO
###### [io.Reader in depth](https://medium.com/@matryer/golang-advent-calendar-day-seventeen-io-reader-in-depth-6f744bb4320b)
###### [A simple beginners tutorial to io.Writer in Golang](https://medium.com/@as27/a-simple-beginners-tutorial-to-io-writer-in-golang-2a13bfefea02)
###### [Streaming IO in Go](https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185)


#### io.Reader
io.Reader should be used whereever posible where we take byte arrays or strings as then we can use it to pass any kind of information. A lot of important Reader types return an io.Reader type object, eg. strings.NewReader, os.Open, bytes.Buffer. 

##### Method of Use
We can Read() from the io.Reader in use by getting the stream in a buffer periodically and using it for our purpose. We can even create our own reader which can read elements in a certain way. This is good in case we need to preprocess our data, we can use io.Reader to preprocess it before we take it in as input.


#### io.Writer
With io.Writer allows us to write a serialized stream of data into a buffer which we can read with io.Reader.

##### Method of Use
We use io.Writer to simply transfer what we have written to the transfer buffer to the Destination we want to write our data to.

#### io.Copy
This function allows us to chain source reader to target writer and it is very easy to use as we just have to pass in the writer annd the reader where the reader or the place we ahve to gfe tht e data from has to be initialized with a new instead of a make or we can get an error as make initializes without a pointer and new does.

### 19> Protocol Buffers
###### [Basics of Protocol Buffers in Golang](https://tutorialedge.net/golang/go-protocol-buffer-tutorial/)

#### *.proto Files
This is where we have to code our schemas. The way we define payload in these .proto files is using schemas called messages and they are nomianally ordered with a distinct identifier given to each of them.

```go
syntax="proto3";
package main;
message SocialFollowers {
    uint32 youtube = 1;
    uint32 twitter = 2;
}
```

Now they can be imported in other .proto files as below.

```go
syntax = "proto3";
package main;
import "social.proto";
message Person {
    string name = 1;
    uint32 age = 2;
    main.SocialFollowers SocialFollowers = 3;
}
```

#### Compilling .proto Schemas
The *.proto files are compiled to give us *.pb.go files which are the actual go files which would be used. This go code which is generated contains Getters for each of the messages which are described.

To compile .proto files we need to use the `protoc` command and we also need the `protoc-gen-go` command which will allow for the generation of the *.pb.go files.

```sh
cd 019-protobufbasics
export GOPATH=$(pwd)
export PATH=$PATH:$GOPATH/bin

go get -u github.com/golang/protobuf/proto

cd src
protoc --go_out=. social.proto person.proto
```

After running the above command you should see that the *.pb.go files have been generated. Run the command below to see how the app runs.

#### Checking out protobuf data representation

```sh
go install simpleApp
cd 019-protobufbasics/bin
./simpleApp
```

The above command should return how protobuf looks at data and what the Getters return.

```
[10 6 86 105 98 104 97 118 16 22 26 4 8 40 16 30]
22
Vibhav
youtube:40 twitter:30
```

### 20> RPC Server and Client implementation
###### [RPC Server and Client](https://medium.com/@OmisNomis/creating-an-rpc-server-in-go-3a94797ab833)
RPC or Remote Procedure Calls are exactly what the name suggests. One can remotely call procedures which aren't present locally. So what this means is that a system in this case is able to invoke procedures or methods on a remote server. We are going to use the `net/rpc` package in go to help us make this happen.

For the RPC Protocol, functions need to be specified in a particular way in an interface which would be registered to the RPC service we are running.
These functions are supposed to have only two arguments, and one return(`error`). The first argument is the payload and the second is the `reply` which points to a `var reply <typeofreply>`. 

Once the RPC is registered, we create a HTTPListener where these requests arrive and are handled by the RPC HTTPHandler instantiated beforehand.

On the client side we need the datatypes of the objects which are being returned in the reply. We make Calls to the RPC Server using a RPC HTTPClient by passing the method name in the RPC Call which would be defined on the interface containing the methods.

### 21> The `context` Package
###### [justforfunc #9: The Context Package](https://www.youtube.com/watch?v=LSzR0VEraWw&t=1546s)

Context is used in HTTP servers and clients to make them less wasteful and more efficient. It is usually used in cases when one needs to propogate a Cancellation Request. In a scenario where we make a certain request, it should be possible for us to cancel that request from our end without wasting resources on the request. 

HTTP requests have a context which can be used to supplement our needs and we can also attach a predefined context to a New HTTP Request which and see to it that we are propogating cancellation whenever we want. This can be used to cancel the HTTP requests on certain events which might happen while handling the same

### 22> Protobuf Server and Client implementation
###### [justforfunc #30](https://www.youtube.com/watch?v=_jQ3i_fyqGA)

We have extended upon `019-protobufbasics` and learnt how to use encode the data, store it and vice-a-versa in the form of a todoapp, where we add the tasks in a file which and then read it from there when we need to.

### 23> GRPC Basics
###### [justforfunc #31](https://www.youtube.com/watch?v=uolTUtioIrc&t=322s)

Extending upon what we have done in `022-protobuf`, we are going to convert the entire application that we have made into a GRPC Server and also write a Client for the same which can invoke Remote Procedure Calls. Before this, we had done RPC which is the protocol on which GRPC is build. The difference between GRPC and RPC is, GRPC uses protobuf as it's medium of communication and all the schema is defined in protobuf itself.

We modify the `todo.proto` file to have a new message type.
```go
message TaskList {
    repeated Task tasks = 1;
}
```
Here we know that Task has already been defined as another message. We are using it like a type in this case and then saying that we are expecting a slice of Tasks. That is why we have mentioned the `repeated` keyword which basically points this out.

```go
service Tasks {
    rpc 
}
```
If we add the above code to the .proto file and try to compile it with `protoc -I todo.proto --go_out=.`. We can notice that the compilation hasn't worked. Now this is because, when we say service we are specifically saying that it needs to be a GRPC service. In this case we are going to pass a grpc plugin which will help compile the service and the command would be `protoc -I todo.proto --go_out=plugins=grpc:.`

```go
type taskServer struct{}

func (s taskServer) List(ctx context.Context, void *todo.Void) (*todo.TaskList, error) {
	return nil, nil
}

func main() {
	var tasks taskServer

	srv := grpc.NewServer()
	todo.RegisterTasksServer(srv, tasks)
}
```
Similar to RPC we have to register a GRPC Server with the correct type interface containing all the functions we need to make the GRPC calls.

We define the .proto file with the services and messages which have to be passed, as done below.
```go
syntax = "proto3";

package todo;

message Task {
    string todo = 1;
    bool done = 2;
}

message TaskList {
    repeated Task tasks = 1;
}

message Text {
    string text = 1;
}

message Void {}

service Tasks {
    rpc List(Void) returns(TaskList) {}
    rpc Add(Text) returns(Task) {}
}
```
Once we define and compile the above with the grpc plugin, we register the GRPC Server function and get the client definition from the generated .pb.go file. 
Everything else we do here is similar to what we do in 020-rpc and 022-protobuf. Even the way we encode and save the protobuf data. 

GRPC allows us for faster communication with HTTP/2 over TCP/IP with a lot more functionality than rest allowing for bidirectional streams as well a lot more functionality than the usual GET, POST, PUT, PATCH.

### 24> Developing and testing Kubernetes applications using `client-go` 
###### (Unit Testing with Kubernetes Client Go)[https://medium.com/@e_frogers/unit-testing-with-kubernetes-client-go-283b11aaa7db]
###### (Testing Kubernetes Go Applications)[https://itnext.io/testing-kubernetes-go-applications-f1f87502b6ef]

## Things to Cover
#### Reflection [1](https://blog.golang.org/laws-of-reflection), [2](https://medium.com/capital-one-developers/learning-to-use-go-reflection-822a0aed74b7)
#### Goroutine Leak [1](https://medium.com/golangspec/goroutine-leak-400063aef468)
#### Syncronization [1](https://medium.com/golangspec/synchronization-queues-in-golang-554f8e3a31a4), [2](https://medium.com/@fharding/mutexes-in-go-df0c30a138e4), [3](https://medium.com/golangspec/sync-rwmutex-ca6c6c3208a0), [4](https://www.sohamkamani.com/blog/2017/08/24/golang-channels-explained/)
#### Conversions and Interfaces [1](https://medium.com/golangspec/conversions-in-go-4301e8d84067), [2](https://medium.com/golangspec/interfaces-in-go-part-ii-d5057ffdb0a6), [3](https://medium.com/golangspec/interfaces-in-go-part-iii-61f5e7c52fb5) 
#### Protobuf [1](https://codeburst.io/protocol-buffers-part-1-serialization-library-for-microservices-37418e72908b) 
