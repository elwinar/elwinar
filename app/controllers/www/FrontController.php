<?php namespace Www;

class FrontController extends \BaseController {
    
    public function index()
    {
        $this->home();
    }

	public function home()
	{
        $articles = \Article::orderBy('created_at', 'desc')->take(5)->get();
        $this->display('front.home', array(
            'title' => 'Accueil',
            'description' => 'Site personnel de Romain Baugue, développeur spécialisé dans le web et l\'intelligence artificielle'
        ), array('articles' => $articles));
	}
    
    public function contact()
    {
        $this->display('front.contact', array('title' => 'Contact'));
    }
    
    public function send()
    {
        Mail::send('emails.contact', array('text' => \Input::get('text')), function($message)
        {
            $message->from(\Input::get('mail'));
            $message->to('romain.baugue@elwinar.com', 'Romain Baugue')->subject(\Input::get('subject'));
        });
		return Redirect::to('sent');
    }
	
	public function sent()
	{
        $this->display('front.sent', array('title' => 'Envoyé'));
	}
    
    public function articles()
    {
        $articles = \Article::orderBy('created_at', 'desc')->get();
        $this->display('front.articles', array(
            'title' => 'Articles',
            'description' => 'Liste des articles publiés sur ce site par ordre antéchronologique'
        ), array('articles' => $articles));
    }
    
    public function article($article)
    {
        $this->display('front.article', array(
            'title' => $article->title,
            'description' => $article->tagline,
            'keywords' => $article->tags
        ), array('article' => $article));
    }
	
	public function tools()
	{
        $this->display('front.tools', array('title' => 'Outils'));
	}

}