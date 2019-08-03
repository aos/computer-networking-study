#include <string.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netdb.h>
#include <netinet/in.h>

#define MYPORT "3490" // the port users will be connecting to
#define BACKLOG 10    // how many pending connections queue will hold

int main(void)
{
  struct sockaddr_storage their_addr;
  socklen_t addr_size;
  struct addrinfo hints, *res;
  int sockfd, new_fd;

  // don't forget error checking for these calls!

  // Load up address structs with getaddrinfo():
  memset(&hints, 0, sizeof hints);
  hints.ai_family = AF_UNSPEC; // use either IPv4 or IPv6
  hints.ai_socktype = SOCK_STREAM;
  hints.ai_flags = AI_PASSIVE; // fill in IP for me

  getaddrinfo(NULL, MYPORT, &hints, &res);

  // make a socket, bind it, and listen on it:
  sockfd = socket(res->ai_family, res->ai_socktype, res->ai_protocol);
  bind(sockfd, res->ai_addr, res->ai_addrlen);
  listen(sockfd, BACKLOG);

  // now accept incoming connection:
  addr_size = sizeof their_addr;
  new_fd = accept(sockfd, (struct sockaddr *)&their_addr, &addr_size);

  // ready to communicate on socket descriptor new_fd!
  // new_fd will be used for all send() and recv() calls
}
