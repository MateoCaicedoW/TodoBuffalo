
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


const userSelectHidden = document.querySelector("#task-UserID")
const  userSelect= document.querySelector("#form-select-user")

let arrayPath = window.location.pathname.split("/")

for (const iterator of arrayPath) {
    if (iterator.includes("new") || iterator.includes("edit")) {
    
        userSelect.addEventListener("change", ()=>{
            userSelectHidden.value = userSelect.value
            if (userSelectHidden.value == "Open this select menu") {
                userSelectHidden.value= "00000000-0000-0000-0000-000000000000"
               
            }
        })
    
        if (window.location.pathname != "/new" )  {
            userSelect.value = userSelectHidden.value    
        }
        
        
    }
  
}


    

