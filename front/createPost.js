function makePost() {
    document.getElementById('myForm').addEventListener('submit', function(e) {
        // Предотвращаем стандартное поведение формы
        e.preventDefault();
      
        // Получаем данные формы
        var formData = new FormData(this);
      
        // Создаем объект для отправки данных в формате JSON
        var object = {};
        formData.forEach((value, key) => {
          object[key] = value;
        });
        var json = JSON.stringify(object);
      
        // Отправляем POST запрос на сервер
        fetch('http://localhost:80/post/', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: json,
        })
        .then(response => {
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          return response.json();
        })
        .then(data => {
          console.log(data);
          alert('Форма успешно отправлена!');
        })
        .catch(error => {
          console.error('Произошла ошибка при отправке формы:', error);
        });
      });
}

window.onload = makePost;