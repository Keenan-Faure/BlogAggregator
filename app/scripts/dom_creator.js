/**
 * Creates Registration DOM inside HTML body
 *
 * @param {HTMLElement} element Element that triggers the creation of the register element
 * @returns {void}
 */
function createRegister(element) {
    if (!registerExists()) {
        let register_name = document.createElement("input");
        register_name.id = "register.name";
        register_name.type = "text";
        register_name = appendStyleInput(register_name);

        let register_tag = document.createElement("p");
        register_tag.innerHTML = "Name: ";

        let register_button = document.createElement("button");
        register_button = appendStyleBtn(register_button);
        register_button.setAttribute("onclick", "registers()");
        register_button.innerHTML = "Register";

        element.insertAdjacentElement("afterend", register_button);
        element.insertAdjacentElement("afterend", document.createElement("br"));
        element.insertAdjacentElement("afterend", document.createElement("br"));
        element.insertAdjacentElement("afterend", register_name);
        element.insertAdjacentElement("afterend", register_tag);
    }
}

/**
 * Determines if the register function has been called already
 *
 * @return {bool}
 */
function registerExists() {
    register = document.getElementById("register");
    if (register != null) {
        return true;
    }
    return false;
}

/**
 * Appends style to the input element and returns it
 *
 * @param {HTMLElement} input_element Element to have styles added on to it
 * @returns {HTMLElement}
 */
function appendStyleInput(input_element) {
    input_element.style.border = "1px solid black";
    return input_element;
}

/**
 * Appends style to the button element and returns it
 *
 * @param {HTMLElement} button_element Element to have styles added on to it
 * @returns {HTMLElement}
 */
function appendStyleBtn(button_element) {
    button_element.style.position = "relative";
    button_element.style.left = "50%";
    button_element.style.transform = "translate(-50%, 0%)";
    return button_element;
}

/**
 * Creates an error message on the screen
 * @param {string} message Message from response
 * @param {Number} responseCode Response code
 */
function Message(message, responseCode) {
    let button = document.createElement("button");
    /** styles button */
    button.style.zIndex = "5";
    button.style.height = "25px";
    button.style.bottom = "15px";
    button.style.right = "15px";
    button.style.position = "absolute";
    button.style.border = "rgb(228 228 228)";
    button.style.opacity = "0.9";
    button.style.padding = "4px 12px";
    button.style.borderRadius = "5px";
    if (![200, 201].includes(responseCode)) {
        button.style.background = "linear-gradient(rgb(209, 125, 125), red)";
    } else {
        button.style.background = "linear-gradient(rgb(28, 243, 233), blue)";
    }
    button.innerHTML = message;

    referenceNode = document.body;
    referenceNode.parentNode.insertBefore(button, referenceNode.nextSibling);

    setTimeout(() => {
        button.remove();
    }, 1500);
}
