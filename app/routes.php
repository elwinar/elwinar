<?php
foreach(File::files(app_path().'/routes') as $file) {
	Route::group(array(
		'domain' => basename($file, '.php').'.'.domain()
	),function() use($file) {
		include($file);
	});
}
?>