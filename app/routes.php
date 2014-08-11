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

?>