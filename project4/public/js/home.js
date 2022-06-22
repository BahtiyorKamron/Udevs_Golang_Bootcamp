let blogSection = document.querySelector(".blogs-section")

fetch('/allThePosts',{
    method: "get"
}).then(res => res.json())
.then(data => {
    
    data.data.forEach(element => {
        console.log(element);
        createBlog(element)
    });
})

const createBlog = (blog) =>{
   blogSection.innerHTML += `
   <div class="blog-card">
      <h1 class="blog-title">${blog.title.substring(0,100) + "..."}</h1>
      <p class="blog-overview">${blog.article.substring(0,100) + "..."}</p>
      <a href=/blog/${blog.id} class="btn dark">read</a>
   </div>

   `
}