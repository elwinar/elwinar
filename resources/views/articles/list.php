<?php extend('app') ?>
			<h1>Articles in ante-chronological order</h1>
<?php if(Auth::check()) { ?>
			<p><a href="/articles/new">New</a></p>
<?php } ?>
			<ul class="articles">
<?php foreach($articles as $article) { ?>
				<li>
					<?= $article->created_at->format('Y.m.d') ?> / <a href="/article/<?= $article->slug ?>" title="<?= $article->tagline ?>"><?= $article->title ?></a>
<?php 	if(Auth::check()) { ?>
					<nav>
<?php		if($article->is_published) { ?>
						<a href="/article/<?= $article->slug ?>/publish" title="publish"><span class="fa fa-eye-slash"></span></a>
<?php		} else { ?>
						<a href="/article/<?= $article->slug ?>/unpublish" title="unpublish"><span class="fa fa-eye"></span></a>
<?php		} ?>
						<a href="/article/<?= $article->slug ?>/edit" title="edit"><span class="fa fa-pencil"></span></a>
						<a href="/article/<?= $article->slug ?>/delete" title="delete"><span class="fa fa-trash-o"></span></a>
					</nav>
<?php 	} ?>
				</li>
<?php } ?>
			</ul>
