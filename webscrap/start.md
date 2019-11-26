# Web Scraping

Web crawling is another aspect of scraping, which involves following hyperlinks
to other pages. Breadth-first crawling refers to finding as many different
websites as you can and following them to find more sites. Depth-first crawling
refers to crawling a single site to find all pages possible before moving on to the
next site.

HTTP responses contain the body as a reader interface. We can extract the data
from the reader using any function that accepts a reader interface. This includes
functions such as the io.Copy(), io.ReadAtLeast(), io.ReadlAll(), and bufio buffered
readers. In this example, ioutil.ReadAll() is used to quickly store the full contentsof 
the HTTP response into a byte-slice variable.

# Using regular expressions to find email addresses in a page

Regular expressions have a reputation for being slow. The implementation used
is guaranteed to run in linear time, as opposed to exponential time, based on the
input length. This means it will run faster than many other implementations of
regular expressions that do not provide that guarantee, such as Perl. Russ Cox,
one of Go's authors, published a deep comparison of the two different
approaches in 2007, which is available at [a link](https://swtch.com/~rsc/regexp/regexp1.html).
This is very important for our use case of searching the contents of an HTML
page. If the regular expression ran in exponential time, based on the input length,
it could take quite literally years to perform a search of certain expressions.

This example uses a regular expression that searches for email address links
embedded in HTML. It will search for any mailto links and extract the email
address. We'll use the default HTTP client and call http.Get() instead of creating a
custom client to modify the timeout.
A typical email link looks like one of these:

<a href="mailto:nanodano@devdungeon.com">
<a href="mailto:nanodano@devdungeon.com?subject=Hello">

The regular expression used is in this example is this:
"mailto:.*?["?]

# Cookie

Here is a simple example of cookie headers sent from a server:

Set-Cookie: preferred_background=blue
Set-Cookie: session_id=PZRNVYAMDFECHBGDSSRLH

Here is an example header from a client:

Cookie: preferred_background=blue; session_id=PZRNVYAMDFECHBGDSSRLH

# Finding unlisted files on a web server

There is a popular program called DirBuster, which penetration testers use for
finding unlisted files. DirBuster is an OWASP project that comes preinstalled on
Kali, the popular penetration testing Linux distribution. With nothing but the
standard library, we can create a fast, concurrent, and simple clone of DirBuster
with just a few lines. More information about DirBuster is available at 
[a link](https://www.owasp.org/index.php/Category:OWASP_DirBuster_Project).

# Changing the user agent of a request

A common technique to block scrapers and crawlers is to block certain user
agents. Some services will blacklist certain user agents that contain keywords
such as curl and python . You can get around most of these by simply changing
your user agent to firefox.

# Fingerprinting web application technology stacks

I recommend that you inspect the HTTP headers first since they are simple key-
value pairs, and generally, there are only a few returned with each request. It
doesn't take very long to go through the headers manually, so you can inspect
them first before moving on to the application. Fingerprinting at the application
level is more complicated and we'll talk about that in a moment. Earlier in this
chapter, there was a section about extracting HTTP headers and printing them
out for inspection (the Extracting HTTP headers from an HTTP response
section). You can use that program to dump the headers of different web pages
and see what you can find.

The basic idea is simple. Look for keywords. Some headers in particular contain
the most obvious clues, such as the X-Powered-By , Server , and X-Generator headers.
The X-Powered-By header can contain the name of the framework or Content
Management System (CMS) being used, such as WordPress or Drupal.
There are two basic steps to examining the headers. First, you need to get the
headers. Use the example provided earlier in this chapter for extracting HTTP
headers. The second step is to do a string search to look for the keywords. You
can use strings.ToUpper() and strings.Contains() to search directly for keywords, or
use regular expressions. Refer to the earlier examples in this chapter that explain
how to use regular expressions. Once you are able to search through the headers,
you just need to be able to generate the list of keywords to search for.