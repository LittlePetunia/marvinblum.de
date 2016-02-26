$(document).ready(function(){
	// login
	var login_error = $('#login_error');
	login_error.hide();

	$('#dologin').click(function(){
		var login = $('#login').val();
		var pwd = $('#pwd').val();

		if(login == "" || pwd == ""){
			login_error.show();
		}
		else{
			login_error.hide();
			pwd = CryptoJS.SHA256(pwd).toString();

			$.post('/login', JSON.stringify({login: login, password: pwd}), function(resp){
				var json = JSON.parse(resp);

				if(json.success){
					window.location = '/';
				}
				else{
					login_error.show();
				}
			});
		}
	});
});
