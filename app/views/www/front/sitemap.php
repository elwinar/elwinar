<?php
foreach(Route::getRoutes() as $route)
{
	echo $route->getPath().PHP_EOL;
}
?>