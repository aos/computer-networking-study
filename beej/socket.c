// 5.2 socket()

#include <sys/types.h>
#include <sys/socket.h>

// domain -> what type of socket
// type -> stream or datagram
// protocol -> TCP or UDP
int socket(int domain, int type, int protocol);

// Using socket
int s;
struct addrinfo hints, *res;

// do the lookup
// pretend we already filled out "hints" struct
getaddrinfo("www.example.com", "http", &hints, &res);

// Error check, walk the "res" linked list looking for valid entries
sockfd = socket(res->ai_family, res->ai_socktype, res->ai_protocol);

// socket returns the socket descriptor

// 5.3 bind() -- binds a local port number
int bind(int sockfd, struct sockaddr *my_addr, int addrlen);

// Remove the pesky "Address already in use" error message to
// clear the port
int yes = 1;
if (setsockopt(listener, SOL_SOCKET, SO_REUSEADDR, &yes, sizeof yes) == -1) {
  perror("setsockopt");
  exit(1);
}

// 5.4 connect() -- connecting to a remote host (AS A CLIENT)
// the kernel will choose a local port for us, we only care about
// the remote port
// serv_addr -> destination port and IP address
int connect(int sockfd, struct sockaddr *serv_addr, int addrlen);

// 5.5 listen()
// wait for incoming connections and handle them:
// 1. listen(), 2. accept()
// backlog -> number of connections allowed on the incoming queue
// incoming connections are going to wait in queue until you accept()
int listen(int sockfd, int backlog);

// listening requires calling bind() so the server is running on
// a specific port, the sequence is:
getaddrinfo();
socket();
bind();
listen();
/* accept() goes here */

// 5.6 accept()
// Returns a new socket file descriptor for the accepted connection
// The original one is still listening for more new connections
// New one is used for send() and recv()
int accept(int sockfd, struct sockaddr *addr, socklen_t *addrlen);
