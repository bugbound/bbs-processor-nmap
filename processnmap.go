package main

import (
    "github.com/tomsteele/go-nmap"
    "strings"
    "fmt"
    "io/ioutil"
    "os"
    "net/http"
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
            fmt.Println(ports[pi].PortId, " - ", ports[pi].State.State, " - ", ports[pi].Service.Name)
            addServiceToApi(host.Addresses[0].Addr, ports[pi].PortId, ports[pi].State.State, ports[pi].Service.Name)
        }
        fmt.Println("")      
        
    }
    
}

func addServiceToApi(ip string, portid int, state string, name string) {
    url := "http://servicestore-service:7010/api/service_store"
    jsonbody := fmt.Sprintf(`{"ip":"%s","port":"%d","state":"%s","name":"%s"}`, ip, portid, state, name)
    req, err := http.NewRequest("POST", url, strings.NewReader(jsonbody))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    //fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
    
}
