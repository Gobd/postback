# Index.js

The index.js hosts an Express server listening for the data sent by GO. This is jsut a quick simple test to make sure it's actually sending data. The server runs on port 3000.

POST with method GET created by
```shell
curl -X POST -H "Content-Type: application/json" -d '{ "endpoint": { "method":"GET", "url":"http://localhost:3000/data?key={key}&value={value}&foo={bar}" }, "data":[ { "key":"Azureus", "value":"Dendrobates" }, { "key":"Phyllobates", "value":"Terribilis" } ] }' http://localhost/ingest.php
```
POST with method POST created by
```shell
curl -X POST -H "Content-Type: application/json" -d '{ "endpoint": { "method":"POST", "url":"http://localhost:3000/data" }, "data":[ { "key":"Azureus", "value":"Dendrobates" }, { "key":"Phyllobates", "value":"Terribilis" } ] }' http://localhost/ingest.php
```

These two curl requests will result in GO sending the data to this Express server and having it logged to the console.