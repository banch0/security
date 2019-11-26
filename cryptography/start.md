Cryptography is the practice of securing communications even when a third-
party can view those communications. There are two-way symmetric and
asymmetric encryption methods, as well as one-way hashing algorithms.

# Symmetric encryption

Symmetric encryption is when the same key or password is used to encrypt and
decrypt the data. Advanced Encryption Standard, also known as AES or
Rijndael, is a symmetric encryption algorithm made standard by NIST in 2001.
Data Encryption Standard, or DES, is another symmetric encryption algorithm
that is older and less secure than AES. It should not be used over AES unless
there is a specific requirement or specification to do so. Go standard library
includes AES and DES packages.

# AES 
This program will encrypt and decrypt a file using a key, which is basically a 
32-byte (256-bit) password.


# Asymmetric encryption

Asymmetric is when there are two keys for each party. A public and private key
pair is required on each side. Asymmetric encryption algorithms include RSA,
DSA, and ECDSA. The Go standard library has packages for RSA, DSA, and
ECDSA. Some applications that use asymmetric encryption include Secure
Shell (SSH), Secure Sockets Layer (SSL), and Pretty Good Privacy (PGP).
SSL is the Secure Sockets Layer originally developed by Netscape, and version
2 was publicly released in 1995. It is used to encrypt communication between a
server and a client providing confidentiality, integrity, and authentication. TLS,
or Transport Layer Security, is the new version of SSL, with 1.2 being defined
in 2008 as RFC 5246. The Go package for TLS does not completely implement
the specification, but it implements the major parts. Read more about Go's
crypto/tls package at https://golang.org/pkg/crypto/tls/ .

## Generating a public and private key pair

Before using asymmetric encryption, you need a public and private key pair. 
The private key must be kept secure and not shared with anyone. The public 
key should be shared with others.

RSA (Rivest-Shamir-Adleman) and ECDSA (Elliptic Curve Digital Signing
Algorithm) algorithms are available in the Go standard library. ECDSA is
considered more secure, but RSA is the most common algorithm used in SSL
certificates.

If you want to password protect your private key file using a symmetric
encryption algorithm, such as AES, you can use some of the standard library
functions. The primary functions you will need are x509.EncryptPEMBlock() ,
x509.DecryptPEMBlock() , and x509.IsEncryptedPEMBlock().

To perform the equivalent operation of generating a private and public key
file using OpenSSL, use the following:

### Generate the private key
$ openssl genrsa -out priv.pem 2048
### Extract the public key from the private key
$ openssl rsa -in priv.pem -pubout -out public.pem

You can learn more about PEM encoding with Go at [a link](https://golang.org/pkg/encoding/pem/)

### Creating a certificate signing request

If you don't want to create a self-signed certificate, you have to create a
certificate signing request and have it signed by a trusted certificate authority.
You create a certificate request by calling x509.CreateCertificateRequest() and
passing it an x509.CertificateRequest object with the private key.

The equivalent operation using OpenSSL is as follows:

#### Create CSR
openssl req -new -key priv.pem -out csr.pem
#### View details to verify request was created properly
openssl req -verify -in csr.pem -text -noout

### OpenPGP

PGP stands for Pretty Good Privacy, and OpenPGP is standard RFC 4880. PGP
is a convenient suite for encrypting text, files, directories, and disks. All the
principles are the same as discussed in the previous section with SSL and TLS
key/certificates. The encrypting, signing, and verification are all the same. Go
provides an OpenPGP package. Read more about it at 
[a link](https://godoc.org/golang.org/x/crypto/openpgp).

### Off The Record (OTR) messaging
Off The Record or OTR messaging is a form of end-to-end encryption for users
to encrypt their communication over whatever message medium is being used. It
is convenient because you can implement an encrypted layer over any protocol
even if the protocol itself is unencrypted. For example, OTR messaging works
over XMPP, IRC, and many other chat protocols. Many chat clients such as
Pidgin, Adium, and Xabber have support for OTR either natively or via plugin.
Go provides a package for implementing OTR messaging. Read more about Go's
OTR support at [a link](https://godoc.org/golang.org/x/crypto/otr/).


