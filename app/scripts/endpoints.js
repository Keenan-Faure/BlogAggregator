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
 * @param {string} api_key APIKey of user
 * @returns {null}
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
 * @param {string} api_key APIKey of user
 * @returns {null}
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

/**
 * Fetches posts for the user
 * @param {string} api_key APIKey of user
 * @param {string} page Current page number
 * @returns {null}
 */
function posts(api_key, page) {
    getEndpoint(
        "posts",
        "GET",
        {
            Authorization: "ApiKey " + api_key,
        },
        {
            page: page,
            sort: "desc",
        }
    );
}

/**
 * Fetches posts for the user
 * @param {string} api_key APIKey of user
 * @param {string} page Current page number
 * @returns {null}
 */
function bookmarks(api_key, page) {
    getEndpoint(
        "bookmark",
        "GET",
        {
            Authorization: "ApiKey " + api_key,
        },
        {
            page: page,
            sort: "desc",
        }
    );
}

/**
 * Fetches posts for the user
 * @param {string} api_key APIKey of user
 * @param {string} page Current page number
 * @returns {null}
 */
function liked(api_key, page) {
    getEndpoint(
        "liked",
        "GET",
        {
            Authorization: "ApiKey " + api_key,
        },
        {
            page: page,
            sort: "desc",
        }
    );
}

/**
 * Fetches all followed feeds for the user
 * @param {string} api_key APIKey of user
 * @param {string} page Current page number
 * @returns {null}
 */
function liked(api_key, page) {
    getEndpoint(
        "liked",
        "GET",
        {
            Authorization: "ApiKey " + api_key,
        },
        {
            page: page,
            sort: "desc",
        }
    );
}

/**
 * Adds a new bookmark to the user
 * @param {string} apikey APIKey of user
 * @param {string} post_id Post ID
 * @returns {null}
 */
function bookmark(api_key, postid) {
    postEndpoint(
        "bookmark",
        "POST",
        {
            Authorization: "ApiKey " + api_key,
        },
        {},
        {
            post_id: postid,
        }
    );
}

/**
 * Adds a new bookmark to the user
 * @param {string} apikey APIKey of user
 * @param {string} post_id Post ID
 * @returns {null}
 */
function follow_feed(api_key, feedid) {
    postEndpoint(
        "feed_follows",
        "POST",
        {
            Authorization: "ApiKey " + api_key,
        },
        {},
        {
            feed_id: feedid,
        }
    );
}

/**
 * Adds a new liked feed to the user
 * @param {string} apikey APIKey of user
 * @param {string} page Current page number
 * @returns {null}
 */
function feedFollows(api_key, page) {
    getEndpoint(
        "feed_follows",
        "GET",
        {
            Authorization: "ApiKey " + api_key,
        },
        {},
        {
            page: page,
            sort: "desc",
        }
    );
}


