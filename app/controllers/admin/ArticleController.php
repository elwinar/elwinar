<?php namespace Admin;

class ArticleController extends \BaseController {
	
    protected $layout = 'admin.layout';
    
    public function __construct()
    {
        $this->beforeFilter('auth');
    }
	
	public function index()
	{
        $articles = \Article::orderBy('created_at', 'desc')->get();
        $this->display('admin.article.index', array('title' => 'Liste des articles'), array('articles' => $articles));
	}
	
	public function show($article)
	{
        $this->display('admin.article.show', array('title' => 'Article : '.$article->title), array('article' => $article));
	}
	
	public function create()
	{
        $this->display('admin.article.create', array('title' => 'Nouvel article'));
	}
	
	public function store()
	{
        $article = new \Article;
        $article->title = \Input::get('title');
        $article->slug = \Str::slug($article->title);
        $article->tagline = \Input::get('tagline');
        $article->text = \Input::get('text');
        $article->tags = \Input::get('tags');
        
        if(\Article::where('slug', $article->slug)->count() == 0)
        {
            $article->save();
            return \Redirect::to('article/'.$article->slug);
        }
        else
        {
            \Input::flash();
            \Session::flash('error', 'Nom de l\'article trop proche d\'un nom existant.');
            return \Redirect::to(\URL::previous());
        }
	}
	
	public function edit($article)
	{
        $this->display('admin.article.edit', array('title' => 'Article : '.$article->title), array('article' => $article));
	}
	
	public function update($article)
	{
        $article->title = \Input::get('title');
        $article->slug = \Str::slug($article->title);
        $article->tagline = \Input::get('tagline');
        $article->text = \Input::get('text');
        $article->tags = \Input::get('tags');
        
        if(\Article::where('slug', $article->slug)->where('id', '!=', $article->id)->count() == 0)
        {
            $article->save();
            return \Redirect::to('article/'.$article->slug);
        }
        else
        {
            \Input::flash();
            \Session::flash('error', 'Nom de l\'article trop proche d\'un nom existant.');
            return \Redirect::to(\URL::previous());
        }
	}
	
	public function destroy($article)
	{
        $article->delete();
        return \Redirect::to(\URL::previous());
	}
	
}

?>