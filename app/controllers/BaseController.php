<?php

class BaseController extends Controller {

    protected $layout = 'layout';

	protected function display($view, $metadatas = array(), $data = array())
	{
		foreach($metadatas as $key => $value)
		{
			$this->layout->$key = $value;
		}
		
		if(View::exists(subdomain().'.nav')) {
			$this->layout->nav = View::make(subdomain().'.nav', $data);
		}
		
		$this->layout->content = View::make(subdomain().'.'.$view, $data);
	}

	protected function setupLayout()
	{
		if ( ! is_null($this->layout))
		{
			$this->layout = View::make($this->layout);
		}
	}

};

?>