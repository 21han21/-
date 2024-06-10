package main

import (
	"fmt"
	"myapp/client/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.NewWithID("com.example.events")
	myApp.Settings().SetTheme(theme.DarkTheme()) // Устанавливаем темную тему
	myWindow := myApp.NewWindow("Events")

	// Задаем размер окна для мобильного устройства
	myWindow.Resize(fyne.NewSize(400, 600))

	events, err := ui.FetchEvents()
	if err != nil {
		myWindow.SetContent(widget.NewLabel("Error fetching events: " + err.Error()))
		myWindow.ShowAndRun()
		return
	}

	// Создаем список для отображения событий
	list := widget.NewList(
		func() int {
			return len(events)
		},
		func() fyne.CanvasObject {
			titleLabel := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
			descriptionLabel := widget.NewLabel("")
			locationLabel := widget.NewLabel("")
			datetimeLabel := widget.NewLabel("")
			registrationButton := widget.NewButton("Register", nil)

			return container.NewVBox(
				titleLabel,
				descriptionLabel,
				locationLabel,
				datetimeLabel,
				registrationButton,
			)
		},
		func(i int, item fyne.CanvasObject) {
			itemContainer := item.(*fyne.Container).Objects
			title := itemContainer[0].(*widget.Label)
			description := itemContainer[1].(*widget.Label)
			location := itemContainer[2].(*widget.Label)
			datetime := itemContainer[3].(*widget.Label)
			registrationButton := itemContainer[4].(*widget.Button)

			title.SetText(fmt.Sprintf("Title: %s", events[i].Title))
			description.SetText(fmt.Sprintf("Description: %s", events[i].Description))
			location.SetText(fmt.Sprintf("Location: %s", events[i].Location))
			datetime.SetText(fmt.Sprintf("Datetime: %s", events[i].Datetime))

			// Установка действия для кнопки регистрации
			registrationButton.OnTapped = func() {
				showRegistrationForm(myApp, myWindow)
			}
		})

	// Создаем изображение из файла
	img := canvas.NewImageFromFile("img.jpg.jpg")
	img.FillMode = canvas.ImageFillOriginal
	img.SetMinSize(fyne.NewSize(300, 200)) // Задаем размер изображения

	myWindow.SetContent(container.NewVBox(
		widget.NewLabelWithStyle("Events", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container.NewMax(list),
		img, // Добавляем изображение после списка
	))

	myWindow.ShowAndRun()
}

// Функция для отображения формы регистрации
func showRegistrationForm(app fyne.App, parent fyne.Window) {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Name")

	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Email")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: nameEntry},
			{Text: "Email", Widget: emailEntry},
		},
		OnSubmit: func() {
			fmt.Printf("Name: %s\nEmail: %s\n", nameEntry.Text, emailEntry.Text)
			dialog.ShowInformation("Success", "Registration Successful", parent)
		},
	}

	registrationWindow := app.NewWindow("Register for Event")
	registrationWindow.SetContent(container.NewVBox(
		widget.NewLabelWithStyle("Register for Event", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		form,
	))
	registrationWindow.Resize(fyne.NewSize(300, 200))
	registrationWindow.Show()
}
