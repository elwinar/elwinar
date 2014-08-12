            <h1>Articles par ordre antéchronologique</h1>
            <ul>
<?php
foreach($articles as $article)
{
?>
                <li><?= $article->created_at->format('d.m.Y'); ?> / <a href="article/<?= $article->slug; ?>"><?= $article->title; ?></a></li>
<?php
}
?>
            </ul>
