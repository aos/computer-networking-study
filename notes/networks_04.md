# Network Layer

Shortcomings of switches:
1. don't scale across large networks
2. don't work across more than one link layer technology (3G, 802.11, etc.)
3. don't give much traffic control

Routing vs. Forwarding
- **Routing**: process of deciding which direction to send traffic.
    Network wide (global) and expensive
- **Forwarding**: process of sending a packet on its way.
    Node process (local) and fast

## Network Service Models

1. Datagrams (postal letters)
2. Virtual circuits (connection-oriented)

### Store-and-Forward Packet Switching
- Routers receive a complete packet, storing it temporarily if necessary before
  forwarding it onwards (internal buffering)
- Use statistical multiplexing to share link bandwidth over time

### Datagram Model

Each router has a forwarding table keyed by **address**.

## IP (Internet Protocol)

Network layer of the internet, uses datagrams. IPv4 carries 32-bit addresses on
each packet (often 1.5 KB).

### Virtual Circuits

Packets only contain a short label to identify the circuit. Labels don't have
any global meaning, only unique for a link.

Each router has a forwarding table keyed by **circuit**.

- MPLS (Multi-Protocol Label Switching): virtual-circuit like technology widely
  used by ISPs
- ISP sets up circuit inside their backbone ahead of time
- Adds MPLS label to IP packet at ingress, undoes at egress

## Internetworking

How to combine multiple networks together into a larger network.

- Pushes IP to be a "lowest common denominator" protocol, as the need to
  connect disparate networks that support different services (QOS, security,
  etc.) exists
- Asks little and gives little for lower and higher layer networks

## IP Prefixes

Addresses are allocated in blocks called "prefixes". Addresses in an "L-bit"
prefix have the same top L bits. There are 2^(32-L) aligned on 2^(32-L)
boundary.

Example: 10.0.0.0/24 -> 32 - 24 = 8 bits (256 addresses): 10.0.0.0 - 10.0.0.255

## IP Forwarding

- IP addresses on one network belong to the same prefix
- Node uses a table that lists the next hop for IP prefixes

### Longest Matching Prefix

Prefixes might overlap, combines hierarchy with flexibility.

Rule:
- For each packet, find the longest prefix that contains the destination
  address, i.e. the most specific entry
- Forward the packet to the next hop router for that prefix

Example:

```
| Prefix         | Last Address  | Next Hop |
| -------------- | ------------  | -------- |
| 192.24.0.0/18  | 192.24.63.255 | D        |
| 192.24.12.0/22 | 192.24.15.255 | B        | <- Most specific

- 192.24.6.0    -> D
- 192.24.14.32  -> B
- 192.24.54.0   -> D
```

### Host Forwarding Table

- 0.0.0.0/0 is a default route that catches all IP addresses

## ARP and DHCP

### DHCP (Dynamic Host Configuration Protocol)

- It leases IP address to nodes
- Provides other parameters: network prefix, address of local router, etc.
- Uses UDP ports 67, 68

#### Bootstrapping

- Node sends **broadcast** messages that delivered to all nodes on the network
- Broadcast address is all 1s: IPv4 255.255.255.255, MAC ff:ff:ff:ff:ff:ff

1. Client: `DISCOVER` (broadcast)
2. Server: `OFFER`
3. Client: `REQUEST`
4. Server: `ACK`

### ARP (Address Resolution Protocol)

- Node needs link layer addresses to send a frame over the local link
- Node maps a local IP address to its link layer address

```
Ethernet frame:

      Link layer            IP Header
--------------------------------------------------
| Source   | Dest.    | Source | Dest. | Payload |
| Ethernet | Ethernet | IP     | IP    | ...     |
--------------------------------------------------
     ^          ^          ^       ^
    NIC        ARP        DHCP   Target
```

- ARP sits right on top of link layer
- No servers, just asks node with target IP to identify itself
- Uses broadcast to reach all nodes

1. Node: `REQUEST` (broadcast)
2. Target: `REPLY`

## Packet Fragmentation

- Packet size problem: different networks have different maximum packet sizes
(MTU: maximum transmission unit)

- Routers fragment packets that are too large to forward and receiving host
  reassembles to reduce load on routers

- Fragmentation header fields on IP packet (identification, fragment offset,
  MF/DF control bits). MF - more fragments, DF - don't fragment

### IPv4 Fragmentation Procedure

Routers split a packet that is too large:

- Copy ID header to pieces
- Adjust length on pieces
- Set offset to indicate position
- Set MF (More Fragments) on all pieces except last

## Path MTU Discovery

- Discover the MTU that will fit (can avoid fragmentation)
- Host tests path with large packet, routers provide feedback if too large:
  they tell host what size would have fit
- Implemented with ICMP (set DF bit in IP header to get feedback messages)

## Error Handling with ICMP (Internet Control Message Protocol)

- Companion protocol to IP

### ICMP Errors

When router encounters an error while forwarding:
- sends an ICMP error report back to the IP source address

### Message Format

- Type, Code and Checksum
- Often carry the start of the offending packet as payload
- Carried in an IP packet

### Traceroute

IP header contains TTL (time to live) field
- Decremented every router hop, with ICMP error if it hits zero
- Protects against forwarding loops

Traceroute sends probe packets increasing TTL starting from 1
- ICMP errors identify routers on the path

## IPv6

IP header now has source/destination address parts using 16 bytes.

Also:
- no checksum field
- no length field
- flow label (to group packets)

### Transitioning

Approaches:
- dual stack (speak both IPv4 & IPv6)
- translators (convert packets)
- tunnels (carry IPv6 over IPv4)

### Tunneling

Tunnel carries IPv6 packets across IPv4 network.

```
---------------        -----------------------
| IPv6 packet | --|--> | IPv4 | IPv6 packet ||
---------------        -----------------------
```

## NAT (Network Address Translation)

Middleboxes: sit inside the network but perform "more than IP" processing on
packets to add new functionality.
- NAT box, firewall / intrusion detection system

### NAT box

- Connects an internal network to an external network
- Motivated by IP address scarcity
- Keeps an internal/external table
