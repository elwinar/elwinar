<?php extend('app') ?>
			<form method="POST">
				<input type="hidden" name="_token" value="<?= csrf_token() ?>">
				<input type="email" name="email" placeholder="Email">
				<input type="password" name="password" placeholder="Password">
				<input type="submit" value="Login">
			</form>
