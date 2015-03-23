<!DOCTYPE html>
<html lang="<?= Config::get('app.locale') ?>">
	<head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<base href="<?= Request::root(); ?>">
<?php if(!empty($title)) { ?>
		<title>Elwinar - <?= $title; ?></title>
<?php } else { ?>
		<title>Elwinar</title>
<?php } ?>
		<meta name="author" content="Romain Baugue">
<?php if(isset($description)) { ?>
		<meta name="description" content="<?= $description; ?>">
<?php } ?>
<?php if(isset($keywords)) { ?>
		<meta name="keywords" content="<?= $keywords; ?>">
<?php } ?>
		<link rel="stylesheet" href="app.css">
		<script src="app.js"></script>
	</head>
	<body>
		<header>
			<a href="/">
				<h1>Elwinar</h1>
				<p class="headline">Minimalist by design</p>
			</a>
			<nav>
				<ul>
					<li><a href="//stackoverflow.com/users/3472656/elwinar">Ask</a></li>
					<li><a href="//github.com/elwinar">Code</a></li>
					<li><a href="mailto:<?= env('APP_EMAIL') ?>">Contact</a></li>
					<li><a href="articles">Write</a></li>
				</ul>
			</nav>
		</header>
		<section>
			<?= $_view ?>
		</section>
		<footer>
			<a href="http://validator.w3.org/check/referer"><span class="fa fa-html5"></span></a>
			<a href="http://jigsaw.w3.org/css-validator/check/referer"><span class="fa fa-css3"></span></a>
<?php if(Auth::check()) { ?>
			<a href="/logout"><span class="fa fa-lock"></span></a>
<?php } else { ?>
			<a href="/login"><span class="fa fa-unlock-alt"></span></a>
<?php } ?>
		</footer>
	</body>
</html>
