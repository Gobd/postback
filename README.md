# Postback Mini-Project

### Kochava Miniproject :: "Postback Delivery"

##### Instructions:
1. Provision provided linux server (see Resources - Server) with software stack required to complete project.
2. Change default redis port immediately on startup
3. Build a php application to ingest http requests, and a go application to deliver http responses. Use Redis to host a job queue between them.
4. Reach out to (Resources - Contact) once your project is ready to demo, or if you encounter a block during development. Pursue independent troubleshooting prior to escalating questions with contact resource.
5. Maintain development notes, provide support documentation, and commit your project (application code / stack config) to Github.

##### Extra Merit:
- [x] Clean, descriptive VCS commit history.
- [x] Clean, easy-to-follow support documentation for an engineer attempting to troubleshoot your system.
- [ ] All services should be configured to run automatically, and service should remain functional after system restarts.
- [ ] High availability infrastructure considerations.
- [ ] Data integrity considerations, including safe shutdown.
- [ ] Modular code design.
- [ ] Configurable default value for unmatched url {key}s.
- [ ] Performance of system under external load.
- [ ] Performance of system with single request in infinite loop.
- [ ] Minimal bandwidth utilization between ingestion and delivery servers.
- [ ] Configurable response delivery retry attempts.
- [x] Ingestion endpoint functional at /i in addition to /ingest.php.
- [x] Data validation / error handling.
- [ ] Ability to deliver POST (as well as GET) responses.
- [ ] Service monitoring / application profiling.
- [ ] Delivery volume / success / failure visualizations.
- [ ] Internal benchmarking tool.

##### Data flow:
1. Web request (see sample request) >
2. "Ingestion Agent" (php) >
3. "Delivery Queue" (redis)
4. "Delivery Agent" (go) >
5. Web response (see sample response)

##### App Operation - Ingestion Agent (PHP):
1. Accept incoming http request
2. Push a "postback" object to Redis for each "data" object contained in accepted request.

##### App Operation - Delivery Agent (GO):
1. Continuously pull "postback" objects from Redis
2. Deliver each postback object to http endpoint:
  1. Endpoint method: request.endpoint.method.
  2. Endpoint url: request.endpoint.url, with {xxx} replaced with values from each request.endpoint.data.xxx element.
3. Log delivery time, response code, response time, and response body.

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

Sample Response:
	`http://sample_domain_endpoint.com/data?key=Phyllobates&value=Terribilis&foo=`