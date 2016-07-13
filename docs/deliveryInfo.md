# Delivery Agent (GO)

##### App Operation - Delivery Agent (GO):
1. Continuously pull "postback" objects from Redis
2. Deliver each postback object to http endpoint:
  1. Endpoint method: request.endpoint.method.
  2. Endpoint url: request.endpoint.url, with {xxx} replaced with values from each request.endpoint.data.xxx element.
3. Log delivery time, response code, response time, and response body.

##### Ingestion Agent Info

The delivery agent's log file can be changed by passing in a different argument to the `makeLogger` function.

It logs everything to this one file, both info and errors.