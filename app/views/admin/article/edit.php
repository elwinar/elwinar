            <h1>Modifier un article</h1>
<?php
if(Session::has('error'))
{
?>
            <p class="error"><?php echo Session::get('error'); ?></p>
<?php
}
?>
            <form method="POST">
                <input type="text" name="title" placeholder="Titre" value="<?php echo $article->title; ?>" required><br>
                <input type="text" name="tagline" placeholder="Accroche" value="<?php echo $article->tagline; ?>" required><br>
                <textarea name="text"><?php echo $article->text; ?></textarea><br>
                <input type="text" name="tags" placeholder="Etiquettes" value="<?php echo $article->tags; ?>"><br>
                <input type="submit" value="Modifier">
            </form>
