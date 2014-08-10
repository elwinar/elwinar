            <nav>
                <ul>
                    <li><a href="articles/new">Nouvel article</a></li>
                </ul>
            </nav>
            <table>
                <thead>
                    <tr>
                        <th>Titre</th>
                        <th>Date</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
<?php
foreach($articles as $article)
{
?>
                    <tr>
                        <td><?php echo $article->title; ?></td>
                        <td><?php echo $article->created_at->format('d.m.Y'); ?></td>
                        <td><a href="article/<?php echo $article->slug; ?>">Voir</a> <a href="article/<?php echo $article->slug; ?>/edit">Modifier</a> <a href="article/<?php echo $article->slug; ?>/delete">Supprimer</a></td>
                    </tr>
<?php
}
?>
                </tbody>
            </table>
