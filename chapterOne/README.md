### 1.4: gif.go  
usage: 
```
go build gif.go  
./gif > filename.gif
```
Creates a simple black and white gif based on sinusoidal oscillation.  
<img src="https://github.com/robertnowell/golang/blob/master/chapterOne/Screen%20Shot%202017-02-06%20at%201.11.44%20PM.png" 
alt="fdf" width="200" height="200" border="10"/>

### 1.5: fetch.go  
usage:  
```
go build fetch.go  
./fetch url  
```
Inspired by Unix curl, this program fetches data from a webpage and prints it to standard output.  

### 1.6: fetchall.go  

"A goroutine is a concurrent function execution. A channel is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine (18)." 

This program takes in multiple URLs as input and processes their data concurrently, each in separate goroutine. The total time taken is the amount of time taken for the longest goroutine, rather than the sum of the time taken for all goroutines.  
The output for this program is the amount of time taken, the quantity of bytes read, and the website URL.  
```
go build fetchall.go
./fetch URL URL ... URL
```

### 1.7 server1.go
Using this program will rely on the fetch program. Server1 spins up a quick local server and allows for the retrieval of some metadata from that server. 
```
go run server1.go
go build fetch.go
./fetch http://localhost:8000
```

the output will look something like the following:
```
GET / HTTP/1.1
Header["User-Agent"] = ["Go-http-client/1.1"]
Header["Accept-Encoding"] = ["gzip"]
Host= "localhost:8000"
RemoteAddr = "127.0.0.1:59911"
```

