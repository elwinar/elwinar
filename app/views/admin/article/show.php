			<nav class="pull-right contextual">
				<ul class="nav nav-stacked">
					<li><a href="articles">Retour à la liste</a></li>
					<li><a href="article/<?= $article->slug; ?>/edit">Modifier</a></li>
				</ul>
			</nav>
<?= Markdown::string($article->text); ?>
