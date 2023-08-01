/**
 * Login function for the user
 */
function login() {
    const json = fetchEndpoint(
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
function register() {
    json = fetchEndpoint("users", "POST", {}, {});
}
