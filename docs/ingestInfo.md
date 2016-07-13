# Ingestion Agent (PHP)

1. Accept incoming http request
2. Push a "postback" object to Redis for each "data" object contained in accepted request.

##### Sample Request

POST

`http://{server_ip}/ingest.php`

POST Data
```json
    {
  	  "endpoint":{
  	    "method":"GET",
  	    "url":"http://sample_domain_endpoint.com/data?key={key}&value={value}&foo={bar}"
  	  },
  	  "data":[
  	    {
  	      "key":"Azureus",
  	      "value":"Dendrobates"
  	    },
  	    {
  	      "key":"Phyllobates",
  	      "value":"Terribilis"
  	    }
  	  ]
  	}
```

POST with method GET created by
```shell
curl -X POST -H "Content-Type: application/json" -d '{ "endpoint": { "method":"GET", "url":"http://localhost:3000/data?key={key}&value={value}&foo={bar}" }, "data":[ { "key":"Azureus", "value":"Dendrobates" }, { "key":"Phyllobates", "value":"Terribilis" } ] }' http://localhost/ingest.php
```
POST with method POST created by
```shell
curl -X POST -H "Content-Type: application/json" -d '{ "endpoint": { "method":"POST", "url":"http://localhost:3000/data" }, "data":[ { "key":"Azureus", "value":"Dendrobates" }, { "key":"Phyllobates", "value":"Terribilis" } ] }' http://localhost/ingest.php
```
##### Ingestion Agent Info

The ingestion agent tries to log all errors to `/var/www/html/php.log` including the time, request headers, the data, and some kind of helpful error message.