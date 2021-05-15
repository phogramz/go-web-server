package main
import ("fmt"   //доступ функциям (напр: вывод на сайт,терминал)
        "net/http"   //показывать инф-ю пользователю, отслежевать его действия
        "html/template" //для работы с шаблонами html, графическая чать
        "database/sql"
        _ "reflect"
	      _ "github.com/go-sql-driver/mysql"
        //_ "pkg/mod/github.com/go-sql-driver/mysql"
        //_ "pkg/mod/github.com/go-sql-driver/mysql@v1.6.0"
)

const(
  user = "mysql"
  pass = "mysql"
  host = "localhost"
  port = "3306"
  dbname = "golangdb"
)
  var db *sql.DB //база mysql

type TUser struct {
  Password string
  Email string
}

type TCrew struct {
  member1 string
  member2 string
  member3 string
  member4 string
  member5 string
  member6 string
  member7 string
}

type Mission struct {
  Name string `json:"Name"` //теги json (свойство Name будет преобразовано в ключ "Name")
  StartData string `json:"StartData"`
  FinishData string `json:"FinishData"`
  DayDuration uint16 `json:"DayDuration"`//uint16-хранит целое неотрицательное число
  Target string `json:"Target"`
  CarrierRocket string `json:"CarrierRocket"`
  Crew TCrew `json:"Crew"`
  CustomerCountry string `json:"CustomerCountry"`
  LaunchSite string `json:"LaunchSite"`
  Success bool `json:"Success"`
  CauseFailure string `json:"CauseFailure"`
}

func (this Mission) getAllInfo() string { //вывод содержимого объекта
  //if (this.Success == true)
  var returnStr string = "\nНазвание миссии: %s. \nДата запуска: %s." +
                     "\nМиссия завершена: %s. \nПродолжительность: %d дней." +
                     "\nОсновная цель миссии: %s. \nРакета-носитель: %s." +
                     "\nКоманда/Аппараты: %d. \nСтрана: %s." +
                     "\nКосмодром: %s. \nУспех миссии: %t." +
                     "\nПричина неудачи: %s."
  return fmt.Sprintf(returnStr,
      this.Name,
      this.StartData,
      this.FinishData,
      this.DayDuration,
      this.Target,
      this.CarrierRocket,
      this.Crew,
      this.CustomerCountry,
      this.LaunchSite,
      this.Success,
      this.CauseFailure)
}

func (this *Mission) setNewTarget(newTarget string) { //явно определяем УКАЗАТЕЛЬ(не копию); принимает string; ничего не возвращает;
    this.Target = newTarget
}

func main() {
  db, _ = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbname)) //подключение к БД
  defer db.Close()

