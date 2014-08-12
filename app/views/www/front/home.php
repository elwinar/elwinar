			<h1>Site personnel et minimaliste</h1>
			<p itemscope itemtype="http://schema.org/Person">Je m'appelle <span itemprop="name">Romain Baugue</span>, <meta itemprop="birthDate" content="1990-11-05"><?= intval(substr(date('Ymd') - date('Ymd', strtotime('19901105')), 0, -4)); ?> ans, <span itemprop="jobTitle">développeur</span>, passionné, laconique.</p>
			<h2>Derniers articles</h2>
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
