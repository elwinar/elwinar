			<h1>Hasheur</h1>
			<form method="POST" id="hasher">
				<textarea name="data"><?php echo Input::old('data', ''); ?></textarea>
				<select name="algorithm">
<?php
foreach(hash_algos() as $algorithm) {
?>
²					<option value="<?php echo $algorithm; ?>"<?php if(Input::old('algorithm', '') == $algorithm){ ?> selected<?php }?>><?php echo $algorithm; ?></option>
<?php
}
?>
				</select>
				<input type="submit" name="submit" value="Hasher">
				<input type="text" name="result" readonly value="<?php echo (Session::get('result', '') !== 'null')?Session::get('result', ''):''; ?>">
			</form>
