## Standard kid

Cloud-native extract kid from JWT (using no dependencies) for multi-tenant gateways

### Example Usage
[https://go.dev/play/p/-tBycmHbXC2](https://go.dev/play/p/-tBycmHbXC2)


```go
package main

import (
    "fmt"
    "standardkid"
)

func main() {
    token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImtleV9pZCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
    kid, err := standardkid.ExtractKidFromJWT(token)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("kid:", kid)
}
```
