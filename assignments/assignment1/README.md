# Assignment 1: Socket Programming

As discussed in lecture, socket programming is the standard way to write programs that communicate over a network. While originally developed for Unix computers programmed in C, the socket abstraction is general and not tied to any specific operating system or programming language. This allows programmers to use the socket mental model to write correct network programs in many contexts.

This part of the assignment will give you experience with basic socket programming.  You will write 2 pairs of TCP client and server programs for sending and receiving text messages over the Internet. 

### Server specification
* Each server program should listen on a socket, wait for a client to connect, receive a message from the client, print the message to stdout, and then wait for the next client indefinitely.
* Each server should take one command-line argument: the port number to listen on for client connections.
* Each server should accept and process client communications in an infinite loop, allowing multiple clients to send messages to the same server. The server should only exit in response to an external signal (e.g. SIGINT from pressing `ctrl-c`).
* Each server should maintain a short (5-10) client queue and handle multiple client connection attempts sequentially. In real applications, a TCP server would fork a new process to handle each client connection concurrently, but that is **not necessary** for this assignment.
* Each server should gracefully handle error values potentially returned by socket programming library functions (see specifics for each language below). Errors related to handling client connections should not cause the server to exit after handling the error; all others should.

### Client specification
* Each client program should contact a server, read a message from stdin, send the message, and exit.
* Each client should read and send the message *exactly* as it appears in stdin until reaching an EOF (end-of-file).
* Each client should take two command-line arguments: the IP address of the server and the port number of the server.
* Each client must be able to handle arbitrarily large messages by iteratively reading and sending chunks of the message, rather than reading the whole message into memory first.
* Each client should handle partial sends (when a socket only transmits part of the data given in the last `send` call) by attempting to re-send the rest of the data until it has all been sent.
* Each client should gracefully handle error values potentially returned by socket programming library functions.

### Testing

You should test your implementations by attempting to send messages from your clients to your servers. The server can be run in the background (append a `&` to the command) or in a separate SSH window. You should use `127.0.0.1` as the server IP and a high server port number between 10000 and 60000. You can kill a background server with the command `fg` to bring it to the foreground then `ctrl-c`.

The Bash script `test_client_server.sh` will test your implementation by attempting to send several different messages between all 4 combinations of your clients and servers (C client to C server, C client to Python/Go server, etc.). The messages are the following:

0. The short message "Go Tigers!\n"
0. A long, randomly generated alphanumeric message
0. A long, randomly generated binary message
0. Several short messages sent sequentially from separate clients to one server
0. Several long, random alphaumeric messages sent concurrently from separate clients to one server

Run the script as

`./test_client_server.sh [python|go] [server port]`

If you get a permissions error, run `chmod 744 test_client_server.sh` to give the script execute privileges.

For each client/server pair, the test script will print "SUCCESS" if the message is sent and received correctly. Otherwise it will print a diff of the sent and received message if the diff output is human-readable, i.e., just for tests 1 and 4.

### Debugging hints
Here are some debugging tips. If you are still having trouble, ask a question on Piazza or see an instructor during office hours.

* There are defined buffer size and queue length constants in the scaffolding code. Use them. If they are not defined in a particular file, you don't need them. If you are not using one of them, either you have hard-coded a value, which is bad style, or you are very likely doing something wrong.
* There are multiple ways to read and write from stdin/stdout in C, Python, and Go. Any method is acceptable as long as it does not read an unbounded amount into memory at once and does not modify the message.
* If you are using buffered I/O to write to stdout, make sure to call `flush` or the end of a long message may not write.
* Remember to close the socket at the end of the client program.
* When testing, make sure you are using `127.0.0.1` as the server IP argument to the client and the same server port for both client and server programs.
* If you get "address already in use" errors, make sure you don't already have a server running. Otherwise, restart your ssh session with the command `logout` followed by `vagrant ssh`.  
* If you are getting other connection errors, try a different port between 10000 and 60000.
