<?php extend('app') ?>
<?php 	if(Auth::check()) { ?>
			<nav>
<?php		if($article->is_published) { ?>
				<a href="/article/<?= $article->slug ?>/unpublish" title="unpublish"><span class="fa fa-eye-slash"></span></a>
<?php		} else { ?>
				<a href="/article/<?= $article->slug ?>/publish" title="publish"><span class="fa fa-eye"></span></a>
<?php		} ?>
				<a href="/article/<?= $article->slug ?>/edit" title="edit"><span class="fa fa-pencil"></span></a>
				<a href="/article/<?= $article->slug ?>/delete" title="delete"><span class="fa fa-trash-o"></span></a>
			</nav>
<?php 	} ?>
			<h1><?= $article->title ?> <small><?= $article->created_at->format('Y.m.d') ?></small></h1>
			<?= Markdown::string($article->text) ?>
