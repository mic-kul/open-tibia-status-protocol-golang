package main

import (
    "fmt"
    "io/ioutil"
    "net"
    "time"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
        os.Exit(1)
    }
    service := os.Args[1]
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err)
    timeoutDuration := time.Now().Add(5 * time.Second)
    conn.SetDeadline(timeoutDuration)
    conn.SetReadDeadline(timeoutDuration)
    conn.SetWriteDeadline(timeoutDuration)
    var msg = []byte{0x06, 0x00, 0xFF, 0xFF, 0x69, 0x6E, 0x66, 0x6F}
    _, err = conn.Write(msg[:])
    checkError(err)
    result, err := ioutil.ReadAll(conn)
    checkError(err)
    fmt.Println(string(result))
    os.Exit(0)
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}