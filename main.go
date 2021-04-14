package main
import ("fmt"   //доступ функциям (напр: вывод на сайт,терминал)
        //"net/http"   //показывать инф-ю пользователю, отслежевать его действия
        //"html/template" //для работы с шаблонами html, графическая чать
        "database/sql"

	      _ "github.com/go-sql-driver/mysql"
        //_ "pkg/mod/github.com/go-sql-driver/mysql"
        //_ "pkg/mod/github.com/go-sql-driver/mysql@v1.6.0"
)

type Mission struct {
  Name string
  StartData string
  FinishData string
  DayDuration uint16 //uint16-хранит целое неотрицательное число
  Target string
  CarrierRocket string
  Crew []string
  CustomerCountry string
  LaunchSite string
  Success bool
  CauseFailure string
}

func (this Mission) getAllInfo() string { //вывод содержимого объекта
  //if (this.Success == true)
  var returnStr string = "\nНазвание миссии: %s. \nДата запуска: %s." +
                     "\nМиссия завершена: %s. \nПродолжительность: %d дней." +
                     "\nОсновная цель миссии: %s. \nРакета-носитель: %s." +
                     "\nКоманда/Аппараты: %s. \nСтрана: %s." +
                     "\nКосмодром: %s. \nУспех миссии: %b." +
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
  db, err := sql.Open("mysql", "mysql:mysql@tcp(localhost:3306)/golangdb")
  if err != nil {
    panic(err)
  }
  defer db.Close()

  //Установка данных
  // mode, err := db.Query("alter database golangdb set read_write")
  // if err != nil {
  //   panic(err)
  // }
  insert, err := db.Query("INSERT INTO `missions` (`ID`, `Name`, `StartData`, `FinishData`, `DayDuration`, `Target`, `CarrierRocket`, `Crew`, `CustomerCountry`, `LaunchSite`, `Success`, `CauseFailure`) VALUES('1', 'Alex', '25 aug', '26 feb', '45', 'Cosmos', 'Atlas', '3', 'USA', 'Beach', '1', '-')")
  //insert, err := db.Query("INSERT INTO `crewtable` (`member1`, `member2`, `member3`, `member4`, `member5`, `member6`, `member7`, `member8`, `member9`) VALUES('Bob', 'Alex', 'Frank', 'Prev', 'Ing', '-', '-', '-', '-')")
  if err != nil {
    panic(err)
  }
  //defer mode.Close()
  defer insert.Close()

  fmt.Println("Подключено к MySql")
  //handleRequest() //отслеживаем url, запускаем сервер
}
/*
func handleRequest() {
  http.HandleFunc("/", home_page) //arg1- отслеживание перехода по url; arg2-метод при arg1; /-главная страница (/about)
  http.HandleFunc("/faq/", faq_page)
  http.HandleFunc("/Pioner5/", Pioner5_page)
  http.HandleFunc("/HeliosB/", HeliosB_page)
  http.HandleFunc("/PionerE/", PionerE_page)
  http.HandleFunc("/Moon3/", Moon3_page)
  http.HandleFunc("/Moon19/", Moon19_page)
  http.HandleFunc("/Appolo11/", Appolo11_page)
  http.HandleFunc("/MoonWalker2/", MoonWalker2_page)
  http.HandleFunc("/Voyager2/", Voyager2_page)
  http.HandleFunc("/Akatciki/", Akatciki_page)
  http.HandleFunc("/NewHorizons/", NewHorizons_page)
  http.HandleFunc("/Mars2020/", Mars2020_page)

  http.ListenAndServe(":9080", nil) //запуск локально сервера на порту:8080(любой);...
  //...arg1-порт по чтению сервера(любой свободный на ПК); arg2-настройки для сервера; nil-NULL
}

func home_page(w_page http.ResponseWriter, r *http.Request) { // arg2(r)-запрос для проверки передачи данных
  templ, _ := template.ParseFiles("templates/home_page.html") //v1-хранит шаблон, v2-обработка ошибок
  templ.Execute(w_page, home_page) //отображение шаблона; arg2-объект Mission
}

func faq_page(w_page http.ResponseWriter, r *http.Request) {
  // fmt.Fprintf(w_page, `<h1>Main Text</h1>
  // <b>Main Text</b>`)
}

func Pioner5_page(w_page http.ResponseWriter, r *http.Request)  {
  Pioner5 := Mission{"Пионер-5", "11 марта 1960", "30 апреля 1960", 50, "Исследование солнечных частиц и космического пространства.",
                     "Thor Able IV 219", []string{"Bob","Alex","Kate"}, "США", "Мыс Канаверал", true, "-"} //создание объекта
  fmt.Fprintf(w_page, Pioner5.getAllInfo()) //вывод форматируемой строки; arg1-куда вывод; arg2-что выводим;
}

func HeliosB_page(w_page http.ResponseWriter, r *http.Request)  {
  HeliosB := Mission{"Helios-B", "15 января 1976", "23 декабря 1979", 1438, "Приближение к Солнцу на 0,291 а.е.",
                     "Titan IIIE/Centaur", []string{"Helios spacecraft","500kb Computer"}, "Германия/США", "Мыс Канаверал", true, "-"}
  fmt.Fprintf(w_page, HeliosB.getAllInfo())
}

func PionerE_page(w_page http.ResponseWriter, r *http.Request)  {
  PionerE := Mission{"Пионер-E", "27 августа 1969", "-", 0, "Изучение солнечной плазмы и ветра, физики частиц.",
                     "Thor-Delta L", []string{"Module: Pioner E"}, "США", "Мыс Канаверал", true, "Взрыв ракеты-носителя через 8 минут после старта"}
  fmt.Fprintf(w_page, PionerE.getAllInfo())
}

func Moon3_page(w_page http.ResponseWriter, r *http.Request)  {
  Moon3 := Mission{"Луна-3", "4 октября 1959", "20 апреля 1960", 200, "Снимки обратной стороны Луны.",
                     "", []string{"Module E-2A"}, "СССР", "Байконур", true, "-"}
  fmt.Fprintf(w_page, Moon3.getAllInfo())
}

func Moon19_page(w_page http.ResponseWriter, r *http.Request)  {
  Moon19 := Mission{"Луна-19", "28 сентября 1971", "1 ноября 1972", 388, "Картографирование лунной поверхности.",
                     "Проток-К/Блок Д", []string{"Станция У-8ЛС №202"}, "СССР", "Байконур Пл. 81/24", false, "Отказ системы управления-> Не выход на нужную орбиту, потеря связи."}
  fmt.Fprintf(w_page, Moon19.getAllInfo())
}

func Appolo11_page(w_page http.ResponseWriter, r *http.Request)  {
  Appolo11 := Mission{"Аполон-11", "16 июля 1969", "24 июля 1969", 8, "Высадка на поверхность Луны.",
                     "Saturn V (SA-506)", []string{"Neil Alden Armstrong", "Michael Collins", "Edwin Eugene Aldrin, Jr."}, "США", "John F. Kennedy Space Center (Launch Complex 39A)", true, "-"}
  fmt.Fprintf(w_page, Appolo11.getAllInfo())
}

func MoonWalker2_page(w_page http.ResponseWriter, r *http.Request)  {
  MoonWalker2 := Mission{"Луноход-2", "15 января 1973", "4 июня 1973", 50, "Исследование поверхности Луны.",
                     "Проток-К/Блок Д", []string{"Луноход-2"}, "СССР", "Байконур Пл. 81/23", true, "-"}
  fmt.Fprintf(w_page, MoonWalker2.getAllInfo())
}

func Voyager2_page(w_page http.ResponseWriter, r *http.Request)  {
  Voyager2 := Mission{"Вояджер-2", "20 августа 1977", "активен", 15940, "Исследование дальних планет Солнечной системы.",
                     "Titan IIIE/ Центавр", []string{""}, "США", "Мыс Канаверал", true, "-"}
  fmt.Fprintf(w_page, Voyager2.getAllInfo())
}

func Akatciki_page(w_page http.ResponseWriter, r *http.Request)  {
  Akatciki := Mission{"Акацуки (PLANET-C)", "20 мая 2010", "активна", 3980, "Изучение Венеры.",
                     "H-IIA202 (F17)", []string{"IR1(Инфракрасная камера)", "IR2", "LIR(Болометр)", "UVI(Ультрафиолетовая камера)", "LAC(Детектор молний)", "USO (Ультра-стабильный генератор X-диапазона)"}, "Япония", "Tanegashima Space Center", true, "-"}
  fmt.Fprintf(w_page, Akatciki.getAllInfo())
}

func NewHorizons_page(w_page http.ResponseWriter, r *http.Request)  {
  NewHorizons := Mission{"Новые Горизонты", "19 января 2006", "активна", 5562, "Изучение Плутона и его естественного спутника Харона.",
                     "Atlas V", []string{"Alice(Спектрометр)","Ralph(Фотокамера)","LORRI(Камера для детальной съемки)", "SWAP", "PEPSSI", "REX", "VB-SDC"}, "США", "Мыс Канаверал", true, "-"}
  fmt.Fprintf(w_page, NewHorizons.getAllInfo())
}

func Mars2020_page(w_page http.ResponseWriter, r *http.Request)  {
  Mars2020 := Mission{"Марс-2020", "30 июля 2020", "активна", 0, "Исследование поверхности Марса. Первый управляемый полет на другой планете.",
                     "Atlas V (541)", []string{"Rover: Perseverance","Coptet: Ingenuity"}, "США", "Мыс Канаверал", true, "-"}
  fmt.Fprintf(w_page, Mars2020.getAllInfo())
}
*/
