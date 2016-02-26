# IQeye motion data examples

## Multipart MIME stream

```
$ telnet 192.168.1.83 80
Trying 192.168.1.83...
Connected to 192.168.1.83.
Escape character is '^]'.
GET /now.jpg?snap=spush0.1&pragma=motion&noimage HTTP/1.0

HTTP/1.0 200 OK
Cache-Control: no-cache
Content-Type: multipart/x-mixed-replace; boundary=--ImageSeparator

--ImageSeparator
Content-Type: image/jpeg
Content-Length: 0
Pragma: trigger=motion 1

--ImageSeparator
Content-Type: image/jpeg
Content-Length: 0
Pragma: trigger=motion 1 2
```

## Single request

```
$ telnet 192.168.1.83 80
Trying 192.168.1.83...
Connected to 192.168.1.83.
Escape character is '^]'.
GET /now.jpg?pragma=motion&noimage HTTP/1.0

HTTP/1.0 200 OK
Server: IQinVision Embedded 1.0
Content-Type: image/jpeg
Content-Length: 0
Pragma: trigger=none
Date: Fri, 26 Feb 2016 16:06:09 GMT
Last-Modified: Fri, 26 Feb 2016 16:06:09 GMT
Expires: -1
Pragma: no-cache
Cache-Control: no-cache
```
