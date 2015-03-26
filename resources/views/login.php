<?php extend('app') ?>
			<form method="POST">
				<input type="hidden" name="_token" value="<?= csrf_token() ?>">
				<input type="hidden" name="email" value="<?= env('APP_EMAIL') ?>">
				<input type="password" name="password" placeholder="Password">
				<input type="submit" value="Login">
			</form>
