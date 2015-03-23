<?php

Route::get('/', function() {
	return view('home');
});

Route::get('/login', function() {
	return view('login');
});

Route::post('/login', function() {
	if(Auth::attempt([
		'email' => Input::get('email'),
		'password' => Input::get('password'),
	])) {
		return redirect()->intended('/');
	}
	return redirect()->back();
});

Route::get('/logout', function() {
	Auth::logout();
	return redirect('/');
});
