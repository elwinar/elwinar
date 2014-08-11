<?php
foreach(File::files(app_path().'/routes') as $file) {
	Route::group(array(
		'domain' => basename($file, '.php').'.'.Config::get('app.domain')
	),function() use($file) {
		include($file);
	});
}
?>