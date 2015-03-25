<?php extend('app') ?>
			<header>
				<h1><?= $article->title ?></h1>
				<?= Markdown::string($article->tagline) ?>
			</header>
			<?= Markdown::string($article->text) ?>
