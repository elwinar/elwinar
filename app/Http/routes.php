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
	if(Auth::check()) {
		$articles = App\Article::all();
	} else {
		$articles = App\Article::where('is_published', true)->get();
	}
	return view('articles.list', [
		'articles' => $articles,
	]);
});

Route::get('/article/{article}', function(App\Article $article)
{
	return view('articles.show', [
		'article' => $article,
		'title' => $article->title,
		'keywords' => $article->tags,
		'description' => $article->tagline,
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
	
	Route::get('/article/{article}/publish', function(App\Article $article)
	{
		$article->is_published = true;
		$article->save();
		return redirect()->back();
	});
	
	Route::get('/article/{article}/unpublish', function(App\Article $article)
	{
		$article->is_published = false;
		$article->save();
		return redirect()->back();
	});

	Route::get('/article/{article}/edit', function(App\Article $article)
	{
		return view('articles.edit', [
			'article' => $article,
		]);
	});

	Route::post('/article/{article}/edit', function(App\Http\Requests\EditArticleRequest $request, App\Article $article)
	{
		$article->fill($request->all())->save();
		return redirect('/article/'.$article->slug);
	});
	
	Route::get('/article/{article}/delete', function(App\Article $article)
	{
		$article->delete();
		return redirect()->back();
	});
});
