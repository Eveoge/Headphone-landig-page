const button = document.getElementsByTagName("button")[0]
const todoinput = document.getElementsByName("todoinput")[0]
const todos = document.getElementsByClassName("todos")[0]

console.log(button)
console.log(todoinput)
console.log(todos)

button.addEventListener("click", function(){
    if(todoinput.value==""){
        return ;}

    console.log ("hello world")
    var todo=todoinput.value

var newtodo=document.createElement("div")
newtodo.classList.add("todo")
newtodo.innerHTML=todo

todos.append(newtodo)
todoinput.value=""
})



console.log(new Date().getTime())