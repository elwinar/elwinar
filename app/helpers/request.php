<?php

function subdomain() {
	if(!isset($_SERVER['HTTP_HOST'])) {
		return '';
	}
	$chunks = explode('.', $_SERVER['HTTP_HOST']);
	return array_shift($chunks);
}

function domain() {
	if(!isset($_SERVER['HTTP_HOST'])) {
		return 'localhost';
	}
	$chunks = explode('.', $_SERVER['HTTP_HOST']);
	array_shift($chunks);
	return implode('.', $chunks);
}

?>