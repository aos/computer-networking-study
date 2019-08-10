## Web Proxy

In this assignment, you will implement a simple web proxy that passes requests
and data between a web client and a web server. This will give you a chance to
get to know one of the most popular application protocols on the Internet- the
Hypertext Transfer Protocol (HTTP)v. 1.0- and give you an introduction to the
Berkeley sockets API. When you're done with the assignment, you should be able
to configure your web browser to use your personal proxy server as a web proxy.

## Introduction: The Hypertext Transfer Protocol

The Hypertext Transfer Protocol or (HTTP) is the protocol used for
communication on this web. That is, it is the protocol which defines how your
web browser requests resources from a web server and how the server responds.
For simplicity, in this assignment we will be dealing only with version 1.0 of
the HTTP protocol, defined in detail in [RFC 1945](http://web.archive.org/web/20080704073753/http://www.ietf.org/rfc/rfc1945.txt).
You should read through this RFC and refer back to it when deciding on the
behavior of your proxy.

HTTP communications happen in the form of transactions, a transaction consists
of a client sending a request to a server and then reading the response.
Request and response messages share a common basic format:

* An initial line (a request or response line, as defined below)
* Zero or more header lines
* A blank line (CRLF)
* An optional message body.

For most common HTTP transactions, the protocol boils down to a relatively
simple series of steps (important sections of [RFC 1945](http://web.archive.org/web/20080704073753/http://www.ietf.org/rfc/rfc1945.txt) are in parenthesis):

1. A client creates a connection to the server.
2. The client issues a request by sending a line of text to the server. This
   **request line** consists of a HTTP _method_ (most often GET, but others
   are defined in the RFCs), a _request URI_ (like a URL), and the protocol
   version that the client wants to use (HTTP/1.0). The message body of the
   initial request is typically empty. (5.1-5.2, 8.1-8.3, 10, D.1)
3. The server sends a response message, with its initial line consisting of a
   **status line**, indicating if the request was successful. The status line
   consists of the HTTP version (HTTP/1.0), a _response status code_ (a
   numerical value that indicates whether or not the request was completed
   successfully), and a _reason phrase_, an English-language message providing
   description of the status code. Just as with the the request message, there
   can be as many or as few header fields in the response as the server wants
   to return. Following the CRLF field separator, the message body contains
   the data requested by the client in the event of a successful request.
   (6.1-6.2, 9.1-9.5, 10)
4. Once the server has returned the response to the client, it closes the
   connection.

It's fairly easy to see this process in action without using a web browser.
From a Unix prompt, type:

`telnet www.yahoo.com 80`

This opens a TCP connection to the server at www.yahoo.com listening on port
80- the default HTTP port. You should see something like this:

```
Trying 209.131.36.158...
Connected to www.yahoo.com (209.131.36.158).
Escape character is '^]'.
```

type the following:

`GET / HTTP/1.0`

and hit enter twice. You should see something like the following:

```
HTTP/1.1 200 OK
Date: Fri, 10 Nov 2006 20:31:19 GMT
Connection: close
Content-Type: text/html; charset=utf-8

<html><head>
<title>Yahoo!</title>
(More HTML follows)
```


There may be some additional pieces of header information as well- setting
cookies, instructions to the browser or proxy on caching behavior, etc. What
you are seeing is exactly what your web browser sees when it goes to the Yahoo
home page: the HTTP status line, the header fields, and finally the HTTP
message body- consisting of the HTML that your browser interprets to create a
web page.


### HTTP Proxies

Ordinarily, HTTP is a client-server protocol. The client (usually your web
browser) communicates directly with the server (the web server software).
However, in some circumstances it may be useful to introduce an intermediate
entity called a proxy. Conceptually, the proxy sits between the client and the
server. In the simplest case, instead of sending requests directly to the
server the client sends all its requests to the proxy. The proxy then opens a
connection to the server, and passes on the client's request. The proxy
receives the reply from the server, and then sends that reply back to the
client. Notice that the proxy is essentially acting like both a HTTP client (to
the remote server) and a HTTP server (to the initial client).

Why use a proxy? There are a few possible reasons:

1. **Performance:** By saving a copy of the pages that it fetches, a proxy
   can reduce the need to create connections to remote servers. This can
   reduce the overall delay involved in retrieving a page, particularly if a
   server is remote or under heavy load.
2. **Content Filtering and Transformation:** While in the simplest case the
   proxy merely fetches a resource without inspecting it, there is nothing
   that says that a proxy is limited to blindly fetching and serving files.
   The proxy can inspect the requested URL and selectively block access to
   certain domains, reformat web pages (for instances, by stripping out
   images to make a page easier to display on a handheld or other
   limited-resource client), or perform other transformations and filtering.
3. **Privacy:** Normally, web servers log all incoming requests for
   resources. This information typically includes at least the IP address of
   the client, the browser or other client program that they are using
   (called the User-Agent), the date and time, and the requested file. If a
   client does not wish to have this personally identifiable information
   recorded, routing HTTP requests through a proxy is one solution. All
   requests coming from clients using the same proxy appear to come from the
   IP address and User-Agent of the proxy itself, rather than the individual
   clients. If a number of clients use the same proxy (say, an entire
   business or university), it becomes much harder to link a particular HTTP
   transaction to a single computer or individual.

**Links:**

* [RFC 1945](http://web.archive.org/web/20080704073753/http://www.w3.org/Protocols/rfc1945/rfc1945)
The Hypertext Transfer Protocol, version 1.0


### The Basics

Your first task is to build a basic web proxy capable of accepting HTTP
requests, making requests from remote servers, and returning data to a client.

You shouldn't assume that your server will be running on a particular IP
address, or that clients will be coming from a pre-determined IP.


### Listening

When your proxy starts, the first thing that it will need to do is establish a
socket connection that it can use to listen for incoming connections. Your
proxy should listen on the port specified from the command line, and wait for
incoming client connections.

Once a client has connected, the proxy should read data from the client and
then check for a properly-formatted HTTP request. An invalid request from the
client should be answered with an appropriate error code.


### Parsing the URL

Once the proxy sees a valid HTTP request, it will need to parse the requested
URL. The proxy needs at most three pieces of information: the requested host
and port, and the requested path. See the `URL (7)` manual page for more info.


### Getting Data from the Remote Server

Once the proxy has parsed the URL, it can make a connection to the requested
host (using the appropriate remote port, or the default of 80 if none is
specified) and send a HTTP request for the appropriate file. The proxy then
sends the HTTP request that it received from the client to the remote server.


### Returning Data to the Client

After the response from the remote server is received, the proxy should send
the response message to the client via the appropriate socket. Once the
transaction is complete, the proxy should close the connection.


### Performance and Design Considerations

* You must not use a hard-coded port number.
* You must be able to handle multiple simultaneous connections. You can
  implement this feature however you like- forking, threading, etc.
* Your proxy should not statically allocate large amounts of memory for use
  in reading requests or responses. Instead, allocate a reasonably sized
  buffer and then grow it to some reasonable maximum.
* You are only required to implement HTTP 1.0 methods. Additional methods
  specified in the HTTP 1.1 RFC should be responded to with a 501 error..
* You must support the basic HTTP 1.0 methods- GET, HEAD, and POST.
  Additional methods specified in the RFC can be implemented as an additional
  feature.
* You do not need to implement checking of header fields. As long as they
  follow the basic format specified in the RFC (section 4.2), they should be
  passed through unmodified.


### Testing Your Proxy

Run your client with the following command:

`./proxy <port>`, where `port` is the port number that the proxy should listen
on. As a basic test of functionality, try requesting a page using telnet:

```
telnet localhost <port>
Trying 127.0.0.1...
Connected to localhost.localdomain (127.0.0.1).
Escape character is '^]'.
GET http://www.google.com HTTP/1.0

```

If your proxy is working correctly, the headers and HTML of the Google homepage
should be displayed on your terminal screen.

For a slightly more complex test, you can configure your web browser to use
your proxy server as its web proxy. See the section below for details.


## Configuring a Web Browser to Use a Proxy


### A Caveat

**Using your proxy with a web browser will not work 100% correctly until you
have multi-threading/forking working.** Because a web browser like Firefox or
IE issues multiple HTTP requests for each URL you request (for instance, to
download images and other embedded content), a single-threaded proxy will
likely miss some requests, resulting in missing images or other minor errors.
That's OK for testing, but you must correctly handle multiple connections in
your final submission.


### Firefox

**Version 2.0:**

1. Select Tools->Options from the menu.
2. Click on the 'Advanced' icon in the Options dialog.
3. Select the 'Network' tab, and click on 'Settings' in the 'Connections'
   area.
4. Select 'Manual Proxy Configuration' from the options available. In the
   boxes, enter the hostname and port where proxy program is running.

To stop using the proxy server, select 'Direct connection to the Internet' in
the connection settings dialog.


#### Configuring Firefox to use HTTP/1.0

Because Firefox defaults to using HTTP/1.1 and your proxy speaks HTTP/1.0,
there are a couple of minor changes that need to be made to Firefox's
configuration. Fortunately, Firefox is smart enough to know when it is
connecting through a proxy, and has a few special configuration keys that can
be used to tweak the browser's behavior.

1.  Type 'about:config' in the title bar.
2.  In the search/filter bar, type 'network.http.proxy'
3.  You should see three keys: `network.http.proxy.keepalive`,
    `network.http.proxy.pipelining`, and `network.http.proxy.version`.
4.  Set `keepalive` to false. Set `version` to 1.0\. Make sure that
    `pipelining` is set to false.


## Socket Programming

In order to build your proxy you will need to learn and become comfortable
programming sockets. The Berkeley sockets library is the standard method of
creating network systems on Unix.

You can find the details of these functions in the Unix `man` pages (most of
them are in section 2) and in the Stevens _Unix Network Programming_ book,
particularly chapters 3 and 4\. Other sections you may want to browse include
the client-server example system in chapter 5 (you will need to write both
client and server code for this assignment) and the name and address conversion
functions in chapter 9.

**Links:**

* [Guide to Network Programming Using Sockets](http://web.archive.org/web/20080704073753/http://www.beej.us/guide/bgnet/)
* [An Introduction to Sockets Programming](http://web.archive.org/web/20080704073753/http://www.uwo.ca/its/doc/courses/notes/socket/)
* [HTTP Made Really Easy- A Practical Guide to Writing Clients and Servers](http://web.archive.org/web/20080704073753/http://www.jmarshall.com/easy/http/)


### A Note on Network Programming

Writing code that will interact with other programs on the Internet is a little
different than just writing something for your own use. The general guideline
often given for network programs is: **be lenient about what you accept, but
strict about what you send**. That is, even if a client doesn't do exactly the
right thing, you should make a best effort to process their request if it is
possible to easily figure out their intent. On the other hand, you should
ensure that anything that you send out conforms to the published protocols as
closely as possible. If an incoming request has a single field out of whack
(such as sending you a request using HTTP 0.9 or 1.1), uses non-standard line
terminators (some clients only send \r instead of the standard \r\n), or does
something you don't quite expect with HTTP headers, you should still handle the
request rather than dropping the request. Pay attention to parts of the RFC
that specify areas where not all clients may conform exactly to what you
expect. We'll be looking for this kind of interoperability in both the second
round of tests that we run and in the style portion of your grade.

When in doubt, try to follow the behavior specified in
[RFC 1945](http://www.ietf.org/rfc/rfc1945.txt).


### Possible Extensions

While it may not be obvious at first, proxies are very flexible tools that can
serve a number of different purposes on the web. Common uses for proxies
include improving giving performance boosts to dial-up users (through caching
and pre-fetching), privacy protection (through anonymous proxies), content
filtering and blocking (used in many "NetNanny"-type applications), and content
transformation.

**Sample Proxy Applications:**

* [Anonymizer](http://web.archive.org/web/20080704073753/http://www.anonymizer.com/)
    - A privacy protection/anonymous browsing service.
* [Foxy](http://web.archive.org/web/20080704073753/http://www.2-power-n.com/)
    - A filtering web proxy.
* [Google Web Accelerator](http://web.archive.org/web/20080704073753/http://webaccelerator.google.com/)
    - The latest of a number of 'accelerators'.

#### Content Transformation

Content transformation is the process of a proxy inserting, removing, or
changing the contents of a resource requested from a remote server. After the
resource has been retrieved from the server, the proxy is free to do whatever
it would like to the content. Since the data returned from a web server is
usually just text, this means that we can change the page almost any way we
want- add or remove dirty words, change the text to Pig-Latin, rotate the
images on the page 90 degrees, etc.


#### Caching

Caching is one of the most common performance enhancements that web proxies
implement. Caching takes advantage of the fact that most pages on the web don't
change that often, and that any page that you visit once you (or someone else
using the same proxy) are likely to visit again. A caching proxy server saves a
copy of the files that it retrieves from remote servers. When another request
comes in for the same resource, it returns the saved (or _cached_) copy instead
of creating a new connection to a remote server. This saves a modest amount of
time and CPU if the remote server is nearby and lightly trafficked, but can
create more significant savings in the case of a more distant server or a
remote server that is overloaded (it can also help reduce the load on heavily
trafficked servers).

Caching introduces a few new complexities as well. First of all, a great deal
of web content is dynamically generated, and as such shouldn't really be
cached. Second, we need to decide how long to keep pages around in our cache.
If the timeout is set too short, we negate most of the advantages of having a
caching proxy. If the timeout is set too long, the client may end up looking at
pages that are outdated or irrelevant.

There are a few steps to implementing caching behavior for your web proxy:

1. First, alter your proxy so that you can specify a timeout value (probably
   in seconds) on the command line.
2. Second, you'll need to alter how your proxy retrieves pages. It should now
   check to see if a page exists in the proxy before retrieving a page from a
   remote server. If there is a valid cached copy of the page, that should be
   presented to the client instead of creating a new server connection.
3. Finally, you will need to somehow implement cache expiration. The timing
   does not need to be exact (i.e. it's okay if a page is still **in** your
   cache after the timeout has expired, but it's not okay to **serve** a
   cached page after its timeout has expired), but you want to ensure that
   pages that are older than the user-set timeout are not served from the
   cache.


#### Link Prefetch

Building on top of your caching and content transformation code, the last piece
of functionality that you will implement is called link prefetching. The idea
behind link prefetching is simple: if a user asks for a particular page, the
odds are high that he or she will next request a page linked from that page.
Link prefetching uses this information to attempt to speed up browsing by
parsing requested pages for links, and then fetching the linked pages in the
background. The pages fetched from the links are stored in the cache, ready to
be served to the client when they are requested without the client having to
wait around for the remote server to be contacted.

Parsing and fetching links can take an appreciable amount of time, especially
for a page with a lot of links. For this reason, if you haven't already, at
this stage you should make your proxy into a multi-threaded application. One
thread should remain dedicated to the tasks that you have already implemented:
reading requests from the client and serving pages from either the cache or a
remote server. In a separate thread, the proxy will parse a page and extract
the HTTP links, request those links from the remote server, and add them to the
cache.


#### Other Possible Extensions

* HTTP 1.1 Support
* HTTP Connection Keep-Alive
