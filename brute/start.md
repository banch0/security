HTTP basic authentication is when you provide a username and password with
your HTTP request. You can pass it as part of the URL in modern browsers.
Consider this example:
http://username:password@www.example.com

When adding basic authentication programmatically, the credentials are provided
as an HTTP header named Authorization , which contains a value of
username:password base64 encoded and prefixed with Basic , separated by a space.
Consider the following example:
Authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ=

## Brute forcing SSH
To protect against attacks like these, implement rate-limiting or a tool such as
fail2ban that locks out accounts for a short duration when a number of failed
login attempts are detected. Also disable the root remote login. Some people like
to put SSH on a non-standard port, but end up putting it on high number non-
restricted ports such as 2222 , which is not a good idea. If you use a high number
non privileged port such as 2222 , another low privilege user could hijack the port
and start running their own service in its place if it ever went down. Put your
SSH daemon on a port lower than 1024 if you want to change it from the default.