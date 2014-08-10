            <h1>Articles par ordre antéchronologique</h1>
            <ul>
<?php
foreach($articles as $article)
{
?>
                <li><?php echo $article->created_at->format('d.m.Y'); ?> / <a href="article/<?php echo $article->slug; ?>"><?php echo $article->title; ?></a></li>
<?php
}
?>
            </ul>
