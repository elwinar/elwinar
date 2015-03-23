<?php

Route::get('/', function()
{
	return view('home');
});

Route::get('/login', function()
{
	return view('login');
});

Route::post('/login', function()
{
	if(Auth::attempt([
		'email' => Input::get('email'),
		'password' => Input::get('password'),
	])) {
		return redirect()->intended('/');
	}
	return redirect()->back();
});

Route::get('/logout', function()
{
	Auth::logout();
	return redirect('/');
});

Route::get('/articles', function()
{
	return view('articles.list');
});

Route::get('/article/{article}', function(App\Article $article)
{
	return view('articles.show', [
		'article' => $article,
	]);
});

Route::group(['middleware' => 'auth'], function()
{
	Route::get('/articles/new', function()
	{
		return view('articles.new');
	});
	
	Route::post('/articles/new', function(App\Http\Requests\NewArticleRequest $request)
	{
		$article = App\Article::create($request->all());
		return redirect('/article/'.$article->slug);
	});
});
