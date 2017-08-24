package main

import (
    "fmt"
    "log"
    "net/http"
    _ "net/http/pprof"
    "runtime"
    "time"
    "github.com/julienschmidt/httprouter"
)

func count() int {
    count := 20000000
    return count
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Println("INDEX PAGE")
}

func Scenario1(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    count := count()
    start := time.Now().UnixNano() 
    for i := 0; i < count; i++ { 
    }
    end := time.Now().UnixNano() 
    process := int(end-start)
    log.Println("PROCESS TIME :",process)
    log.Println(process/count)
    fmt.Println("------------------------------------------------------------------------")
}

func Scenario2(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    var m runtime.MemStats
    count := count()
    runtime.ReadMemStats(&m)
    fmt.Println("START HEAP : ",m.HeapAlloc)
    return
    for i := 0; i < count; i++ {
        runtime.ReadMemStats(&m)
        fmt.Println("PROCESS HEAP : ",m.HeapAlloc)
        
    }
    fmt.Println("------------------------------------------------------------------------")
}

func Scenario3(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    var m runtime.MemStats
    count := count()
    runtime.ReadMemStats(&m)
    fmt.Println("START Allocation IN KB : ",m.Alloc/1024)
    return
    for i := 0; i < count; i++ {
        runtime.ReadMemStats(&m)
        fmt.Println("PROCESS Allocation IN KB: ",m.Alloc/1024)     
    }
    fmt.Println("------------------------------------------------------------------------")
}

func Scenario4(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    var m runtime.MemStats
    count := count()
    runtime.ReadMemStats(&m)
    fmt.Println("START Allocation : ",m.Alloc/1024)
    fmt.Println("START HEAP : ",m.HeapAlloc)
    return
    start := time.Now().UnixNano() 
    for i := 0; i < count; i++ {
        runtime.ReadMemStats(&m)
        fmt.Println("PROCESS Allocation : ",m.Alloc/1024)
        fmt.Println("PROCESS HEAP : ",m.HeapAlloc)
        
    }
    end := time.Now().UnixNano()  
    process := int(end-start)
    log.Println("PROCESS TIME :",process)
    log.Println(process/count)
    fmt.Println("------------------------------------------------------------------------")
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/Scenario1", Scenario1)
    router.GET("/Scenario2", Scenario2)
    router.GET("/Scenario3", Scenario3)
    router.GET("/Scenario4", Scenario4)
    log.Fatal(http.ListenAndServe(":8080", router))
}