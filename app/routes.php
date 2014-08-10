<?php

$host = explode('.', Request::server('HTTP_HOST'));
$domain = (count($host) == 3)?$host[0]:'www';

switch($domain)
{
    case 'www':
		include app_path().'/routes/www.php';
    break;
    
    case 'admin':
		include app_path().'/routes/admin.php';
    break;
}

App::missing(function($exception)
{
	return View::make('errors.layout', array('title' => 'Erreur 404'))->nest('content', 'errors.404');
});

App::fatal(function($exception)
{
	return View::make('errors.layout', array('title' => 'Erreur dans l\'application'))->nest('content', 'errors.fatal', array('exception' => $exception));
});

?>