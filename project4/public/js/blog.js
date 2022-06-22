
fetch('/posts',{
    method: "get"
}).then(res => res.json())
.then(data => {
    // console.log(data[data.length-1]);
    setupBlog(data[data.length-1])
})

const setupBlog = (data) => {
    const banner = document.querySelector('.banner')
    const blogTitle = document.querySelector('.title')
    const titleTag = document.querySelector('title')
    const article = document.querySelector('.article')
    const publish = document.querySelector('.published')
  
    titleTag.innerHTML += blogTitle.innerHTML = data.title
    article.innerHTML += article.innerHTML = data.article
    publish.innerHTML += data.publishedAt;
   
}