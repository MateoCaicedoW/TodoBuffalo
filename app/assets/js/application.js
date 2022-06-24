require("expose-loader?exposes=$,jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");

let btnCloseFlash = document.getElementById("close-flash")
let flashError = document.getElementById("flash")
btnCloseFlash.addEventListener("click", () => {
    flashError.classList.add("d-none")
})

