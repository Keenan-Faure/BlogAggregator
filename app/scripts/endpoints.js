/**
 * Login function for the user
 */
function login() {
    fetchEndpoint(
        "users",
        "GET",
        {
            Authorization:
                "ApiKey " + document.getElementById("login.psw").value,
        },
        {},
        {}
    );
}

/**
 * Login function for the user
 */
function registers() {
    fetchEndpoint(
        "users",
        "POST",
        {},
        {},
        {
            name: document.getElementById("register.name").value,
        }
    );
}
