			<h1>Liste des articles</h1>
			<table class="table table-striped">
				<thead>
					<tr>
						<th>Titre</th>
						<th>Date</th>
						<th colspan="3">Actions</th>
					</tr>
				</thead>
				<tfoot>
					<tr>
						<td colspan="5"><a href="articles/new">Nouvel article</a></td>
					</tr>
				</tfoot>
				<tbody>
<?php
if(count($articles) == 0) {
?>
					<tr>
						<td colspan="5">Aucun article</td>
					</tr>
<?php
} else {
	foreach($articles as $article)
	{
?>
					<tr>
						<td><?= $article->title; ?></td>
						<td><?= $article->created_at->format('d.m.Y'); ?></td>
						<td><a href="article/<?= $article->slug; ?>"><span class="entypo-eye"></span></a></td>
						<td><a href="article/<?= $article->slug; ?>/edit"><span class="entypo-brush"></span></a></td>
						<td><a href="article/<?= $article->slug; ?>/delete"><span class="entypo-cancel"></span></a></td>
					</tr>
<?php
	}
}
?>
				</tbody>
			</table>
