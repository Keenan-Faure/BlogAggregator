/**
 * Async function that fetches data from the API (GET, DELETE)
 *
 * Headers & Params are arrays in key, value pairs
 * ```json
 * {
 *  "key": "value"
 * }
 * ```
 *
 * @param {string} endpoint API Endpoint to request data from
 * @param {string} method HTTP Option either `GET`, `POST`, `PUT`, `DELETE`
 * @param {Object} headers Array of headers to be sent along with request
 * @param {Object} params Array of params to be sent as query_params with the request
 * @returns {any}
 */
const getEndpoint = async function (
	endpoint,
	method,
	headers = {},
	params = {}
) {
	let url = createURL(endpoint, params);
	const resp = await fetch(url, {
		method: method, // *GET, POST, PUT, DELETE, etc.
		mode: "cors", // no-cors, *cors, same-origin
		cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
		credentials: "include", // include, *same-origin, omit
		headers: AppendHeaders(headers),
		redirect: "follow", // manual, *follow, error
		referrerPolicy: "no-referrer", // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
	});
	const json = await resp.json();
	console.log(json);
	console.log(endpoint);
	let adHocMessage = EndpointAdHoc(endpoint, method, json);
	if (adHocMessage != "undefined") {
		Message(adHocMessage, resp.status);
	}
	return json;
};

/**
 * Async function that fetches data from the API (POST, PUT)
 *
 * Headers & Params are arrays in key, value pairs
 * ```json
 * {
 *  "key": "value"
 * }
 * ```
 *
 * @param {string} endpoint API Endpoint to request data from
 * @param {string} method HTTP Option either `GET`, `POST`, `PUT`, `DELETE`
 * @param {Object} headers Array of headers to be sent along with request
 * @param {Object} params Array of params to be sent as query_params with the request
 * @param {Object} bodyData Object containing key-value pairs to be sent with POST request
 * @returns {any}
 */
const postEndpoint = async function (
	endpoint,
	method,
	headers = {},
	params = {},
	bodyData = {}
) {
	let url = createURL(endpoint, params);
	const resp = await fetch(url, {
		method: method, // *GET, POST, PUT, DELETE, etc.
		mode: "cors", // no-cors, *cors, same-origin
		cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
		credentials: "include", // include, *same-origin, omit
		headers: AppendHeaders(headers),
		redirect: "follow", // manual, *follow, error
		referrerPolicy: "no-referrer", // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
		body: JSON.stringify(bodyData),
	});
	const json = await resp.json();
	console.log(json);
	console.log(endpoint);
	let adHocMessage = EndpointAdHoc(endpoint, method, json);
	if (adHocMessage != "undefined") {
		Message(adHocMessage, resp.status);
	}
	return json;
};

/**
 *
 * @param {string} endpoint endpoint that was queried
 * @param {any} json response from API
 * @returns {string}
 */
function EndpointAdHoc(endpoint, method, json) {
	if ((endpoint == "users" && method) == "GET") {
		setTimeout(() => {
			window.location.href = "app/dashboard.html";
		}, 500);
		return "Success";
	} else if (endpoint == "users" && method == "POST") {
		document.getElementById("login.psw").value = json.api_key;
		localStorage.setItem("api_key", json.api_key);
		return "Pasted token inside password";
	}
}

/**
 * Creates the request url
 *
 * @param {string} endpoint Endpoint to request data from
 * @param {Object} params Object containing params to send with request url
 * @returns {string}
 */
function createURL(endpoint, params) {
	arrayUrl = document.URL.split("/");
	url = "http://" + arrayUrl[2] + "/v1/" + endpoint;
	return appendParams(url, params);
}

/**
 * Appends the params key-value pairs to the `url`
 *
 * @param {string} url Current URL
 * @param {Object} params Params object
 * @returns {string}
 */
function appendParams(url, params) {
	if (Object.keys(params).length !== 0) {
		url += "?";
	}
	for (const key in params) {
		url += key + "=" + params[key];
	}
	return url;
}

/**
 * Appends the basic headers to the headers object
 *
 * @param {Object} headers
 * @return {Object}
 */
function AppendHeaders(headers) {
	headers["Access-Control-Allow-Origin"] = "*";
	headers["Content-Type"] = "application/json";
	return headers;
}

// sessionStorage.setItem("lastname", "Smith");
// let personName = sessionStorage.getItem("lastname");
// document.getElementById("demo").innerHTML = personName;
