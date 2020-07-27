package main

import (
    "github.com/tomsteele/go-nmap"
    
    "fmt"
    "io/ioutil"
    "os"
)


func main() {
    fmt.Println("w00p")
    b, err := ioutil.ReadFile(os.Args[1]) // just pass the file name
    if err != nil {
        fmt.Print(err)
    }

    bah, nmaperr := nmap.Parse(b)
    if nmaperr != nil {
        fmt.Print(nmaperr)
    }
    
    for currentIndex := range bah.Hosts {
        var host = bah.Hosts[currentIndex]
        fmt.Println(host.Addresses[0].Addr)  
        
        var ports = host.Ports
        for pi := range ports {
            fmt.Println(ports[pi].PortId, " - ", ports[pi].State.State)
        }
        fmt.Println("")      
        
    }
    
}

