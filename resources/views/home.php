<?php extend('app') ?>
			<h1>Welcome</h1>
			<p vocab="http://schema.org/" typeof="Person">My name is <span property="name">Romain Baugue</span>, <meta property="birthDate" content="1990-11-05"><?= intval(substr(date('Ymd') - date('Ymd', strtotime('19901105')), 0, -4)); ?> years old, <span property="jobTitle">developer</span>, passionate, laconic.</p>
