<?php
$json = file_get_contents('php://input');
$decoded = json_decode($json, true);
$redis = new Redis();
$connection = $redis->connect('127.0.0.1', 6400);

$headers = json_encode(apache_request_headers());

$date = new DateTime();
$date = $date->format("y:m:d h:i:s");

/*
IF LOGGING ISN'T WORKING THEN CREATE A PHP.LOG AND 'sudo chmod 777 php.log'
That should fix it
*/

if (!$connection) {
	return error_log("{$date} Unable to connect to Redis on port 6400." . PHP_EOL, 3, "/var/www/html/php.log");
}

if (json_last_error()) {
	$decodingError = json_last_error_msg();
	return error_log("{$date} Error deconding: {$decodingError} from {$headers}" . PHP_EOL, 3, "/var/www/html/php.log");
}

if (!isset($decoded['data']) || !isset($decoded['endpoint'])) {
	return error_log("{$date} JSON: {$json} is missing either the 'endpont' or 'data' field from {$headers}" . PHP_EOL, 3, "/var/www/html/php.log");
}

foreach($decoded['data'] as & $data) {
	$postback = array(
		"endpoint" => $decoded['endpoint'],
		"data" => $data,
	);
	$dataToPush = json_encode($postback);
	$dataSuccess = $redis->rPush('requests', $dataToPush);
	if (!$dataSuccess) {
		error_log("{$date} Error pushing data to Redis: {$dataSuccess} from {$headers}" . PHP_EOL, 3, "/var/www/html/php.log");
	}
}

unset($data);
?>