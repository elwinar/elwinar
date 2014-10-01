			<h1>Hasheur</h1>
			<form method="POST" id="hasher">
				<textarea name="data"><?= Input::old('data', ''); ?></textarea>
				<select name="algorithm">
<?php
foreach(hash_algos() as $algorithm) {
?>
²					<option value="<?= $algorithm; ?>"<?php if(Input::old('algorithm', '') == $algorithm){ ?> selected<?php }?>><?= $algorithm; ?></option>
<?php
}
?>
				</select>
				<input type="submit" name="submit" value="Hasher">
				<input type="text" name="result" readonly value="<?= (Session::get('result', '') !== 'null')?Session::get('result', ''):''; ?>">
			</form>
