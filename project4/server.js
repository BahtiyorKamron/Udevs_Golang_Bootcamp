const express = require('express')
const path = require('path')
const fs = require('fs')

const { json } = require('express/lib/response');
const exp = require('constants');
const bodyParser = require('body-parser');
const cors = require('cors')
const write = require('./lab/w&r')


let initial_path = path.join(__dirname,'public');

const app = express()
app.use(cors())
app.use(express.json())
app.use(express.static(initial_path))

app.set('views',path.join(process.cwd(),'public','js'))
app.set('view engine','ejs')

app.get("/",(req,res)=>{
    res.sendFile(path.join(initial_path,"uploads/home.html"))

})

app.get("/editor",(req,res)=>{
    res.sendFile(path.join(initial_path,"uploads/editor.html"))
})


app.post("/post-blog",(req,res)=>{

    let date = new Date;
    let seconds = date.getSeconds();
    let minutes = date.getMinutes();
    let hour = date.getHours()

    let year = date.getFullYear();
    let month = date.getMonth() + 1; // beware: January = 0; February = 1, etc.
    let day = date.getDate();
    let publishedAt = (`${year}-${month}-${day} ${hour}:${minutes}`);


    
    write(path.join(process.cwd(),'data','posts.json'),req.body.value)
    
    res.sendFile(path.join(initial_path,"uploads/blog.html"))
    console.log(req.body.value);


    
})
app.get("/blog/:id",(req,res)=>{
    let posts = fs.readFileSync(path.join(process.cwd(),'data','posts.json'))
    posts = JSON.parse(posts)
   
    let post = posts.find(el => el.id == req.params.id)
    
    res.render('post',{title: post.title, article: post.article , publishedAt: post.publishedAt})
})
app.get('/blogs',(req,res)=>{
    res.sendFile(path.join(initial_path,"uploads/blog.html"))
})

app.get("/posts",(req,res)=>{
    let data = fs.readFileSync(path.join(process.cwd(),"data",'posts.json'),'utf-8')

    data = JSON.parse(data)
    res.json(data)
})

app.get("/allThePosts",(_,res)=>{
    let data = fs.readFileSync(path.join(process.cwd(),"data",'posts.json'),'utf-8')

    if(!data.length){
        data = []
    }else{
        data = JSON.parse(data)
    }
    res.json({
        data
    })
})

app.listen(1111,()=> console.log('http://localhost:'+1111))