$(document).ready(function(){
	// add article
	var article_input_error = $('#article_input_error');
	article_input_error.hide();

	$('#create').click(function(){
		article_input_error.hide();

		var title = $('#title').val();
		var link = $('#link').val();
		var picture = $('#picture').val();

		if(title == "" && link == ""){
			article_input_error.show();
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
					article_input_error.show();
				}
			});
		}
	});

	// edit article
	$('#save').click(function(){
		article_input_error.hide();

		var article = $('#article').val();
		var title = $('#title').val();
		var link = $('#link').val();
		var picture = $('#picture').val();
		var content = CKEDITOR.instances.content.getData();

		if(title == "" && link == ""){
			article_input_error.show();
		}
		else{
			$.post('/saveArticle', JSON.stringify({article: article,
				title: title,
				link: link,
				picture: picture,
				content: content}), function(resp){
				var json = JSON.parse(resp);

				if(json.success){
					location.reload();
				}
				else{
					article_input_error.show();
				}
			});
		}
	});
});
