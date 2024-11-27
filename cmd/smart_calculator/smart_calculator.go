package main

import (
	"encoding/json"
	"fmt"
	calcadapter "leftrana/smartcalc/pkg/calcadapter"
	"net/http"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

type WeatherInfo struct {
	List []WeatherListItem `json:"list"`
}

type WeatherListItem struct {
	Dt      int           `json:"dt"`
	Main    WeatherMain   `json:"main"`
	Weather []WeatherType `json:"weather"`
}

type WeatherMain struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	Humidity  int     `json:"humidity"`
}

type WeatherType struct {
	Icon string `json:"icon"`
}

const openWeatherMapApiKey = "jklsjdfgoijlsdkajfijlgkjdf"

func main() {
	str := "5+5*2"
	got, err := calcadapter.Calculate(str)
	if err != nil || got != 15 {
		fmt.Printf("5+5*2 = %f; want 15 ", got)
	}
	fmt.Println("hello", got)
	var w WeatherInfo
	err = getWeatherForecast(w)
	if err != nil {
		fmt.Println("no wether")
	}
	setupUi(w)
}

func getWeatherForecast(result interface{}) error {
	var url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?q=Voronezh&cnt=4&units=metric&appid=%s", openWeatherMapApiKey) // инициализируйте со своим ключом
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(result)
}

func setupUi(weatherInfo WeatherInfo) {
	app := app.New()

	w := app.NewWindow("Программа для просмотра погоды")

	var vBox = widget.NewVBox() // создаем новый контейнер с вертикальным порядком дочерних виджетов

	for i := 0; i < len(weatherInfo.List); i++ {
		var weatherForDay = weatherInfo.List[i] // погода дня для текущей итерации
		var weatherMainGroup = widget.NewVBox(
			widget.NewLabel(fmt.Sprintf("Температура: %.2f °C", weatherForDay.Main.Temp)),
			widget.NewLabel(fmt.Sprintf("Ощущается как: %.2f °C", weatherForDay.Main.FeelsLike)),
			widget.NewLabel(fmt.Sprintf("Влажность: %d%%", weatherForDay.Main.Humidity)),
		) // отображаем 3 лейбла один под другим

		var weatherTypeGroup = widget.NewVBox()
		for weatherTypeI := 0; weatherTypeI < len(weatherForDay.Weather); weatherTypeI++ {
			var resource, _ = fyne.LoadResourceFromURLString(fmt.Sprintf("http://openweathermap.org/img/wn/%s.png", weatherForDay.Weather[weatherTypeI].Icon)) // создаем статический ресурс, содержащий иконку погоды
			var icon = widget.NewIcon(resource)
			weatherTypeGroup.Append(icon)
		}

		var time = time.Unix(int64(weatherInfo.List[i].Dt), 0).String()
		vBox.Append(widget.NewGroup(time))
		vBox.Append(widget.NewHBox(weatherMainGroup, weatherTypeGroup))
	}
	vBox.Append(widget.NewButton("Закрыть", func() {
		app.Quit()
	}))

	w.SetContent(vBox) //устанавливаем контент для окна приложения

	w.ShowAndRun() // запускаем сконфигурированное приложение
}
