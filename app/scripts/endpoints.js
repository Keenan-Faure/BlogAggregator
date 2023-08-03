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

/**
 * Fetch Feeds for user
 */
function feed(api_key) {
	getEndpoint(
		"feed",
		"GET",
		{
			Authorization: "ApiKey " + api_key,
		},
		{
			page: "1",
			sort: "asc",
		}
	);
}

/**
 * Adds a new feed to the database for the user
 */
function createFeed(api_key) {
	postEndpoint(
		"feed",
		"POST",
		{
			Authorization: "ApiKey " + api_key,
		},
		{},
		{
			name: document.getElementById("feed.name").value,
			url: document.getElementById("feed.url").value,
		}
	);
}
