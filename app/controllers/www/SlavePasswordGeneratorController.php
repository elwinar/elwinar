<?php namespace Www;

class SlavePasswordGeneratorController extends \BaseController {

    protected $layout = 'www.layout';
	
	public function show()
	{
		$this->display('www.tools.slave-password-generator', array(
			'title' => 'Générateur de mots de passe',
			'description' => 'Un générateur de mots de passe intelligent',
		));
	}
	
	public function process()
	{
		$hash = hash("md5", \Input::get('master').\Input::get('slave'));
		$chars = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789&#@\',.()[]{}<>+-/*';
		$password = '';
		for($i = 0; $i < strlen($hash) / 2; ++$i)
		{
			$password = $password.$chars[intval(substr($hash, $i * 2, 2), 16) % strlen($chars)];
		}
		\Session::flash('result', $password);
		\Input::flash();
		return \Redirect::to(\URL::previous());
	}
};

?>