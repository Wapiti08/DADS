package main

import (
    "flag"
    "net"
    "encoding/json"
    "math/rand"
    "time"
)


// information/metadata about node
type NodeInfo struct {
    NodeId int  'json:"nodeid"'
    NodeIp string   'json:"nodeip"'
    Port string 'json:"port"'
}

// standard format for a Rq/Rp for adding node to cluster
type AddToClusterMessage struct {

}

// format print for node info

// format print for Rq/Rp info

// entrypoint for system
func main() {
    clusterip := flag.String("clusterip", "127.0.0.1","if not connected to another cluster, set itself as master")
    port := flag.Int("port"."8001", "connect to this port")
    masterNodeError := flag.Bool("masterNodeError")
    flat.parse()
    // generate the unique id for cluster with rand and time
    rand.seed(time.Now().UTC().UnixNano())
    randid := rand.Intn(999999)
    // assign values to node
    ip, _ := net.InterfaceAddrs()
    source := NodeInfo{Nodeid: randid, NodeIp:ip, Port: *port}

}
/* format the json packet to send requests */

// unexported function
func connectToCluster() {
    // protocol type, address, timout duration
    conn, err := net.DailTimeout()
}

func listenOnPort() {
    // protocol type, address
    In, err := net.Listen("")
}
