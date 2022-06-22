// const blogTitleField = document.querySelector('.title')
const article = document.querySelector('.article')
const title = document.querySelector('.title');
const btn = document.querySelector('.publish-btn')




btn.addEventListener('click', () => {
    let value = {article: article.value ,title:  title.value}
    
    var xhr = new XMLHttpRequest();
    xhr.open("POST", 'http://localhost:1111/post-blog', true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify({
    value
    }));
    
    setTimeout(()=>{
        location.href = '/blogs'
    },500)

})