// Получаю ссылку на элемент типом 
const typeOfCategory = document.getElementById('type_of_category');

const category = document.getElementById('category')

// Добавляем обработчик события 'change'
typeOfCategory.addEventListener('change', function(){
    // Получаем тип категории
    const tCategory = typeOfCategory.value;

    switch (tCategory) {
        case "Доход":
            console.log("Выбран доход")
            break;
        case "Расход":
            console.log("Выбран расход")
            break;
        case "Перевод":
            console.log("Выбран перевод")
            break;
    }

    /*fetch('http://127.0.0.1:5050/addcategories')
        .then(response =>{
            if (!response.ok) {
                throw new Error('Ошибка сети: ${response.statusText}');
            }
            return response.json();
        })
        .then(data =>{
            console.log('Данные получены:', data);
        })*/
});