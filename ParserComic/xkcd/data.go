package xkcd

import(
	"fmt"
)

const Link = "https://xkcd.com/"

type DataComic struct{
	Year string `json: "year"`
	Month string `json: "month"`
	Day string `json: "day"`
	Num int `json: "num"`
	Title string `json: "title"`
	Transcript string `json: "transcript"`
	Img string `json: "img"`

}

func (d *DataComic) Print(){
	fmt.Printf("%v.%v.%v - Дата выпуска комикса\n",d.Day,d.Month,d.Year)
	fmt.Printf("%v - Номер коммикса\n",d.Num)
	fmt.Printf("%v - Название коммикса\n",d.Title)
	fmt.Printf("Основное содержание: \n %v\n",d.Transcript)
	fmt.Printf("Cсылка на картинку:\n%v\n",d.Img)
}