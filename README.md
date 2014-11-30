# BBB-gif-server
This is a simple BBB server that activates an attached camera and returns a tiny lil' gif.  

## Preparation
- buy a BeagleBone Black
- buy a Logitech Camera
- watch [Derek Molloy's very helpful boneCV video](https://www.youtube.com/watch?v=8QouvYMfmQo)

## Acquire all the needed software
- plug in an ethernet router and plug BBB into your computer via USB
- plug the camera in
- get into the BBB: `ssh 192.168.7.2 -l root`
- `apt-get update`
- `apt-get install v4l-utils libv4l-dev imagemagick ffmpeg gifsicle`
- [install golang](http://golang.org/doc/install)
- git Derek Molloy's [boneCV files](https://github.com/derekmolloy/boneCV.git)
- git this project and copy the files into the same dir as the boneCV

## Turn off some services
- since this project has to run on port 80, it has been recommended to disable some other services:
```
systemctl disable cloud9.service
systemctl disable gateone.service
systemctl disable bonescript.service
systemctl disable bonescript.socket
systemctl disable bonescript-autorun.service
systemctl disable avahi-daemon.service
systemctl disable gdm.service
systemctl disable mpd.service
```
- restart BBB after disabling stuff

## start it up/test it out
- cd into project dir
- `build` to build the binaries used by the _getGif.sh_ script
- `ifconfig` to get the ip address
- start up the server: `go run picServer.go`
- visit it on your browser
- tinker with the script settings

### Todos
- figure out how to make this guy accessable from outside of router using NAT
- include Hubot script to ping this endpoint
