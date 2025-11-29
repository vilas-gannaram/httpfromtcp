## Web Dev → Networking Translation

### What You Already Know (HTTP/Web):

```
Browser (Client)                 API Server
     |                               |
     |------ GET /api/users -------->|  Request
     |<----- 200 OK + JSON ----------|  Response
     |                               |
```

- **Client makes request** → waits for response
- **Server responds** with data
- **HTTP** handles all the connection stuff for you

### What We've Been Doing (Raw TCP/UDP):

Think of it as **building the layer BELOW HTTP**. HTTP is built on top of TCP!

```
HTTP Request/Response
        ↓
    TCP/UDP (what we're coding)
        ↓
    Internet
```

## The Analogy

### Your Web Dev World:

```javascript
// Client (browser/fetch)
fetch('http://api.example.com/users')
  .then(res => res.json())
  .then(data => console.log(data))

// Server (Express/Node)
app.get('/users', (req, res) => {
  res.json({ users: [...] })
})
```

### What We Just Built:

**TCP Server** = Like an Express server, but more low-level
```go
// Listens for connections (like app.listen())
listener.Accept()  

// Handle each connection (like route handler)
for line := range getLines(conn) {
    fmt.Println(line)  // Process the "request"
}
```

**UDP Client** = Like `fetch()`, but fire-and-forget
```go
// Send data (like POST request, but no response expected)
conn.Write([]byte(line))
```

## Key Differences

| Web Dev (HTTP) | Raw Networking (TCP/UDP) |
|----------------|--------------------------|
| `fetch('/api')` | `net.Dial("tcp", "host:port")` |
| JSON request/response | Raw bytes |
| Always request → response | TCP: yes, UDP: no response |
| Browser handles connection | You handle everything |
| Built on TCP | You're building the TCP layer |

## Why Learn This?

As a web dev, you might need this for:

- **WebSockets** (persistent connections, chat apps)
- **Real-time features** (gaming, live updates)
- **Custom protocols** (not HTTP)
- **Understanding performance** (why is my API slow?)
- **Microservices** that talk to each other directly

## The Layers

```
Your Web Apps (React, API routes)
         ↓
    HTTP/HTTPS
         ↓
    TCP (what we built)  ← You're here now!
         ↓
    IP/Internet
         ↓
    Physical Network
```

Think of what we're doing as **understanding the engine under the hood of your car** (HTTP). You don't need it to drive (build web apps), but it helps you understand why things work the way they do!

Does this make more sense now? Want me to show you how to build a simple HTTP server from scratch using TCP to see how it all connects?
