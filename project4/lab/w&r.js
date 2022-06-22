const path = require('path')
const fs = require('fs')

function write(path,data){
    let date = new Date;
    let seconds = date.getSeconds();
    let minutes = date.getMinutes();
    let hour = date.getHours()

    let year = date.getFullYear();
    let month = date.getMonth() + 1; // beware: January = 0; February = 1, etc.
    let day = date.getDate();
    let publishedAt = (`${year}-${month}-${day} ${hour}:${minutes}`);

    let base = fs.readFileSync(path)
    base = JSON.parse(base)

    data.id = base.length  ?  base[base.length-1].id + 1 : 1
    data.publishedAt = publishedAt

    base.push(data)


    fs.writeFileSync(path,JSON.stringify(base,null,4))

}

module.exports = write