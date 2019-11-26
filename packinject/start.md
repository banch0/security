sudo apt-get install libpcap-dev
go get github.com/google/gopacket

=========================================================
The equivalent tcpdump command to list devices is as follows:

tcpdump -D

Alternatively, you can use this command:
tcpdump --list-interfaces

==========================================================
The dst net filter will capture incoming traffic
that is going to a 192.168.0.* address. The port filter is watching only for port 22
traffic. The not broadcast and not multicast filter demonstrates how you can negate
and combine multiple filters. Filtering out broadcast and multicast is useful because
they tend to clutter a capture.


The equivalent tcpdump command for a basic capture is simply running it and
passing it an interface:
tcpdump -i eth0

If you want to pass filters, you just pass them as command-line arguments, like
this:
tcpdump -i eth0 tcp port 80
==========================================================
To do the equivalent with tcpdump, just pass it the -w flag with a filename, 
as shown in the following command:
tcpdump -i eth0 -w my_capture.pcap
===========================================================
After getting a handle, whether it was from pcap.OpenLive() or
pcap.OpenOffline() , the handle is treated the same. No distinction is made between
a live device and a capture file once the handle is created, except that a live
device will continue to deliver packets, and a file will eventually end.
You can read pcap files that were captured with any libpcap client, including
Wireshark, tcpdump , or other gopacket applications. This example opens a file
named test.pcap using pcap.OpenOffline() and then iterates through the packets
using range and prints the basic packet information.
============================================================
