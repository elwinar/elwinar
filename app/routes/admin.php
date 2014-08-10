<?php

Route::get('/', '\Admin\FrontController@index');
Route::get('home', '\Admin\FrontController@home');
Route::get('login', '\Admin\FrontController@signin');
Route::post('login', '\Admin\FrontController@login');
Route::get('logout', '\Admin\FrontController@logout');

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

Route::get('articles', '\Admin\ArticleController@index');
Route::get('articles/new', '\Admin\ArticleController@create');
Route::post('articles/new', '\Admin\ArticleController@store');
Route::get('article/{article}', '\Admin\ArticleController@show');
Route::get('article/{article}/edit', '\Admin\ArticleController@edit');
Route::post('article/{article}/edit', '\Admin\ArticleController@update');
Route::get('article/{article}/delete', '\Admin\ArticleController@destroy');

?>