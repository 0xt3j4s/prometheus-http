package main

import (
   "fmt"
   "net/http"

   "github.com/prometheus/client_golang/prometheus"
   "github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(
   prometheus.CounterOpts{
       Name: "ping_request_count",
       Help: "No of request handled by Ping handler",
   },
)

func ping(w http.ResponseWriter, req *http.Request){
   pingCounter.Inc()
   fmt.Fprintf(w,"pog!")
}

func main() {
   prometheus.MustRegister(pingCounter)
   http.HandleFunc("/ping",ping)
   http.Handle("/metrics", promhttp.Handler())
   fmt.Println("Server is listening on port 8090...")
   http.ListenAndServe(":8090", nil)
   
}