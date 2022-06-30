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
let form = document.querySelector('#form-entry')

let select = document.querySelector('#entries')
select.addEventListener('change', (e)=>{
    let select = e.target

    let inputHidden = document.querySelector('#entry-select')
    inputHidden.value = select.value
    console.log(inputHidden.value)
    form.submit()
})




