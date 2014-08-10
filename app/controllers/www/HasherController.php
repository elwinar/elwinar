<?php namespace Www;

class HasherController extends \BaseController {

    protected $layout = 'www.layout';
	
	public function show()
	{
		$this->display('www.tools.hasher', array(
			'title' => 'Hasher',
			'description' => 'Formulaire de hashage de chaines',
		));
	}
	
	public function process()
	{
		\Session::flash('result', hash(\Input::get('algorithm'), \Input::get('data')));
		\Input::flash();
		return \Redirect::to(\URL::previous());
	}
};

?>