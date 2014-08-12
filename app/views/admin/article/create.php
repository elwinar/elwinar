			<nav class="pull-right contextual">
				<ul class="nav nav-stacked">
					<li><a href="<?= URL::previous(); ?>">Retour</a></li>
				</ul>
			</nav>
			<h1>Nouvel article</h1>
<?php
if(Session::has('error'))
{
?>
			<p class="error"><?= Session::get('error'); ?></p>
<?php
}
?>
			<form method="POST" class="form">
				<input class="form-control" type="text" name="title" placeholder="Titre" value="<?= Input::old('title', ''); ?>" required><br>
				<input class="form-control" type="text" name="tagline" placeholder="Description" value="<?= Input::old('tagline', ''); ?>" required><br>
				<textarea class="form-control" name="text" rows="10"><?= Input::old('text', ''); ?></textarea><br>
				<input class="form-control" type="text" name="tags" placeholder="Etiquettes" value="<?= Input::old('tags', ''); ?>"><br>
				<input class="form-control" type="submit" value="Créer">
			</form>
