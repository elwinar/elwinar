<?php extend('app') ?>
			<h1>Articles in ante-chronological order</h1>
<?php if(Auth::check()) { ?>
			<p><a href="/articles/new">New</a></p>
<?php } ?>
			<ul class="articles">
<?php foreach(App\Article::all() as $article) { ?>
				<li><a href="/article/<?= $article->slug ?>"><?= $article->created_at->format('Y.m.d') ?> / <?= $article->title ?></a></li>
<?php } ?>
			</ul>
