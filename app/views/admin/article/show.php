            <nav>
                <ul>
                    <li><a href="articles">Retour à la liste</a></li>
                    <li><a href="article/<?php echo $article->slug; ?>/edit">Modifier</a></li>
                </ul>
            </nav>
<?php echo Markdown::string($article->text); ?>
