<?php extend('app') ?>
			<header>
				<h1><?= $article->title ?> <small><?= $article->created_at->format('Y.m.d') ?></small></h1>
				<?= Markdown::string($article->tagline) ?>
			</header>
			<?= Markdown::string($article->text) ?>
