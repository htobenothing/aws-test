
const fileInput = document.getElementById("fileInput")
const result = document.getElementById("result")
var file

fileInput.onchange =(e)=>{
    file= e.target.files[0]
    console.log(file)
}


const baseUrl = "http://localhost:8080"
const upload= ()=>{
    var fileData = new FormData()
    fileData.append("file",file) 
    fetch(baseUrl+"/upload",{
        method:"POST",
        body:fileData
    })
    .then(resp=>resp.json())
    .then(data=> console.log(data))
    .catch(err=>console.log(err))
}

const suButton = document.getElementById("btnSubmit")
suButton.onclick=()=>{
    result.innerHTML+="starting upload </br>"
    upload()
}