//Заполнение таблиц
/*
  insert, err := db.Query("INSERT INTO `tbl_mission` (`ID`, `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, `Crew`, `CustomerCountry`, `LaunchSite`, `Success`, `CauseFailure`) VALUES" +
    "(1, 'Пионер-5', '11 марта 1960', '30 апреля 1960', 50, 'Исследование солнечных частиц и космического пространства.', 'Thor Able IV 219', 3, 'США', 'Мыс Канаверал', 1, '-')," +
    "(2, 'Helios-B', '15 января 1976', '23 декабря 1979', 1438, 'Приближение к Солнцу на 0,291 а.е.', 'Titan IIIE/Centaur', 2, 'Германия, США', 'Мыс Канаверал', 1, '-')," +
    "(3, 'Пионер-E', '27 августа 1969', '27 августа 1969', 0, 'Изучение солнечной плазмы и ветра, физики частиц.', 'Thor-Delta L', 1, 'США', 'Мыс Канаверал', 0, 'Взрыв ракеты-носителя через 8 минут после старта')," +
    "(4, 'Луна-3', '4 октября 1959', '20 апреля 1960', 200, 'Снимки обратной стороны Луны.', '8К72 Л1-8', 4, 'СССР', 'Байконур', 1, '-')," +
    "(5, 'Луна-19', '28 сентября 1971', '1 ноября 1972', 388, 'Картографирование лунной поверхности.', 'Проток-К/Блок Д', 5, 'СССР', 'Байконур Пл. 81/24', 0, 'Отказ системы управления-> Не выход на нужную орбиту, потеря связи.')," +
    "(6, 'Аполон-11', '16 июля 1969', '24 июля 1969', 8, 'Высадка на поверхность Луны.', 'Saturn V (SA-506)', 6, 'США', 'John F. Kennedy Space Center (39A)', 1, '-')," +
    "(7, 'Луноход-2', '15 января 1973', '4 июня 1973', 50, 'Исследование поверхности Луны.', 'Проток-К/Блок Д', 7, 'СССР', 'Байконур', 1, '-')," +
    "(8, 'Вояджер-2', '20 августа 1977', 'активен', 15940, 'Исследование дальних планет Солнечной системы.', 'Titan IIIE/ Центавр', 8, 'США', 'Мыс Канаверал', 1, '-')," +
    "(9, 'Акацуки (PLANET-C)', '20 мая 2010', 'активен', 3980, 'Изучение Венеры.', 'H-IIA202 (F17)', 9, 'Япония', 'Tanegashima Space Center', 1, '-')," +
    "(10, 'Новые Горизонты', '19 января 2006', 'активен', 5562, 'Изучение Плутона и его естественного спутника Харона.', 'Atlas V', 10, 'США', 'Мыс Канаверал', 1, '-')," +
    "(11, 'Марс-2020', '30 июля 2020', 'активен', 286, 'Исследование поверхности Марса. Первый управляемый полет на другой планете.', 'Atlas V (541)', 11, 'США', 'Мыс Канаверал', 1, '-')")

  //insert, err := db.Query("INSERT INTO `crewtable` (`member1`, `member2`, `member3`, `member4`, `member5`, `member6`, `member7`, `member8`, `member9`) VALUES('Bob', 'Alex', 'Frank', 'Prev', 'Ing', '-', '-', '-', '-')")
  if err != nil {
    panic(err)
  }
  defer insert.Close()
*/
  fmt.Println("Подключено к MySql")
  handleRequest() //отслеживаем url, запускаем сервер
}

//Обрабатывает переходы по url, вызывает соответсвующие методы
func handleRequest() {
  http.HandleFunc("/", index) //arg1- отслеживание перехода по url; /-главная страница
  http.HandleFunc("/faq/", faq_page)
  http.HandleFunc("/signup/", signUp_page)
  http.HandleFunc("/create_user/", newUser_page)
  http.HandleFunc("/pioner5/", Pioner5_page)
  http.HandleFunc("/heliosB/", HeliosB_page)
  http.HandleFunc("/pionerE/", PionerE_page)
  http.HandleFunc("/moon3/", Moon3_page)
  http.HandleFunc("/moon19/", Moon19_page)
  http.HandleFunc("/appolo11/", Appolo11_page)
  http.HandleFunc("/moonWalker2/", MoonWalker2_page)
  http.HandleFunc("/voyager2/", Voyager2_page)
  http.HandleFunc("/akatcuki/", Akatcuki_page)
  http.HandleFunc("/newHorizons/", NewHorizons_page)
  http.HandleFunc("/mars2020/", Mars2020_page)

  http.ListenAndServe(":8080", nil) //запуск локально сервера на порту:9080;...
  //...arg1-порт по чтению сервера(любой свободный на ПК); arg2-настройки для сервера; nil-NULL
}

//Отображает html шаблон всех страниц Mission(передает данные в html файл)
func (this *Mission) parsePage(w_page http.ResponseWriter, r *http.Request){
  tmpl, err := template.ParseFiles("templates/index.html",
                                   "templates/header.html",
                                   "templates/footer.html",
                                   "templates/mission.html",
                                   "templates/signup.html")
  if err != nil {
    panic(err)
  }
  tmpl.ExecuteTemplate(w_page, "mission", this)
}

