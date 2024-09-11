document.addEventListener("DOMContentLoaded",()=>{
  const root = document.getElementById('root')

  const list = [1,2,3,4,5]
  for(i=0;i<list.length;i++){
    const page = document.createElement("p")
    page.textContent = list[i]

    root.appendChild(page)
  }
})