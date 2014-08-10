<?php namespace Www;

class JsonPrettyPrinterController extends \BaseController {

    protected $layout = 'www.layout';
	
	public function show()
	{
		$this->display('www.tools.json-pretty-printer', array(
			'title' => 'Formatteur JSON',
			'description' => 'Un simple outil de remise en forme d\'une chaine JSON',
		));
	}
	
	public function process()
	{
		\Session::flash('result', json_encode(json_decode(\Input::get('data')),  JSON_PRETTY_PRINT | JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE));
		\Input::flash();
		return \Redirect::to(\URL::previous());
	}
};

?>