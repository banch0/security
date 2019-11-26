# Adding secure HTTP headers

The following headers are used in this example:

--------------------------------------------------------------
header -----------------------------------------description---
--------------------------------------------------------------
Content-security-Policy --- This defines what scripts or remote
    hosts are trusted and able to provide executable js
--------------------------------------------------------------
X-frame-options --- This defines whether or not frames and iframes
can be used and which domains are allowed to apper in frames
--------------------------------------------------------------
x-xss-Protection --- This tells the browser to stop loading if a 
cross-site-scripting attack is detected; it is largery unnecessary
if a good content-security-policy header is defined
--------------------------------------------------------------
x-content-type-options --- This tells the browser to use MIME 
type provided by the server, and not to modify based upod
guesses by MIME sniffing
--------------------------------------------------------------


server static files 
negroniHandler.Use(negroni.NewStatic(http.Dir("/path/to/static/files")))


# CSRF tokens

Cross-Site Request Forgery, or CSRF, tokens are a way of trying to prevent
one website from taking action on your behalf against a different website.
CSRF is a common attack where a victim will visit a website with malicious
code embedded that tries to make a request to a different site. For example, a
malicious actor embeds JavaScript that makes a POST request to every bank
website attempting to transfer $1,000 to the attacker's bank account. If the victim
has an active session with one of those banks, and the bank does not implement
CSRF tokens, the bank's website may accept the request and process it.

Go provides a xsrftoken package that you can read more about at 
[a link](https://godoc.org/golang.org/x/net/xsrftoken). 
It provides a Generate() function to create tokens and a Valid() function
to validate tokens. You can use their implementation of choose to
develop your own to suit your needs.

To implement CSRF tokens, create a 16-byte random token and store it on the
server associated to the user's session. You can use whatever backend you like to
store the token, whether that is in memory, in a database, or in a file. Embed the
CSRF token in the form as a hidden field. When processing the form on the
server side, verify that the CSRF token is present and matches the user. Destroy
the token after it is used. Do not reuse the same token.
The various requirements for implementing CSRF tokens have been covered in
the previous sections:

* Generating a token: In Chapter 6 , Cryptography, a section titled
Cryptographically secure pseudo-random number generator (CSPRNG)
provides an example of generating random numbers, strings, and bytes.
* Creating, serving, and processing an HTML form: In Chapter 9 , Web
Applications, the section titled HTTP server provides information on
creating a secure web server, and Chapter 12 , Social Engineering, has a
section titled HTTP POST form login honeypot has an example of
processing a POST request.
* Storing a token in a file: In Chapter 3 , Working with Files, the section titled
Write bytes to a file provides an example of storing data in a file.
* Storing a token in a database: In Chapter 8 , Brute Force, the section titled
Brute force database login provides a blueprint for connecting to various
database types.

Preventing user enumeration and
abuse
The important things to remember here are as follows:
Don't let people figure out who has an account
Don't let someone spam your users with your email server
Don't allow people to figure out who is registered by brute force attempts
Let's elaborate on the practical examples.

## Registration

## Preventing LFI and RFI abuse
Local File Inclusion (LFI) and Remote File Inclusion (RFI) are other OWASP
Top 10 vulnerabilities. They refer to the danger of loading files from the local
file system or a remote host that were not intended to be loaded, or loading the
intended files but with contaminated data. Remote file includes are dangerous
because a user may supply a remote file from a malicious server if precaution is
not taken.
Do not open a file from the local file system if the filename is specified by the
user without any sanitization. Consider an example where a file is returned by a
web server upon request. The user may be able to request a file with sensitive
system information, such as /etc/passwd , with a URL like this:

http://localhost/displayFile?filename=/etc/passwd

This could be big trouble if the web server handled it like this (pseudocode):
file = os.Open(request.GET['filename'])
return file.ReadAll()

You can't simply fix it by prepending a specific directory like this:
os.Open('/path/to/mydir/' + GET['filename']).

This isn't enough because attackers can use directory traversal to get back to the
root of the filesystem, as shown here:
http://localhost/displayFile?filename=../../../etc/passwd


## Contaminated files
If an attacker finds an LFI, or you provide a web interface to log files, you need
to make sure that, even if the logs are contaminated, no code will execute.
An attacker can potentially contaminate your logs and insert malicious code by
taking some action on your service that creates a log entry. Any service that
generates a log that is loaded or displayed must be considered.

If an attacker makes a request to
yourwebsite.com/<script>alert("test");</script> , then your HTML log viewer may
actually end up rendering that code, if it is not escaped or sanitized properly.

## Using system proxy
Go's default HTTP client will respect the system's HTTP(S) proxy if set through
environment variables. Go uses the HTTP_PROXY , HTTPS_PROXY and NO_PROXY environment
variables. The lowercase versions are also valid. You can set the environment
variable before running the process or set the environment variable in Go with
this:

os.Setenv("HTTP_PROXY", "proxyIp:proxyPort")

After configuring the environment variable(s), any HTTP request made using the
default Go HTTP client will respect the proxy settings. Read more about the
default proxy settings at https://golang.org/pkg/net/http/#ProxyFromEnvironment .

### Using a SOCKS5 proxy (Tor)
Tor is an anonymity service that attempts to protect your privacy. Do not use Tor
unless you fully understand all of the implications. Read more about Tor at https:
//www.torproject.org . This example demonstrates how to use Tor when making a
request, but this applies equally to other SOCKS5 proxies.

To use a SOCKS5 proxy, the only modification needed is with the URL string of
the proxy. Instead of using the HTTP protocol, use the socks5:// protocol prefix.
The default Tor port is 9050 , or 9150 when using the Tor Browser bundle. The
following example will perform a GET request for check.torproject.org , which will
let you know if you are properly routing through the Tor network.