//Отображает html шаблон Главной страницы
func index(w_page http.ResponseWriter, r *http.Request) { // arg2(r)-запрос для проверки передачи данных
  tmpl, err := template.ParseFiles("templates/index.html",
                                   "templates/header.html",
                                   "templates/footer.html",
                                   "templates/mission.html",
                                   "templates/signup.html",
                                   "templates/signup_success.html") //v1-хранит шаблон, v2-обработка ошибок
                                   if err != nil {
                                     panic(err)
                                   }
  tmpl.ExecuteTemplate(w_page, "index", nil) //отображение шаблона;
}

// func (s Mission) toString() string {
//         return fmt.Sprintf("%b", s)
// }

// func getType(myvar interface{}) string {
//     if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
//         return "*" + t.Elem().Name()
//     } else {
//         return t.Name()
//     }
// }

//Добавление данных о новом пользователе в БД (3)
func (this TUser) newUser() {
  insert, err := db.Query(fmt.Sprintf("INSERT INTO `tbl_user` (`email`, `password`) VALUES ('%s', '%s')", this.Email, this.Password))
  if err != nil {
    panic(err)
  }
  defer insert.Close()
}

//Обработка данных полученных при регистрации нового пользователя (2)
func newUser_page(w_page http.ResponseWriter, r *http.Request){
  var _user TUser
  _user.Email = r.FormValue("inputEmail") //запись значений внесенных пользователем
  _user.Password = r.FormValue("inputPassword")

  check_user_in_tbl, err := db.Query(fmt.Sprintf("SELECT `email` FROM tbl_user WHERE email = '%s'", _user.Email))
  if err != nil {
    panic(err)
  }
  var check_email string
  for check_user_in_tbl.Next() {
    err = check_user_in_tbl.Scan(&check_email)
  }
  defer check_user_in_tbl.Close()

  if check_email == "" {
  _user.newUser() //вызывает ф-ю создания записи о пользователе в БД
  } else {
  _user.Email = "ERROR_USER_ALREADY_EXISTS"
  }
  //После создания записи выводит страницу успешной регистрации
  tmpl, err := template.ParseFiles("templates/index.html",
                                   "templates/header.html",
                                   "templates/footer.html",
                                   "templates/mission.html",
                                   "templates/signup.html",
                                   "templates/signup_success.html")
                                 if err != nil {
                                   panic(err)
                                 }
  tmpl.ExecuteTemplate(w_page, "signup_success", _user)
}

//Вызывает html файлы для отображения страницы регистрации (1)
func signUp_page(w_page http.ResponseWriter, r *http.Request){
  tmpl, err := template.ParseFiles("templates/index.html",
                                   "templates/header.html",
                                   "templates/footer.html",
                                   "templates/mission.html",
                                   "templates/signup.html")
                                 if err != nil {
                                   panic(err)
                                 }
  tmpl.ExecuteTemplate(w_page, "signup", nil)
}

//Вызывает html файлы для отображения страницы FAQ
func faq_page(w_page http.ResponseWriter, r *http.Request) {
  res, err := db.Query("SELECT `title`, `id_genre` FROM tbl_1")
  if err != nil {
    panic(err)
  }
  for res.Next() {
    var title string
    var genre string
    err = res.Scan(&title, &genre)
    if err != nil {
      panic(err)
    }
    fmt.Fprintf(w_page, fmt.Sprintf("\n title: %s, \n genre: %s", title, genre))
  }
}

