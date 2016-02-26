$(document).ready(function(){
	// add comment
	var input_error = $('#input_error');
	var saved = $('#success');
	input_error.hide();
	saved.hide();

	$('#send').click(function(){
		input_error.hide();
		saved.hide();

		var name = $('#name').val();
		var email = $('#email').val();
		var comment = $('#comment').val();
		var article = $('#article').val();

		if(name == "" || email == "" || comment == ""){
			input_error.show();
		}
		else{
			$.post('/addComment', JSON.stringify({name: name,
				email: email,
				comment:comment,
				article: article}), function(resp){
				var json = JSON.parse(resp);

				if(json.success){
					saved.show();
					// TODO add to html
				}
				else{
					input_error.hide();
				}
			});
		}
	});

	// delete comment
	$('.deleteComment').click(function(){
		var article = $('#article').val();
		var created = $(this).attr('comment');

		$.post('/removeComment', JSON.stringify({article: article, created: created}), function(resp){
			var json = JSON.parse(resp);

			if(json.success){
				$('[created="'+created+'"]').remove();
			}
		});
	});
});
