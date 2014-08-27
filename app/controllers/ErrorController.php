<?php 

class ErrorController extends BaseController {

    protected $layout = 'layout';

	protected function display($view, $metadatas = array(), $data = array())
	{
		if(View::exists(subdomain().'.nav')) {
			$this->layout->nav = View::make(subdomain().'.nav', $data);
		}
		
		if(View::exists(subdomain().'.'.$view)) {
			$this->layout->content = View::make(subdomain().'.'.$view, $data);
		} else {
			$this->layout->content = View::make('errors.'.$view, $data);
		}
	}
    
    public function error($code, $exception)
    {
		switch($code) {
			case 404:
				$this->error404();
				break;
				
			default:
				$this->display('error', array(), array(
					'code' => $code,
					'exception' => $exception,
				));
		}
    }
    
    public function error404() {
		$this->display('404');
    }
    
};

?>