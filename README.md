# go_gobgp_api
This Running REST API over GoBGP.
Curl with GoBGP's CLI this HTTP API, REMOTE GoBGP Run That's Command.
(On the Restful Auth is JWT.)

# Examples(Using shell curl command)

root@ubu-client:~# curl -u user:pass -v  http://localhost:3000/api/token*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 3000 (#0)
* Server auth using Basic with user 'user'
> GET /api/token HTTP/1.1
> Host: localhost:3000
> Authorization: Basic dXNlcjpwYXNz
> User-Agent: curl/7.47.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Sun, 27 Aug 2017 10:14:26 GMT
< Content-Length: 154
< Content-Type: text/plain; charset=utf-8
< 
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwibmFtZSI6IkFkbyBLdWtpYyJ9.qsKN2OIk6AW4O4PMgLjyeBYx0BCG7Iopvei-fNuUivo", "expired":"24"}
* Connection #0 to host localhost left intact

root@ubu-client:~# curl -H 'Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwibmFtZSI6IkFkbyBLdWtpYyJ9.qsKN2OIk6AW4O4PMgLjyeBYx0BCG7Iopvei-fNuUivo' -v  -X POST -d '{"command":"/root/go/bin/gobgp global rib add -a ipv4 10.0.0.1/32 community 100:100 med 10 origin igp local-pref 2000"}' http://localhost:3000/api/command
Note: Unnecessary use of -X or --request, POST is already inferred.
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 3000 (#0)
> POST /api/command HTTP/1.1
> Host: localhost:3000
> User-Agent: curl/7.47.0
> Accept: */*
> Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwibmFtZ* upload completely sent off: 119 out of 119 bytes
< HTTP/1.1 200 OK
< Date: Sun, 27 Aug 2017 10:19:11 GMT
< Content-Length: 0
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host localhost left intact
SI6IkFkbyBLdWtpYyJ9.qsKN2OIk6AW4O4PMgLjyeBYx0BCG7Iopvei-fNuUivo
> Content-Length: 119
> Content-Type: application/x-www-form-urlencoded

root@ubu-gobgpd:~# /root/go/bin/gobgp global rib -a ipv4
   Network              Next Hop             AS_PATH              Age        Attrs
*> 10.0.0.1/32          0.0.0.0                                   00:02:01   [{Origin: i} {Med: 10} {LocalPref: 2000} {Communities: 100:100}]

# Examples(Using go_gobgp_client)

root@ubu-bgp:/go-honban/go_gobgp_api/go_gobgp_client# go run main.go 

#########################
  Gobgp Flowspec client
#########################

Do you want to do?(add/del): add

