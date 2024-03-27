document.addEventListener("DOMContentLoaded", function() {
    // После полной загрузки страницы
    var loader = document.getElementById("loader");
    // Показываем загрузочную страницу
    loader.style.display = "block";

    // Скрываем загрузочную страницу после перехода на главную страницу
    window.onload = function() {
        loader.style.display = "none";
    };
});
    