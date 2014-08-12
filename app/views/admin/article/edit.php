			<nav class="pull-right contextual">
				<ul class="nav nav-stacked">
					<li><a href="<?= URL::previous(); ?>">Retour</a></li>
				</ul>
			</nav>
			<h1>Modifier un article</h1>
<?php
if(Session::has('error'))
{
?>
			<p class="error"><?= Session::get('error'); ?></p>
<?php
}
?>
			<form method="POST">
				<input class="form-control" type="text" name="title" placeholder="Titre" value="<?= $article->title; ?>" required><br>
				<input class="form-control" type="text" name="tagline" placeholder="Accroche" value="<?= $article->tagline; ?>" required><br>
				<textarea class="form-control" name="text" rows="10"><?= $article->text; ?></textarea><br>
				<input class="form-control" type="text" name="tags" placeholder="Etiquettes" value="<?= $article->tags; ?>"><br>
				<input class="form-control" type="submit" value="Modifier">
			</form>
