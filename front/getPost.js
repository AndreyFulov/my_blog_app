function fetchPosts() {
    const params = new URLSearchParams(window.location.search);
    const id_param = params.get("id")
    fetch(`http://localhost:8080/post?id=${id_param}`)
        .then(response => response.json())
            .then(data => {
                const postInfo = document.createElement('article');
                    postInfo.innerHTML = `<h2>${data.title}</h2>
                    <p>${data.desc}</p>`;
                    postInfo.classList.add("extended-post")
                    document.getElementById("post-feed").appendChild(postInfo)
            })
}

window.onload = fetchPosts;