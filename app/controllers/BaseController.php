<?php

class BaseController extends Controller {

    protected function display($view, $metadatas, $data = array())
    {
		foreach($metadatas as $key => $value)
		{
			$this->layout->$key = $value;
		}
        $this->layout->content = View::make($view, $data);
    }

	/**
	 * Setup the layout used by the controller.
	 *
	 * @return void
	 */
	protected function setupLayout()
	{
		if ( ! is_null($this->layout))
		{
			$this->layout = View::make($this->layout);
		}
	}

};

?>