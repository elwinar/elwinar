<?php namespace App\Http\Requests;

use App\Http\Requests\Request;
use Auth;

class EditArticleRequest extends Request {

	/**
	 * Determine if the user is authorized to make this request.
	 *
	 * @return bool
	 */
	public function authorize()
	{
		return Auth::check();
	}

	/**
	 * Get the validation rules that apply to the request.
	 *
	 * @return array
	 */
	public function rules()
	{
		return [
			'title' => 'required|max:50',
			'slug' => 'required|unique:articles,slug,'.$this->route('article')->id.'|max:50',
			'tagline' => 'required|max:300',
			'text' => 'required',
			'tags' => 'required',
		];
	}

}
