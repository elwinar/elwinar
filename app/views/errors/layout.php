<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <base href="<?php echo Request::root(); ?>">
<?php
if(!empty($title))
{
?>
        <title>Elwinar - <?php echo $title; ?></title>
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
        <meta name="description" content="<?php echo $description; ?>">
<?php
}

if(isset($keywords))
{
?>
        <meta name="keywords" content="<?php echo $keywords; ?>">
<?php
}
?>
        <link rel="stylesheet" href="css/normalize.min.css">
        <link rel="stylesheet" href="css/futura-light.css">
        <link rel="stylesheet" href="css/style.css">
    </head>
    <body>
        <header>
            <a href=".">
                <h1>Elwinar</h1>
                <p class="headline">Minimalist by design</p>
            </a>
        </header>
        <nav>
            <ul>
            </ul>
        </nav>
<?php
flush();
?>
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
