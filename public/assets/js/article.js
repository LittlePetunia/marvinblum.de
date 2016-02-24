$(document).ready(function(){
	// add article
	var input_error = $('#input_error');
	input_error.hide();

	$('#create').click(function(){
		input_error.hide();

		var title = $('#title').val();
		var link = $('#link').val();
		var picture = $('#picture').val();

		if(title == "" && link == ""){
			input_error.show();
		}
		else{
			$.post('/addArticle', JSON.stringify({title: title,
				link: link,
				picture: picture}), function(resp){
				var json = JSON.parse(resp);

				if(json.success){
					location.reload();
				}
				else{
					input_error.show();
				}
			});
		}
	});
});
