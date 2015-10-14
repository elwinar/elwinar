$(function(){
	$('input.slug').each(function(){
		var $this = $(this);
		var $target = $($this.data('target'));
		
		$target.on('keyup', function(){
			if($this.data('auto') === false) {
				return
			}
			$this.val(getSlug(this.value));
		});
		
		$this.on('keyup', function(){
			$this.data('auto', false);
		});
	});
});
