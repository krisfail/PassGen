package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
func init() {
    rand.Seed(time.Now().UnixNano())
}
 
func main() {
    fmt.Printf("%s", randomText(10))
}
 
func randomText(length int) string {
    const charSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    b := make([]byte, length)
    for i := range b {
        b[i] = charSet[rand.Intn(len(charSet))]
    }
 
    return string(b)
}