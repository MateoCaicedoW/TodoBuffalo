
require("expose-loader?exposes=$,jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");

let flash= document.getElementById("flash")

flash.addEventListener("click", (e)=>{
    ev = e.target
    if (ev.classList.contains("close-flash-error") || ev.classList.contains("close-flash-success")) {
        flash.classList.add("d-none")
    }
})


    

