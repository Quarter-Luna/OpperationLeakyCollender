# OpperationLeakyCollender
 UWEmergingTechnologiesLab

This is a client side port scanner that can take individual IPv4 addresses and scan their ports and read back any open ports

There is functionality to automate the reading of ports on a local network using Google DNS 8.8.8.8:80 and it will input them all
into the main go function.  There is also a way to modify the code to only take one type of input into the range of ports to check
with time out sepecified in milliseconds.

This multithreaded process also uses the maxumum allowed threads for a single core to speed up the port checking process.

There can be added functionality to make a list of ports that are restricted to be recorded and sent to a file to be
checked later.

Brief write up:

This local network port scanner was designed to iterate over a local network finding open ports that shouldn’t be open and any potential un-authorized connections between computers and the internet.  The main challenge I foresaw with this program was keeping the network scanner from cascading outside of the local network, thankfully after some research using go to bounce off of the Google DNS 8.8.8.8:80 port allowed it to simply go back to each computer on the local network.  The IO inputs were simple thanks to previous C++ and C programming, just under some different syntax to allow terminal inputs.  This could be expanded out to be done in a window and could even be used to expand out into a port restriction file.  Most of the additional functionality is for more in-depth logging and for this program to be run more than once a month.  The port scanner was taken from Chapter 2 of Black Hat Go by Tom Steele, Chris Patten, and Dan kottmann, and modified to be done with multithreading using semaphore which to my surprise works on ARM64 architecture.  The multithreading was actually very simple thanks to documentation available but thanks to GO’s C integration it allowed for better threading with simple protection to prevent problems.  This program still needs to be used by an actual computer professional but with its current state, it would be useful for Small businesses in order to make sure their networks aren’t making connections without their permission or any slackers playing Call Of Duty.