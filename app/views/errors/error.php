			<h1>Erreur <?= $code ?></h1>
			<p>Il y a eu un problème pendant le traitement de votre requête. Essayez de <a href="<?= URL::previous(); ?>">revenir en arrière</a> et de recommencer. Si le problème persiste, <a href="mailto:romain.baugue@elwinar.com">faites-le moi savoir</a> afin que je puisse le corriger. Et en attendant, vous pouvez repartir depuis <a href=".">l'accueil</a>.
			<pre><?= $exception ?></pre>
