// 5.7 send() and recv()
// Communicating over stream sockets or connected datagram sockets

// send()
int send(int sockfd, const void *msg, int len, int flags);

// sample
char *msg = "Aos was here!";
int len, bytes_sent;

len = strlen(msg);
bytes_sent = send(sockfd, msg, len, 0);

// send() returns the number of bytes actually sent
// this might be less than the number you told it to send

// recv()
int recv(int sockfd, void *buf, int len, int flags);

// buf -> buffer to read information from
// len -> max length of buffer

// 5.8 sendto() and recvfrom()
// unconnected datagram sockets
// Similar to send() and recv()
int sendto(int sockfd, const void *msg, int len, unsigned int flags,
           const struct sockaddr *to, socklen_t tolen);

int recvfrom(int sockfd, void *buf, int len, unsigned int flags,
             struct sockaddr *from, int *fromlen);

// 5.9 close() and shutdown()
// closing a socket:
close(sockfd);

// shutdown allows more control over how the socket closes:
int shutdown(int sockfd, int how);

// how is one of:
// 0 -> further receives disallowed
// 1 -> further sends disallowed
// 2 -> close()

// 5.10 getpeername()
// who is on the other end of a connected stream socket
int getpeername(int sockfd, struct sockaddr *addr, int *addrlen);

// 5.11 gethostname()
// returns the name of the computer that your program is running on
int gethostname(char *hostname, size_t size);
