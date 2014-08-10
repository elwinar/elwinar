<?php namespace Admin;

use Auth;
use BaseController;
use Input;
use Redirect;

class FrontController extends BaseController {

    protected $layout = 'admin.layout';
    
    public function __construct()
    {
        $this->beforeFilter('auth', array(
			'except' => array(
				'login',
				'signin'
			)
        ));
    }
    
    public function index()
    {
		$this->home();
    }
    
    public function home()
    {
        $this->display('admin.front.home', array('title' => '\o'));
    }

	public function signin()
	{
        $this->display('admin.front.login', array('title' => 'Connexion'));
	}
    
    public function login()
    {
		if (Auth::attempt(array('username' => Input::get('username'), 'password' => Input::get('password')), true))
		{
			return Redirect::intended('/');
		}
		else
		{
			return Redirect::to('login');
		}
    }
    
    public function logout()
    {
		Auth::logout();
		return Redirect::to('/');
    }

}