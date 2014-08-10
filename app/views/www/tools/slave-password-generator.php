			<h1>Générateur de mots de passe esclave</h1>
			<form method="POST" id="password-generator">
				<label>Mot de passe maître <input type="text" name="master" value="<?php echo \Input::old('master', ''); ?>"></label>
				<label>Clé esclave <input type="text" name="slave" value="<?php echo \Input::old('slave', ''); ?>"></label>
				<input type="submit" name="submit" value="Générer">
				<label>Mot de passe esclave <input type="text" name="result" value="<?php echo (Session::get('result', '') !== 'null')?Session::get('result', ''):''; ?>" readonly></label>
			</form>
			<script src="http://crypto-js.googlecode.com/svn/tags/3.1.2/build/rollups/md5.js"></script>
			<script>
var form = document.getElementById('password-generator');
form.removeChild(form.elements.submit);
var chars = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789&#@\',.()[]{}<>+-/*';

function process()
{
	var hash = CryptoJS.MD5(form.elements.master.value + form.elements.slave.value).toString();
	var password = '';
	for(var i = 0; i < hash.length / 2; ++i)
	{
		password += chars[parseInt(hash.substr(i * 2, 2), 16) % chars.length];
	}
	form.elements.result.value = password;
}

var timer;
function listen()
{
	clearTimeout(timer);
	timer = setTimeout(process, 100);
}

form.elements.master.addEventListener('input', listen, false);
form.elements.slave.addEventListener('input', listen, false);
			</script>
			<h2>Utilisation</h2>
			<p>Entrer le mot de passe maître, et une clé esclave. L'outil générera un mot de passe esclave composé de lettres, chiffres et symboles spéciaux d'un longeur de 16 caractères.</p>
			<p><em>Le mot de passe généré n'est pas aléatoire</em>, il est donc inutile de le retenir : le même mot de passe maître et la même clé esclave généreront toujours le même mot de passe.</p>
			<h2>Téléchargement</h2>
			<p>J'ai prit quelques minutes pour faire une version bureau de ce logicel (en C++ &amp; Qt), que je mets ici en téléchargement.</p>
			<ul>
				<li><a href="files/slave-password-generator.zip">Le logiciel prêt-à-exécuter (version Windows 32 bits)</a></li>
				<li><a href="https://github.com/elwinar/slave-password-generator">Le dépôt sur GitHub</a></li>
			</ul>
