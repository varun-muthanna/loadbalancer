wrk.headers["Host"] = "domain2.com"

-- wrk -t4 -c5 -d2s -s host.lua http://localhost:8080 
-- t (threads) 
-- c (number of tcp connections)
-- d (duration)  send as many requests as possible 
-- concurrent requests is bounded by c not t , t is how many OS threads are used to drive those connections.