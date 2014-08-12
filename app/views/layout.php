<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<base href="<?= Request::root(); ?>">
<?php
if(!empty($title))
{
?>
		<title>Elwinar - <?= $title; ?></title>
<?php
}
else
{
?>
		<title>Elwinar</title>
<?php
}
?>
		<meta name="author" content="Romain Baugue">
<?php
if(isset($description))
{
?>
		<meta name="description" content="<?= $description; ?>">
<?php
}

if(isset($keywords))
{
?>
		<meta name="keywords" content="<?= $keywords; ?>">
<?php
}
?>
		<link rel="stylesheet" href="css/style.css">
	</head>
	<body class="container">
		<header class="navbar">
			<a href="." class="navbar-header">
				<h1>Elwinar</h1>
				<p class="headline">Minimalist by design</p>
			</a>
			<nav class="navbar-right">
				<ul class="nav nav-pills">
<?= $nav; ?>
				</ul>
			</nav>
		</header>
		<section>
<?= $content; ?>
		</section>
		<footer class="pull-right">
			<a href="http://www.w3.org/html/wg/drafts/html/master/"><span class="openwebicons-html5 red big"></span></a>
			<a href="http://validator.w3.org/check/referer"><span class="openwebicons-semantics black"></span></a>
			<a href="http://jigsaw.w3.org/css-validator/check/referer"><span class="openwebicons-css3 black"></span></a>
		</footer>
	</body>
</html>