//Обработка перехода по ссылки /pioner5/ (делает выборку из БД, передает в ф-ю вывода шаблона)
func Pioner5_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `member1`, `member2`, `member3`, `member4`, `member5`, `member6`, `member7`" +
    " FROM tbl_crew WHERE ID_Crew = 1")
  if err != nil {
    panic(err)
  }
  var Pioner5_crew TCrew
  for res.Next() {
    err = res.Scan(&Pioner5_crew.member1,
                   &Pioner5_crew.member2,
                   &Pioner5_crew.member3,
                   &Pioner5_crew.member4,
                   &Pioner5_crew.member5,
                   &Pioner5_crew.member6,
                   &Pioner5_crew.member7)
               }
  res, err = db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, " +
    " `CustomerCountry`, `LaunchSite`, `Success`, `CauseFailure` FROM tbl_mission WHERE Name = 'Пионер-5'")
  if err != nil {
    panic(err)
  }
    for res.Next() {
    var Pioner5 Mission
    err = res.Scan(&Pioner5.Name,
                   &Pioner5.StartData,
                   &Pioner5.FinishData,
                   &Pioner5.DayDuration,
                   &Pioner5.Target,
                   &Pioner5.CarrierRocket,
                   &Pioner5.CustomerCountry,
                   &Pioner5.LaunchSite,
                   &Pioner5.Success,
                   &Pioner5.CauseFailure)
    if err != nil {
      panic(err)
    }
    Pioner5.Crew = Pioner5_crew
    Pioner5.parsePage(w_page, r)
    //fmt.Fprintf(w_page, Pioner5.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

//Обработка перехода по ссылки /heliosb/ (делает выборку из БД, передает в ф-ю вывода шаблона)
func HeliosB_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `member1`, `member2`, `member3`, `member4`, `member5`, `member6`, `member7`" +
    " FROM tbl_crew WHERE ID_Crew = 2")
  if err != nil {
    panic(err)
  }
  var HeliosB_crew TCrew
  for res.Next() {
    err = res.Scan(&HeliosB_crew.member1,
                   &HeliosB_crew.member2,
                   &HeliosB_crew.member3,
                   &HeliosB_crew.member4,
                   &HeliosB_crew.member5,
                   &HeliosB_crew.member6,
                   &HeliosB_crew.member7)
               }
  res, err = db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, " +
    " `CustomerCountry`, `LaunchSite`, `Success`, `CauseFailure` FROM tbl_mission WHERE Name = 'Helios-B'")
  if err != nil {
    panic(err)
  }
  for res.Next() {
    var HeliosB Mission
    err = res.Scan(&HeliosB.Name,
                   &HeliosB.StartData,
                   &HeliosB.FinishData,
                   &HeliosB.DayDuration,
                   &HeliosB.Target,
                   &HeliosB.CarrierRocket,
                   //&HeliosB.Crew,
                   &HeliosB.CustomerCountry,
                   &HeliosB.LaunchSite,
                   &HeliosB.Success,
                   &HeliosB.CauseFailure)
    if err != nil {
      panic(err)
    }
    HeliosB.Crew = HeliosB_crew
    HeliosB.parsePage(w_page, r)
    //fmt.Fprintf(w_page, HeliosB.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

func PionerE_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `member1`, `member2`, `member3`, `member4`, `member5`, `member6`, `member7`" +
    " FROM tbl_crew WHERE ID_Crew = 3")
  if err != nil {
    panic(err)
  }
  var PionerE_crew TCrew
  for res.Next() {
    err = res.Scan(&PionerE_crew.member1,
                   &PionerE_crew.member2,
                   &PionerE_crew.member3,
                   &PionerE_crew.member4,
                   &PionerE_crew.member5,
                   &PionerE_crew.member6,
                   &PionerE_crew.member7)
               }
  res, err = db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, " +
    " `CustomerCountry`, `LaunchSite`, `Success`, `CauseFailure` FROM tbl_mission WHERE Name = 'Пионер-E'")
  if err != nil {
    panic(err)
  }
  for res.Next() {
    var PionerE Mission
    err = res.Scan(&PionerE.Name,
                   &PionerE.StartData,
                   &PionerE.FinishData,
                   &PionerE.DayDuration,
                   &PionerE.Target,
                   &PionerE.CarrierRocket,
                   //&PionerE.Crew,
                   &PionerE.CustomerCountry,
                   &PionerE.LaunchSite,
                   &PionerE.Success,
                   &PionerE.CauseFailure)
    if err != nil {
      panic(err)
    }
    PionerE.Crew = PionerE_crew
    PionerE.parsePage(w_page, r)
    //fmt.Fprintf(w_page, PionerE.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

func Moon3_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `member1`, `member2`, `member3`, `member4`, `member5`, `member6`, `member7`" +
    " FROM tbl_crew WHERE ID_Crew = 4")
  if err != nil {
    panic(err)
  }
  var Moon3_crew TCrew
  for res.Next() {
    err = res.Scan(&Moon3_crew.member1,
                   &Moon3_crew.member2,
                   &Moon3_crew.member3,
                   &Moon3_crew.member4,
                   &Moon3_crew.member5,
                   &Moon3_crew.member6,
                   &Moon3_crew.member7)
               }
  res, err = db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, " +
    " `CustomerCountry`, `LaunchSite`, `Success`, `CauseFailure` FROM tbl_mission WHERE ID = 4")
  if err != nil {
    panic(err)
  }
  for res.Next() {
    var Moon3 Mission
    err = res.Scan(&Moon3.Name,
                   &Moon3.StartData,
                   &Moon3.FinishData,
                   &Moon3.DayDuration,
                   &Moon3.Target,
                   &Moon3.CarrierRocket,
                   //&Moon3.Crew,
                   &Moon3.CustomerCountry,
                   &Moon3.LaunchSite,
                   &Moon3.Success,
                   &Moon3.CauseFailure)
    if err != nil {
      panic(err)
    }
    Moon3.Crew = Moon3_crew
    Moon3.parsePage(w_page, r)
    //fmt.Fprintf(w_page, Moon3.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

func Moon19_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `member1`, `member2`, `member3`, `member4`, `member5`, `member6`, `member7`" +
    " FROM tbl_crew WHERE ID_Crew = 5")
  if err != nil {
    panic(err)
  }
  var Moon19_crew TCrew
  for res.Next() {
    err = res.Scan(&Moon19_crew.member1,
                   &Moon19_crew.member2,
                   &Moon19_crew.member3,
                   &Moon19_crew.member4,
                   &Moon19_crew.member5,
                   &Moon19_crew.member6,
                   &Moon19_crew.member7)
               }
  res, err = db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, " +
    " `CustomerCountry`, `LaunchSite`, `Success`, `CauseFailure` FROM tbl_mission WHERE ID = 5")
  if err != nil {
    panic(err)
  }
  for res.Next() {
    var Moon19 Mission
    err = res.Scan(&Moon19.Name,
                   &Moon19.StartData,
                   &Moon19.FinishData,
                   &Moon19.DayDuration,
                   &Moon19.Target,
                   &Moon19.CarrierRocket,
                   //&Moon19.Crew,
                   &Moon19.CustomerCountry,
                   &Moon19.LaunchSite,
                   &Moon19.Success,
                   &Moon19.CauseFailure)
    if err != nil {
      panic(err)
    }
    Moon19.Crew = Moon19_crew
    Moon19.parsePage(w_page, r)
    //fmt.Fprintf(w_page, Moon19.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

func Appolo11_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `member1`, `member2`, `member3`, `member4`, `member5`, `member6`, `member7`" +
    " FROM tbl_crew WHERE ID_Crew = 6")
  if err != nil {
    panic(err)
  }
  var Appolo11_crew TCrew
  for res.Next() {
    err = res.Scan(&Appolo11_crew.member1,
                   &Appolo11_crew.member2,
                   &Appolo11_crew.member3,
                   &Appolo11_crew.member4,
                   &Appolo11_crew.member5,
                   &Appolo11_crew.member6,
                   &Appolo11_crew.member7)
               }
  res, err = db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, " +
    " `CustomerCountry`, `LaunchSite`, `Success`, `CauseFailure` FROM tbl_mission WHERE ID = 6")
  if err != nil {
    panic(err)
  }
  for res.Next() {
    var Appolo11 Mission
    err = res.Scan(&Appolo11.Name,
                   &Appolo11.StartData,
                   &Appolo11.FinishData,
                   &Appolo11.DayDuration,
                   &Appolo11.Target,
                   &Appolo11.CarrierRocket,
                   //&Appolo11.Crew,
                   &Appolo11.CustomerCountry,
                   &Appolo11.LaunchSite,
                   &Appolo11.Success,
                   &Appolo11.CauseFailure)
    if err != nil {
      panic(err)
    }
    Appolo11.Crew = Appolo11_crew
    Appolo11.parsePage(w_page, r)
    //fmt.Fprintf(w_page, Appolo11.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

func MoonWalker2_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `member1`, `member2`, `member3`, `member4`, `member5`, `member6`, `member7`" +
    " FROM tbl_crew WHERE ID_Crew = 7")
  if err != nil {
    panic(err)
  }
  var MoonWalker2_crew TCrew
  for res.Next() {
    err = res.Scan(&MoonWalker2_crew.member1,
                   &MoonWalker2_crew.member2,
                   &MoonWalker2_crew.member3,
                   &MoonWalker2_crew.member4,
                   &MoonWalker2_crew.member5,
                   &MoonWalker2_crew.member6,
                   &MoonWalker2_crew.member7)
               }
  res, err = db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, " +
    " `CustomerCountry`, `LaunchSite`, `Success`, `CauseFailure` FROM tbl_mission WHERE ID = 7")
  if err != nil {
    panic(err)
  }
  for res.Next() {
    var MoonWalker2 Mission
    err = res.Scan(&MoonWalker2.Name,
                   &MoonWalker2.StartData,
                   &MoonWalker2.FinishData,
                   &MoonWalker2.DayDuration,
                   &MoonWalker2.Target,
                   &MoonWalker2.CarrierRocket,
                   //&MoonWalker2.Crew,
                   &MoonWalker2.CustomerCountry,
                   &MoonWalker2.LaunchSite,
                   &MoonWalker2.Success,
                   &MoonWalker2.CauseFailure)
    if err != nil {
      panic(err)
    }
    MoonWalker2.Crew = MoonWalker2_crew
    MoonWalker2.parsePage(w_page, r)
    //fmt.Fprintf(w_page, MoonWalker2.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

func Voyager2_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `member1`, `member2`, `member3`, `member4`, `member5`, `member6`, `member7`" +
    " FROM tbl_crew WHERE ID_Crew = 8")
  if err != nil {
    panic(err)
  }
  var Voyager2_crew TCrew
  for res.Next() {
    err = res.Scan(&Voyager2_crew.member1,
                   &Voyager2_crew.member2,
                   &Voyager2_crew.member3,
                   &Voyager2_crew.member4,
                   &Voyager2_crew.member5,
                   &Voyager2_crew.member6,
                   &Voyager2_crew.member7)
               }
  res, err = db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, " +
    " `CustomerCountry`, `LaunchSite`, `Success`, `CauseFailure` FROM tbl_mission WHERE ID = 8")
  if err != nil {
    panic(err)
  }
  for res.Next() {
    var Voyager2 Mission
    err = res.Scan(&Voyager2.Name,
                   &Voyager2.StartData,
                   &Voyager2.FinishData,
                   &Voyager2.DayDuration,
                   &Voyager2.Target,
                   &Voyager2.CarrierRocket,
                   //&Voyager2.Crew,
                   &Voyager2.CustomerCountry,
                   &Voyager2.LaunchSite,
                   &Voyager2.Success,
                   &Voyager2.CauseFailure)
    if err != nil {
      panic(err)
    }
    Voyager2.Crew = Voyager2_crew
    Voyager2.parsePage(w_page, r)
    //fmt.Fprintf(w_page, Voyager2.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

func Akatcuki_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `member1`, `member2`, `member3`, `member4`, `member5`, `member6`, `member7`" +
    " FROM tbl_crew WHERE ID_Crew = 9")
  if err != nil {
    panic(err)
  }
  var Akatcuki_crew TCrew
  for res.Next() {
    err = res.Scan(&Akatcuki_crew.member1,
                   &Akatcuki_crew.member2,
                   &Akatcuki_crew.member3,
                   &Akatcuki_crew.member4,
                   &Akatcuki_crew.member5,
                   &Akatcuki_crew.member6,
                   &Akatcuki_crew.member7)
               }
  res, err = db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, " +
    " `CustomerCountry`, `LaunchSite`, `Success`, `CauseFailure` FROM tbl_mission WHERE ID = 9")
  if err != nil {
    panic(err)
  }
  for res.Next() {
    var Akatcuki Mission
    err = res.Scan(&Akatcuki.Name,
                   &Akatcuki.StartData,
                   &Akatcuki.FinishData,
                   &Akatcuki.DayDuration,
                   &Akatcuki.Target,
                   &Akatcuki.CarrierRocket,
                   //&Akatcuki.Crew,
                   &Akatcuki.CustomerCountry,
                   &Akatcuki.LaunchSite,
                   &Akatcuki.Success,
                   &Akatcuki.CauseFailure)
    if err != nil {
      panic(err)
    }
    Akatcuki.Crew = Akatcuki_crew
    Akatcuki.parsePage(w_page, r)
    //fmt.Fprintf(w_page, Akatcuki.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

func NewHorizons_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `member1`, `member2`, `member3`, `member4`, `member5`, `member6`, `member7`" +
    " FROM tbl_crew WHERE ID_Crew = 10")
  if err != nil {
    panic(err)
  }
  var NewHorizons_crew TCrew
  for res.Next() {
    err = res.Scan(&NewHorizons_crew.member1,
                   &NewHorizons_crew.member2,
                   &NewHorizons_crew.member3,
                   &NewHorizons_crew.member4,
                   &NewHorizons_crew.member5,
                   &NewHorizons_crew.member6,
                   &NewHorizons_crew.member7)
               }
  res, err = db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, " +
    " `CustomerCountry`, `LaunchSite`, `Success`, `CauseFailure` FROM tbl_mission WHERE ID = 10")
  if err != nil {
    panic(err)
  }
  for res.Next() {
    var NewHorizons Mission
    err = res.Scan(&NewHorizons.Name,
                   &NewHorizons.StartData,
                   &NewHorizons.FinishData,
                   &NewHorizons.DayDuration,
                   &NewHorizons.Target,
                   &NewHorizons.CarrierRocket,
                   //&NewHorizons.Crew,
                   &NewHorizons.CustomerCountry,
                   &NewHorizons.LaunchSite,
                   &NewHorizons.Success,
                   &NewHorizons.CauseFailure)
    if err != nil {
      panic(err)
    }
    NewHorizons.Crew = NewHorizons_crew
    NewHorizons.parsePage(w_page, r)
    //fmt.Fprintf(w_page, NewHorizons.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

func Mars2020_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `member1`, `member2`, `member3`, `member4`, `member5`, `member6`, `member7`" +
    " FROM tbl_crew WHERE ID_Crew = 11")
  if err != nil {
    panic(err)
  }
  var Mars2020_crew TCrew
  for res.Next() {
    err = res.Scan(&Mars2020_crew.member1,
                   &Mars2020_crew.member2,
                   &Mars2020_crew.member3,
                   &Mars2020_crew.member4,
                   &Mars2020_crew.member5,
                   &Mars2020_crew.member6,
                   &Mars2020_crew.member7)
               }
  res, err = db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, " +
    " `CustomerCountry`, `LaunchSite`, `Success`, `CauseFailure` FROM tbl_mission WHERE ID = 11")
  if err != nil {
    panic(err)
  }
  for res.Next() {
    var Mars2020 Mission
    err = res.Scan(&Mars2020.Name,
                   &Mars2020.StartData,
                   &Mars2020.FinishData,
                   &Mars2020.DayDuration,
                   &Mars2020.Target,
                   &Mars2020.CarrierRocket,
                   //&Mars2020.Crew,
                   &Mars2020.CustomerCountry,
                   &Mars2020.LaunchSite,
                   &Mars2020.Success,
                   &Mars2020.CauseFailure)
    if err != nil {
      panic(err)
    }
    Mars2020.Crew = Mars2020_crew
    Mars2020.parsePage(w_page, r)
    //fmt.Fprintf(w_page, Mars2020.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}
