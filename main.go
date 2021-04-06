package main
import ("fmt"   //доступ функциям (напр: вывод на сайт,терминал)
        "net/http")   //показывать инф-ю пользователю, отслежевать его действия
        //"html/template")   //для работы с шаблонами html, графическая чать

type User struct {
  name string
  age uint16 //uint16-хранит целое неотрицательное число
  money int16 //int16-хранит целое число;
  average_grades, happiness float64
}

func (this User) getAllInfo() string { //вывод содержимого объекта
  return fmt.Sprintf("User name is: %s. \nHe is: %d years old." +
    "\nHis money equal: %d.",
      this.name,
      this.age,
      this.money)
}

func (this *User) setNewName(newName string) { //явно определяем УКАЗАТЕЛЬ(не копию); принимает string; ничего не возвращает;
    this.name = newName
}

func home_page(w_page http.ResponseWriter, r *http.Request) { // arg2(r)-запрос для проверки передачи данных
  bob := User{"Bob", 21, -100, 4.3, 0.7} //создание объекта
  bob.setNewName("Alex")
  fmt.Fprintf(w_page, bob.getAllInfo())
  // fmt.Fprintf(w_page, `<h1>Main Text</h1>
  // <b>Main Text</b>`)
  //templ, _ := template.ParseFiles("templates/home_page.html") //v1-хранит шаблон, v2-обработка ошибок
  //templ.Execute(w_page, bob) //отображение шаблона; arg2-объект User
}

func faq_page(w_page http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w_page, "*Text on FAQ page*") //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
}

func handleRequest() {
  http.HandleFunc("/", home_page) //arg1- отслеживание перехода по url; arg2-метод при arg1; /-главная страница (/about)
  http.HandleFunc("/faq/", faq_page)
  http.ListenAndServe(":8080", nil) //запуск локально сервера на порту:8080(любой);...
  //...arg1-порт по чтению сервера(любой свободный на ПК); arg2-настройки для сервера; nil-NULL
}

func main() {
  //var bob User = .... создание объекта
  //bob := User{name: "Bob", age: 21, money: -100, average_grades: 4.3, happiness: 0.7} //создание объекта
  //bob := User{"Bob", 21, -100, 4.3, 0.7} //создание объекта
  handleRequest() //отслеживаем url, запускаем сервер
}
