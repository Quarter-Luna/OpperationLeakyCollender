# OpperationLeakyCollender
 UWEmergingTechnologiesLab

This is a client side port scanner that can take individual IPv4 addresses and scan their ports and read back any open ports

There is functionality to automate the reading of ports on a local network using Google DNS 8.8.8.8:80 and it will input them all
into the main go function.  There is also a way to modify the code to only take one type of input into the range of ports to check
with time out sepecified in milliseconds.

This multithreaded process also uses the maxumum allowed threads for a single core to speed up the port checking process.

There can be added functionality to make a list of ports that are restricted to be recorded and sent to a file to be
checked later.