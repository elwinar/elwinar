<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <base href="<?php echo Request::root(); ?>">
        <title><?php echo $title; ?> - Elwinar Administration</title>
        <link rel="stylesheet" href="css/normalize.min.css">
        <link rel="stylesheet" href="css/futura-light.css">
        <link rel="stylesheet" href="css/style.css">
		<link rel="stylesheet" href="css/highlight-themes/github.css">
		<script src="js/highlight.pack.js"></script>
		<script>hljs.initHighlightingOnLoad();</script>
    </head>
    <body>
        <header>
            <a href=".">
                <h1>Elwinar</h1>
            </a>
        </header>
        <nav>
            <ul>
				<li class="headline">Minimalist by design</li>
<?php
if(Auth::guest())
{
?>
                <li><a href="login">Connexion</a></li>
<?php
}
else
{
?>
                <li><a href="articles">Articles</a></li>
                <li><a href="logout">Déconnexion</a></li>
<?php
}
?>
            </ul>
        </nav>
        <section>
<?php echo $content; ?>
        </section>
        <footer>
            <a href="http://www.w3.org/html/wg/drafts/html/master/"><img src="images/html5.svg" alt="Badge HTML5" title="Utilise HTML5" class="badge"></a>
            <a href="http://validator.w3.org/check/referer"><img src="images/semantics.svg" alt="Insigne sémantique" title="Sémantiquement valide" class="insigna"></a>
            <a href="http://jigsaw.w3.org/css-validator/check/referer"><img src="images/styling.svg" alt="Insigne style" title="Mise en forme valide" class="insigna"></a>
        </footer>
    </body>
</html>
