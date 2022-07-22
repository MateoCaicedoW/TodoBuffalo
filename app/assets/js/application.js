
require("expose-loader?exposes=$,jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");

let flash = document.getElementById("flash")

flash.addEventListener("click", (e) => {
    ev = e.target
    if (ev.classList.contains("close-flash")) {
        flash.classList.add("d-none")
    }
})


// let searchUser = document.querySelector("#search-users")
// let users = document.querySelectorAll(".user")

// let searchTasks = document.querySelector("#search-tasks")
// let tasks = document.querySelectorAll(".tasks")

   

let table = document.querySelector("#table")
if (table.childNodes.item(3).childNodes.length !=1) {
    document.querySelector("#text-table").classList.add("d-none")
    
}
