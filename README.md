# gobgp_clapi_server
This Running REST API GoBGP together.  
You send commands the GoBGP's CLI into the Json's code to the REMOTE GoBGP, Run That's CLI on the Gobgpd!!  
(With the restful auth Token, JWT.)

# gobgp_clapi_lient
This Support tool this HTTP API.
It's very simple 
Now, Only Support BGP ipv4 flowspec.

### supportted client module
- sytax check(ex.address format...)
- token check(if not token is still changed, no password and DO the announce)
- withdraw last announce prefix(opt. --withdraw/-w)
- logging announce prefix

## What you need
- [Golang](https://golang.org/) (You may use later 1.7)
- [Go BGP](https://github.com/osrg/gobgp/releases/latest).
- [Throw Go BGP CLI](https://github.com/osrg/gobgp/blob/master/docs/sources/cli-command-syntax.md)

## Examples(Using shell curl command)

```bash
root@ubu-client:/usr/local# curl -u user:pass -v -H "Content-Type: application/json"  http://localhost:3000/api/token
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 3000 (#0)
* Server auth using Basic with user 'user'
> GET /api/token HTTP/1.1
> Host: localhost:3000
> Authorization: Basic dXNlcjpwYXNz
> User-Agent: curl/7.47.0
> Accept: */*
> Content-Type: application/json
> 
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Mon, 28 Aug 2017 09:40:27 GMT
< Content-Length: 176
< 
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTA0MTcyNDI3LCJuYW1lIjoibnlhIGhva2UifQ.791PWt8-uO2s3Wq_DyjoB3Ju8bIiQZod8MiJzaNitIQ", "expired":"72"}
* Connection #0 to host localhost left intact

root@ubu-client:/go-honban/gobgp_clapi/gobgp_clapi_client# curl -H 'Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTA0MTc0NTM3LCJuYW1lIjoibnlhIGhva2UifQ.TkePeQFBlZUjJwtAIrBuURqlK2fLr3RhhIu5YAPKD5g' -v  POST -d '{"command":"/root/go/bin/gobgp global rib add -a ipv4 10.0.0.1/32 community 100:100 med 10 origin igp local-pref 2000"}' http://localhost:3000/api/command
* Rebuilt URL to: POST/
* Could not resolve host: POST
* Closing connection 0
curl: (6) Could not resolve host: POST
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 3000 (#1)
> POST /api/command HTTP/1.1
> Host: localhost:3000
> User-Agent: curl/7.47.0
> Accept: */*
> Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTA0MTc0NTM3LCJuYW1lIjoibnlhIGhva2UifQ.TkePeQFBlZUjJwtAIrBuURqlK2fLr3RhhIu5YAPKD5g
> Content-Length: 119
> Content-Type: application/x-www-form-urlencoded
> 
* upload completely sent off: 119 out of 119 bytes
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Mon, 28 Aug 2017 10:20:03 GMT
< Content-Length: 0
< 
* Connection #1 to host localhost left intact
```
```bash
root@ubu-gobgpd:~# /root/go/bin/gobgp global rib -a ipv4
   Network              Next Hop             AS_PATH              Age        Attrs
*> 10.0.0.1/32          0.0.0.0                                   00:02:01   [{Origin: i} {Med: 10} {LocalPref: 2000} {Communities: 100:100}]
```
## Examples(Using go_gobgp_client)

```bash
root@ubu-client:/go_gobgp_api/go_gobgp_client# go run main.go 

#########################
  Gobgp Flowspec client
#########################

Do you want to do?(add/del): add
destination_ip(MUST): 192.168.0.1/32
source_ip(MUST): 10.0.0.1/32
protocols(tcp/udp/any): tcp
destion_port: 80
source_port: 53
Do you want to then?(accept/discard/rate-limit <ratelimit>): accept

##########################
    check the hash key
##########################

 100 / 100 [========================================================] 100.00% 1s
 Check is done.

OK,Current HASH key is not still changed.
Go to Next Process.

######################################################################

    Current Hash Code: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwibmFtZSI6IkFkbyBLdWtpYyJ9.qsKN2OIk6AW4O4PMgLjyeBYx0BCG7Iopvei-fNuUivo
	 Post Command: gobgp global rib -a ipv4-flowspec add match destination 192.168.0.1/32 source 10.0.0.1/32 protocol tcp  destination-port =='80'  source-port =='53'  then accept

######################################################################

Do you want to POST this command??(y/n): y

####################
  Working is Done.
####################

```

```bash
root@ubu-bgpd:~# /root/go/bin/gobgp global rib -a ipv4-flowspec
   Network                                                                                                      Next Hop             AS_PATH              Age        Attrs
*> [destination:192.168.0.1/32][source:10.0.0.1/32][protocol:==tcp ][destination-port: ==80][source-port: ==53] fictitious                                00:02:07   [{Origin: ?}]
```
