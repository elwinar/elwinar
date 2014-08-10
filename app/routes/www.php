<?php

Route::bind('article', function($value, $route)
{
	$article = Article::where('slug', $value)->first();
	if($article == null)
	{
		App::abort(404);
	}
	else
	{
		return $article;
	}
});

Route::get('/', '\Www\FrontController@index');
Route::get('home', '\Www\FrontController@home');
Route::get('contact', '\Www\FrontController@contact');
Route::post('contact', '\Www\FrontController@send');
Route::get('sent', '\Www\FrontController@sent');
Route::get('articles', '\Www\FrontController@articles');
Route::get('article/{article}', '\Www\FrontController@article');
Route::get('tools', '\Www\FrontController@tools');
Route::get('sitemap', '\Www\FrontController@sitemap');

Route::get('tools/json-pretty-printer', '\Www\JsonPrettyPrinterController@show');
Route::post('tools/json-pretty-printer', '\Www\JsonPrettyPrinterController@process');

Route::get('tools/slave-password-generator', '\Www\SlavePasswordGeneratorController@show');
Route::post('tools/slave-password-generator', '\Www\SlavePasswordGeneratorController@process');

Route::get('tools/hasher', '\Www\HasherController@show');
Route::post('tools/hasher', '\Www\HasherController@process');

?>