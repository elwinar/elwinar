<?php extend('app') ?>
			<h1>New article</h1>
<?php if( ! $errors->isEmpty()) { ?>
			<ul class="errors">
<?php 	foreach($errors->all() as $error) { ?>
				<li><?= $error ?></p>
<?php 	} ?>
<?php } ?>
			</ul>
			<form method="POST">
				<input type="hidden" name="_token" value="<?= csrf_token() ?>">
				<input type="text" id="title" name="title" placeholder="Title" value="<?= Input::old('title', ''); ?>" required>
				<input type="text" class="slug" data-target="#title" name="slug" placeholder="Slug" value="<?= Input::old('slug', ''); ?>" required>
				<input type="text" name="tagline" placeholder="Tagline" value="<?= Input::old('tagline', ''); ?>" required>
				<textarea name="text" rows="10" placeholder="Text"><?= Input::old('text', ''); ?></textarea>
				<input type="text" name="tags" placeholder="Tags" value="<?= Input::old('tags', ''); ?>">
				<input type="submit" value="Create">
			</form>
