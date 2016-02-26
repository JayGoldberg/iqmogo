# iqmogo
An IQinVision IQeye IP camera motion detector tool written in Go

The utility uses the motion detection headers provided by IQeye embedded IP cameras. It is using heavily adapted code and ideas from [Paparazzogo](https://github.com/putsi/paparazzogo).

## Modes of operation

1. parsing the streaming method (mulitpart MIME)
1. or, just polls on a regular basis to read the motion header

## What makes it special?

- it uses the in-camera motion detection, **it does not do image analysis**
- has no external dependencies, standard Go libraries
- compiles to single binary
- has low resource usage since it's not doing image analysis

## What sucks about it?

- goroutines not implemented
- doesn't have any sort of resiliency
 - if a thread dies hangs, it will indefinitely try to reconnect, it should be 
- have good performance, contrary to everything I just said ;-)
- probably only works for IQeye, but I'll accept pull requests for this feature on Axis, Mobotix, and Acti cameras, which seem to be the brands most likely to have this feature

## Future
- goroutines for concurrency
- automatically save frames to disk on motion detection
- configuration via `flags` or JSON config file
- better error handling
 - timeout on existing connection, reconnects, bail-out
- detect if a non-IQeye camera was provided and exclude it from the list
