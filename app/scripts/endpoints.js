/**
 * Login function for the user
 */
function login() {
	getEndpoint(
		"users",
		"GET",
		{
			Authorization:
				"ApiKey " + document.getElementById("login.psw").value,
		},
		{}
	);
}

/**
 * Login function for the user
 */
function registers() {
	postEndpoint(
		"users",
		"POST",
		{},
		{},
		{
			name: document.getElementById("register.name").value,
		}
	);
}
