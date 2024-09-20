function fetchPosts() {
    fetch("http://localhost:80/posts/")
        .then(response => response.json())
            .then(data => {
                data.forEach(post => {
                    const postInfo = document.createElement('article');
                    postInfo.innerHTML = `<h2>${post.title}</h2>
                    <p>${String(post.desc).slice(0,50)+" ..."}</p>
                    <a href="post.html?id=${post.id}" class="btn">Читать далее</a>`;
                    document.getElementById("post-feed").appendChild(postInfo)
                });
            })
}

window.onload = fetchPosts;