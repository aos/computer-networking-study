### Socket API

1. `socket()`
2. server: `bind()`
3. server: `listen()`
4. client: `connect()`, server: `accept()`
5. `send()`/`recv()`
6. `close()`

### Demultiplexing

Each layer in the protocol adds its own header information. Each header will
have the next layer's (upstream) protocol type, called demultiplexing keys. For
example,
1. Ethernet: Ethertype field (IP, ARP)
2. IP: IP protocol field (TCP, UDP)
3. TCP: TCP port number (80 - HTTP, 53 - DNS, etc.)

### Internet Reference Model

```
Application (7)   -- HTTP, RTP, DNS
-----------
Transport (4)     -- TCP, UDP
---------
Internet (3)      -- IP
--------
Link (1, 2)       -- 3G, Ethernet
```
