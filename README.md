# Small_Distributed_System
the practice with Go to build distributed systems

## Features:
- one master node, several slave nodes
- communicate through TCP protocols
- messages are json packets
- each node has its unique ID
- first node is default the master cluster. If it dies, re-elect a master

