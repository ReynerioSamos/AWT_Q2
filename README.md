# Middleware Functions in Go

Reynerio Samos - 2018119235 - 2024-10-18

**Sources:**
- [Using Express middleware by Express](https://expressjs.com/en/guide/using-middleware.html)
- [Exploring Middlewares in Go by Vishnu Bharathi](https://vishnubharathi.codes/blog/exploring-middlewares-in-go/)

---
## Definition 
Middleware functions are functions that have access to the request object `req` and response object `res` and the next middleware function in a request-response cycle (such as in an incoming HTTP request and outgoing HTTP response).

Such objects in JavaScript are:
```Node
// Request Objects in HTTP 
req.Method // GET, POST, PUT
req.URL
req.Headers
req.Body
req.Params
req.Query
// Response Objects in HTTP
res.Status // HTTP Status code
res.Headers
res.Body
res.Write() //Method for writing
res.End() // Method for ending response
```

Middleware functions can:
- Execute any code
- Make changes to the request and response objects
- End the request-response cycle
- Call the next middleware function in the stack
	- If not, then it calls `next()` to pass control to the next middleware function in main, or else it will be left hanging

Some levels or classifications of middleware functions are:
- Application-level middleware
	- Wrap around entire application regardless of endpoint
	- Eg: Logging requests
- Router-level middleware
	- Executed on requests that match a specific route or pattern
	- Eg: Authentication handling
- Error-handling middleware
	- Executed when an error occurs in the request-response cycle
	- Eg: Closing connections to requests with exceedingly large payloads, to avoid wasting server resources
- Built-in middleware
	- Middleware provided by framework's library and packages
	- Eg: In GO: `net/http` 
- Third-party middleware
	- Middleware provided by external libraries or packages
	- Eg: [Julien Schmidt's HttpRouter](https://github.com/julienschmidt/httprouter)in Go: `"github.com/julienschmidt/httprouter"`
