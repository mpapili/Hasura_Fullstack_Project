#! /usr/bin/bash

for i in {1..1000}; do
	curl http://localhost:8000/meetups
	#curl http://localhost:8081/meetups
	#curl 'http://localhost:8080/v1/graphql' \
  		#-H 'Connection: keep-alive' \
  		#-H 'sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="98"' \
  		#-H 'sec-ch-ua-mobile: ?0' \
  		#-H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.80 Safari/537.36' \
  		#-H 'sec-ch-ua-platform: "Linux"' \
  		#-H 'content-type: application/json' \
  		#-H 'Accept: */*' \
  		#-H 'Origin: http://localhost:3000' \
  		#-H 'Sec-Fetch-Site: same-site' \
  		#-H 'Sec-Fetch-Mode: cors' \
  		#-H 'Sec-Fetch-Dest: empty' \
  		#-H 'Referer: http://localhost:3000/' \
  		#-H 'Accept-Language: en-US,en;q=0.9' \
  		#--data-raw '{"query":"\n            {\n                meetups_meetup {\n                    address\n                    description\n                    id\n                    image\n                    title\n                }    \n            }\n        "}' \
  		#--compressed
done
