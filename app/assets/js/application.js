
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

// let path = window.location.pathname
// let array = path.split("/")

// if (array.includes("users")) {
//     Search(searchUser, users)  
// }
// if (array.includes("todo")) {
//     Search(searchTasks, tasks)    
// }





// function Search(searchInput, items) {
//     searchInput.addEventListener("keyup", (e) => {
//         let search = e.target.value
//         let string = String(search.toLowerCase()) 
//         string = string.split(" ").join(""); 
       

//         items.forEach((item) => {
//             node = item.childNodes[3].innerText

//             if (node.toLowerCase().includes(string.toLowerCase())) {
//                 item.classList.remove("d-none")
//             } else {
//                 item.classList.add("d-none")
//             }
//         })
//     })

// }