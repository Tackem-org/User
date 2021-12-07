# Tackem Master

Master Controller For Tackem System (A Media Server System)

## Install

docker run -d -p 8800:8800 \
--name tackem-master \
--restart=always \
-v /var/run/docker.sock:/var/run/docker.sock \
-v config:/config \
-v logs:/logs \
tackem/tackem-master:latest
## TODO

- split out all cert stuff and keep GPRC as no security for now but use an interface so we can implement the tls options (server and Mutual)
- GPRC Controller and registration stuff (IN PROGRESS)
- config
- logsystem <-- DO THIS TO ALL EXISTING CODE WHERE NEEDED MAINLY DURING BOOT UP AND ALL THE WAY THROUGH SETUP
- Test systems

## Future Ideas to look at later on
- Swarm detection and usage allowing multi machine setups and redundancy of systems (Low Priority)
## Plans

### Setup Plan
1. Generate inital config (In Progress)
1. Checkunderlying system for information. (In Progress)
    1. Check if app has it's network and add if needed. (DONE)
    1. Check if this container is in that network and add if not. (DONE)
    1. Check for any disc drives for the ripper. if found allow them in the setup of systems (ON HOLD)
    1. Check for transcoding hardware. if found then allow hardware transcoding systems (ON HOLD)
1. Install required systems and register them (TODO)
    1. User system. (ON HOLD)
1. Boot up a web server and web socket (DONE)
    1. User setup of admin (ON HOLD)
    1. Ask what libraries you want to add in TV Shows, Movies, Music Etc (ON HOLD)
    1. Show a list of available systems. allow selection of parts and have systems enable its required systems (ON HOLD)
    1. Install selected systems (ON HOLD)
    1. CONFIG SELECTED SYSTEMS IF NEEDED (ON HOLD)
1. REBOOT SYSTEM INTO FULL SYSTEM (ON HOLD)

### Main System Boot Plan (TODO)
1. Load config.
1. Check all underlying system checks to make sure everything is as it should be.
1. Load list of known systems.
    1. Mark all systems as offline in list.
    - How do I store this information? do I want to use SQLite?
1. Start registration server
    - ?? Ping all known systems with server online message ??
    - ?? Wait for all known systems to report they are online ??
1. Start WebServer

## Running mode

## Health Check

## Shutdown

## Used Software

- GOLANG
- Docker
- <https://github.com/spf13/viper>
- <https://github.com/spf13/pflag>

## Links

- <https://pkg.go.dev/github.com/docker/docker>
- <https://registry.hub.docker.com/_/rabbitmq/>
- <https://github.com/asim/go-micro/tree/master/plugins>
- <https://github.com/asim/go-micro/>
- <https://blog.logrocket.com/creating-a-web-server-with-golang/>
- <https://micro.arch.run/docs/go-config.html>

- <https://pkg.go.dev/github.com/spf13/pflag>
- <https://pkg.go.dev/github.com/spf13/viper>
- <https://developers.google.com/protocol-buffers/docs/proto3>
- <http://www.inanzzz.com/index.php/post/7l4u/sending-and-receiving-grpc-client-server-headers-in-golang>
- <http://www.inanzzz.com/index.php/post/gq4x/using-tls-ssl-certificates-for-grpc-client-and-server-communications-in-golang>
- <https://jbrandhorst.com/post/grpc-auth/>
- <https://jbrandhorst.com/post/certify/>
- <https://blog.gopheracademy.com/advent-2019/go-grps-and-tls/>
- <https://bbengfort.github.io/2017/03/secure-grpc/>
- <https://krishicks.com/post/grpc-mutual-tls-golang/>
- <https://github.com/nekonenene/grpc_image/blob/master/hello/src/main.go>
- <https://medium.com/@yeldos/docker-network-example-7aa47ac52285>
- <https://github.com/edbond/Go-gPRC-microservices-Docker-example>
- <https://stackoverflow.com/questions/56870402/gprc-behind-goproxy-returns-certificate-error-works-fine-without-proxy>
- <https://www.digitalocean.com/community/questions/how-to-make-docker-only-use-a-private-network-to-communicate-with-other-hosts>
- <https://docs.docker.com/engine/reference/commandline/network_create/>
- <https://pkg.go.dev/github.com/docker/docker>
- <https://docs.docker.com/engine/api/sdk/examples/>
- <https://docs.docker.com/engine/api/sdk/examples/#run-a-container-in-the-background>
- <https://github.com/nanobox-io/golang-docker-client/blob/62fe52fdab4f61327953080d4eebc7e410ad439a/vendor/github.com/docker/engine-api/types/network/network.go#L51>
- <https://linuxconfig.org/how-to-configure-docker-swarm-with-multiple-docker-nodes-on-ubuntu-18-04>
- <https://blog.raveland.org/post/constraints_swarm/>
- <https://pkg.go.dev/github.com/docker/docker@v20.10.9+incompatible/client#Client.SwarmInit>
- <https://pkg.go.dev/github.com/docker/docker>
- <https://registry.hub.docker.com/_/rabbitmq/>
- <https://github.com/asim/go-micro/tree/master/plugins>
- <https://github.com/asim/go-micro/>
- <https://blog.logrocket.com/creating-a-web-server-with-golang/>
- <https://micro.arch.run/docs/go-config.html>
- <https://pkg.go.dev/github.com/spf13/pflag>
- <https://pkg.go.dev/github.com/spf13/viper>
- <https://pkg.go.dev/github.com/docker/docker>
- <https://registry.hub.docker.com/_/rabbitmq/>
- <https://github.com/asim/go-micro/tree/master/plugins>
- <https://github.com/asim/go-micro/>
- <https://blog.logrocket.com/creating-a-web-server-with-golang/>
- <https://micro.arch.run/docs/go-config.html>
- <https://pkg.go.dev/github.com/spf13/pflag>
- <https://pkg.go.dev/github.com/spf13/viper>
- <https://github.com/grandcat/zeroconf>
- <https://golangexample.com/a-template-to-build-dynamic-web-apps-quickly-using-go/>
- <https://bulma.io/>
- <https://pkg.go.dev/github.com/schollz/websocket>
- <https://developers.google.com/protocol-buffers/docs/proto3>
- <https://www.rodrigoaraujo.me/posts/golang-pattern-graceful-shutdown-of-concurrent-events/>
- <https://stackoverflow.com/questions/39320025/how-to-stop-http-listenandserve>
- <https://pauldigian.hashnode.dev/golang-and-context-an-explanation>
