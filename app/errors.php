<?php

App::error(function(Exception $exception, $code) {
	Log::error($exception);
	if(Config::get('app.debug') == false) {
		return App::make('ErrorController')->callAction('error', array(
			'code' => $code,
			'exception' => $exception,
		));
	}
});

?>