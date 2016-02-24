$(document).ready(function(){
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
				if(resp.success){
					saved.show();
					// TODO add to html
				}
				else{
					input_error.hide();
				}
			});
		}
	});
});
