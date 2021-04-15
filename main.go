package main
import ("fmt"   //доступ функциям (напр: вывод на сайт,терминал)
        "net/http"   //показывать инф-ю пользователю, отслежевать его действия
        "html/template" //для работы с шаблонами html, графическая чать
        "database/sql"

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
  Crew uint16 `json:"Crew"`
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
  db, _ = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbname))
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

func handleRequest() {
  http.HandleFunc("/", home_page) //arg1- отслеживание перехода по url; /-главная страница
  http.HandleFunc("/faq/", faq_page)
  http.HandleFunc("/Pioner5/", Pioner5_page)
  http.HandleFunc("/HeliosB/", HeliosB_page)
  http.HandleFunc("/PionerE/", PionerE_page)
  http.HandleFunc("/Moon3/", Moon3_page)
  http.HandleFunc("/Moon19/", Moon19_page)
  http.HandleFunc("/Appolo11/", Appolo11_page)
  http.HandleFunc("/MoonWalker2/", MoonWalker2_page)
  http.HandleFunc("/Voyager2/", Voyager2_page)
  http.HandleFunc("/Akatcuki/", Akatcuki_page)
  http.HandleFunc("/NewHorizons/", NewHorizons_page)
  http.HandleFunc("/Mars2020/", Mars2020_page)

  http.ListenAndServe(":8080", nil) //запуск локально сервера на порту:9080;...
  //...arg1-порт по чтению сервера(любой свободный на ПК); arg2-настройки для сервера; nil-NULL
}

func home_page(w_page http.ResponseWriter, r *http.Request) { // arg2(r)-запрос для проверки передачи данных
  tmpl, err := template.ParseFiles("templates/home_page.html") //v1-хранит шаблон, v2-обработка ошибок
  if err != nil {
    panic(err)
  }
  tmpl.Execute(w_page, home_page) //отображение шаблона;
}

func faq_page(w_page http.ResponseWriter, r *http.Request) {
  // fmt.Fprintf(w_page, `<h1>Main Text</h1>
  // <b>Main Text</b>`)
}

func Pioner5_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, `ID_Crew`," +
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
                   &Pioner5.Crew,
                   &Pioner5.CustomerCountry,
                   &Pioner5.LaunchSite,
                   &Pioner5.Success,
                   &Pioner5.CauseFailure)
    if err != nil {
      panic(err)
    }
    fmt.Fprintf(w_page, Pioner5.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

func HeliosB_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, `ID_Crew`," +
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
                   &HeliosB.Crew,
                   &HeliosB.CustomerCountry,
                   &HeliosB.LaunchSite,
                   &HeliosB.Success,
                   &HeliosB.CauseFailure)
    if err != nil {
      panic(err)
    }
    fmt.Fprintf(w_page, HeliosB.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

func PionerE_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, `ID_Crew`," +
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
                   &PionerE.Crew,
                   &PionerE.CustomerCountry,
                   &PionerE.LaunchSite,
                   &PionerE.Success,
                   &PionerE.CauseFailure)
    if err != nil {
      panic(err)
    }
    fmt.Fprintf(w_page, PionerE.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

func Moon3_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, `ID_Crew`," +
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
                   &Moon3.Crew,
                   &Moon3.CustomerCountry,
                   &Moon3.LaunchSite,
                   &Moon3.Success,
                   &Moon3.CauseFailure)
    if err != nil {
      panic(err)
    }
    fmt.Fprintf(w_page, Moon3.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

func Moon19_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, `ID_Crew`," +
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
                   &Moon19.Crew,
                   &Moon19.CustomerCountry,
                   &Moon19.LaunchSite,
                   &Moon19.Success,
                   &Moon19.CauseFailure)
    if err != nil {
      panic(err)
    }
    fmt.Fprintf(w_page, Moon19.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

func Appolo11_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, `ID_Crew`," +
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
                   &Appolo11.Crew,
                   &Appolo11.CustomerCountry,
                   &Appolo11.LaunchSite,
                   &Appolo11.Success,
                   &Appolo11.CauseFailure)
    if err != nil {
      panic(err)
    }
    fmt.Fprintf(w_page, Appolo11.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

func MoonWalker2_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, `ID_Crew`," +
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
                   &MoonWalker2.Crew,
                   &MoonWalker2.CustomerCountry,
                   &MoonWalker2.LaunchSite,
                   &MoonWalker2.Success,
                   &MoonWalker2.CauseFailure)
    if err != nil {
      panic(err)
    }
    fmt.Fprintf(w_page, MoonWalker2.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

func Voyager2_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, `ID_Crew`," +
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
                   &Voyager2.Crew,
                   &Voyager2.CustomerCountry,
                   &Voyager2.LaunchSite,
                   &Voyager2.Success,
                   &Voyager2.CauseFailure)
    if err != nil {
      panic(err)
    }
    fmt.Fprintf(w_page, Voyager2.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

func Akatcuki_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, `ID_Crew`," +
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
                   &Akatcuki.Crew,
                   &Akatcuki.CustomerCountry,
                   &Akatcuki.LaunchSite,
                   &Akatcuki.Success,
                   &Akatcuki.CauseFailure)
    if err != nil {
      panic(err)
    }
    fmt.Fprintf(w_page, Akatcuki.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

func NewHorizons_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, `ID_Crew`," +
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
                   &NewHorizons.Crew,
                   &NewHorizons.CustomerCountry,
                   &NewHorizons.LaunchSite,
                   &NewHorizons.Success,
                   &NewHorizons.CauseFailure)
    if err != nil {
      panic(err)
    }
    fmt.Fprintf(w_page, NewHorizons.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}

func Mars2020_page(w_page http.ResponseWriter, r *http.Request)  {
  res, err := db.Query("SELECT `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, `ID_Crew`," +
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
                   &Mars2020.Crew,
                   &Mars2020.CustomerCountry,
                   &Mars2020.LaunchSite,
                   &Mars2020.Success,
                   &Mars2020.CauseFailure)
    if err != nil {
      panic(err)
    }
    fmt.Fprintf(w_page, Mars2020.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
  }
}
