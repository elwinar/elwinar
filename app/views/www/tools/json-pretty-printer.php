			<h1>Formatteur JSON</h1>
			<form method="POST" id="pretty-printer">
				<textarea class="form-control <?php if(Session::get('result', '') === 'null'){ echo 'error'; } ?>" name="data" rows="10"><?= Input::old('data', ''); ?></textarea>
				<input class="form-control" type="submit" name="submit" value="Formatter"><br>
				<textarea class="form-control" name="result" readonly rows="10"><?= (Session::get('result', '') !== 'null')?Session::get('result', ''):''; ?></textarea>
			</form>
			<script>
var form = document.getElementById('pretty-printer');
form.removeChild(form.elements.submit);

function process()
{
	try
	{
		form.elements.data.setAttribute('class', 'form-control');
		form.elements.result.value = JSON.stringify(JSON.parse(form.elements.data.value), undefined, 4);
	}
	catch(error)
	{
		form.elements.data.setAttribute('class', 'form-control error');
	}
}

var timer;
function listen()
{
	clearTimeout(timer);
	timer = setTimeout(process, 100);
}

form.elements.data.addEventListener('input', listen, false);
			</script>
