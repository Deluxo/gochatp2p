# Go p2p chat client

Hey! Here you will find my take on chat-client written in golang.

This is the first release of this app, and I'm pretty new to golang,
but already it enabled me to try playing with networking and multithreading!

## How to use

### Parameters

* *-i ip_address:port* The address that will be waiting for input.
* *-o ip_address:port* The address that will be waiting for output.
* *-u your_user_name* The string that will be shown to the ouput user.

### Usage

Both parties must listen and be able to send to i/o addresses provided,
then the clients will be connected.

The TCP error message will be shown if unable to connect.

If you see no messages repeatedly coming from the console, you're ready to chat!
Write your message into the console window!
To end your message, press *ENTER* and *CTRL+D*.

To end the session, simply close the window.

This app is more of a fun home project than something serious, though feel free to fork,
request for pulls or comment.

Have a good day!
