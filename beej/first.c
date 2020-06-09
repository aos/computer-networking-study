#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>

struct addrinfo {
  int             ai_flags;       // AI_PASSIVE, AI_CANONNAME, etc.
  int             ai_family;      // AF_INET, AF_INET6, AF_UNSPEC
  int             ai_socktype;    // SOCK_STREAM, SOCK_DGRAM
  int             ai_protocol;    // use 0 for "any"
  size_t          ai_addrlen;     // size of ai_addr in bytes
  struct sockaddr *ai_addr;       // struct sockaddr_in or _in6
  char            *ai_canonname;  // full canonical hostname

  struct addrinfo *ai_next;       // linked list, next node
};

struct sockaddr {
  unsigned short  sa_family;    // address family, AF_xxx
  char            sa_data[14];  // 14 bytes of protocol address
};

// Can be cast to a pointer of struct sockaddr (and vice-versa)
// IPv4 only
struct sockaddr_in {
  short int           sin_family;   // Address family, AF_INET
  unsigned short int  sin_port;     // Port number (Network Byte Order 'htons')
  struct in_addr      *sin_addr;    // Internet address
  unsigned char       sin_zero[8];  // Same size as struct sockaddr
};

// A whole struct for an int... a bit historical
struct in_addr {
  uint32_t s_addr;  // 32-bit int (4 bytes)
};

// IPv6
struct sockaddr_in6 {
  uint16_t       sin6_family;   // Address family, AF_INET6
  uint16_t       sin6_port;     // Port number, Network Byte Order
  uint32_t       sin6_flowinfo; // IPv6 flow information
  struct in6_addr *sin6_addr;    // IPv6 address
  uint32_t       sin6_scope_id; // Scope ID
};

struct in6_addr {
  unsigned char   s6_addr[16];   // IPv6 address
};

// Large enough to hold both IPv4 and IPv6
struct sockaddr_storage {
  sa_family_t ss_family;  // address family
  
  // all this is padding, implementation specific, ignore it:
  char    __ss_pad1[_SS_PAD1SIZE];
  int64_t __s_align;
  char    __s_pad2[_SS_PAD2SIZE];
};

// Storing string IP addresses:
struct sockaddr_in sa; // IPv4
struct sockaddr_in6 sa6; // IPv6

// inet_pton converts an IP address in string form to a struct in_addr
// "pton" stands for "printable to network"
inet_pton(AF_INET, "10.12.110.57", &(sa.sin_addr)); // IPv4
inet_pton(AF_INET6, "2001:db8:63b3:1::3490", &(sa6.sin6_addr)); // IPv6

// Reversing: struct -> string address
// IPv4
char ip4[INET_ADDRSTRLEN];  // space to hold the IPv4 string

inet_ntop(AF_INET, &(sa.sin_addr), ip4, INET_ADDRSTRLEN);

printf("The IPv4 address is: %s\n", ip4);
