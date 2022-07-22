
require("expose-loader?exposes=$,jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");

let flash = document.getElementById("flash")
let main = document.getElementById("main")


setInterval(() => {
    if (!flash.classList.contains("d-none")){
        main.classList.add("pb-0")
        
     }else{
        main.classList.remove("pb-0")
     }
}, 100);



flash.addEventListener("click", (e) => {
    ev = e.target
    if (ev.classList.contains("close-flash")) {
        flash.classList.add("d-none")
        modal.classList.add("d-none")
    }
})

setTimeout(() => {
    flash.style.animation ="fadeIn 1s ease-out"
    flash.style.opacity = 0
    setTimeout(() => {
        flash.classList.add("d-none")
    }, 2000)   

    
}, 2000)



// let searchUser = document.querySelector("#search-users")
// let users = document.querySelectorAll(".user")

// let searchTasks = document.querySelector("#search-tasks")
// let tasks = document.querySelectorAll(".tasks")

   

let table = document.querySelector("#table")
if (table.childNodes.item(3).childNodes.length !=1) {
    document.querySelector("#text-table").classList.add("d-none")
    
}
// function preventBack(){window.history.forward();}
// setTimeout(preventBack(), 0);
// window.onunload=function(){null};

