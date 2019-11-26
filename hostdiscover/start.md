# Host Discovery and Enumeration

Host discovery is the process of looking for hosts on a network. This is useful if
you gained access to a machine on a private network, and you want to see which
other machines are on the network and start to gather a picture of what the
network looks like.

A common tool used for
this purpose is nmap. In this chapter, we will cover basic port scanning with a
TCP connect scan and banner grabbing, which are two of the most common use
cases for nmap. We will also cover raw socket connections that can be used to
manually interact and explore a server's ports.

Enumeration is a similar idea, but refers to actively examining a specific
machine to determine as much information as you can. This includes scanning a
server's ports to see which one is open, grabbing banners to inspect services,
making calls to various services to get version numbers and generally search for
attack vectors.

Fuzzing is also covered in this chapter, although it is only touched on very
briefly. Fuzzing warrants its own chapter and, in fact, whole books have been
written about the topic. Fuzzing is more useful when reverse engineering or
searching for vulnerabilities, but can be useful for getting information about a
service. For example, a service may return no response, giving you no clues
about its purpose, but if you fuzz it with bad data and it returns an error, you may
learn what kind of input it is expecting.

# Port scanning
After finding a host on the network, perhaps after doing a ping sweep or
monitoring the network traffic, you typically want to scan the ports and see
which ports are open and accepting connections. You can learn a lot about a
machine just by seeing what ports are open. You might be able to determine
whether it is Windows or Linux or whether it is hosting an email server, a web
server, a database server, and more.

# Fuzzing a network service

Fuzzing is when you send intentionally malformed, excessive, or random data to
an application in an attempt to make it misbehave, crash, or reveal sensitive
information. You can identify buffer overflow vulnerabilities, which can result in
remote code execution. If you cause an application to crash or stop responding
after you send it data of a certain size, it may be due to a buffer overflow.

Sometimes, you will just cause a denial of service by causing a service to use too
much memory or tie up all the processing power. Regular expressions are
notoriously slow and can be abused in the URL routing mechanisms of web
applications to consume all the CPU with few requests.

# Summary
There are numerous other forms of enumeration. With web applications, you can
enumerate usernames, user ids, emails, and more. For example, if a website used
the URL format www.example.com/user_profile/1234 you can potentially start with the
number 1, and increment by 1, crawling through every single user profile
available on the site. Other forms include SNMP, DNS, LDAP, and SMB.
What other forms of enumeration can you think of? What kind of enumeration
can you think of if you were already on a server with a low privilege user? What
kind of information would you want to gather about a server once you had a
shell?