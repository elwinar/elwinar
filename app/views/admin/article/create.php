            <h1>Nouvel article</h1>
<?php
if(Session::has('error'))
{
?>
            <p class="error"><?php echo Session::get('error'); ?></p>
<?php
}
?>
            <form method="POST">
                <input type="text" name="title" placeholder="Titre" value="<?php echo Input::old('title', ''); ?>" required><br>
                <input type="text" name="tagline" placeholder="Description" value="<?php echo Input::old('tagline', ''); ?>" required><br>
                <textarea name="text"><?php echo Input::old('text', ''); ?></textarea><br>
                <input type="text" name="tags" placeholder="Etiquettes" value="<?php echo Input::old('tags', ''); ?>"><br>
                <input type="submit" value="Créer">
            </form>
