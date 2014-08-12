<?php
if(Auth::guest())
{
?>
					<li><a href="login">Connexion</a></li>
<?php
}
else
{
?>
					<li><a href="articles">Articles</a></li>
					<li><a href="logout">Déconnexion</a></li>
<?php
}
?>
