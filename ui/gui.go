package ui

import (
	"fmt"
	"log"
	reservio "restaurant-project/reservation-system"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Gui() {
	//dbh := reservio.NewTempStorage(10, 10)
	dbh, _ := reservio.NewSqliteStorage()
	// dbh,_ := reservio.NewComposedStorage()
	rs := reservio.NewReservationSystem(dbh)
	a := app.New()
	w := a.NewWindow("Restaurant reservation system")
	var loggedName = new(string)
	//w.Resize(fyne.NewSize(500, 500))
	options := []string{"1", "2", "3", "4", "5", "6", "7"}
	hrOptions := []string{"10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}
	prOptions := []string{"1", "2", "3", "4", "5", "6"}
	var day, hour, duration, persons int
	slDay := widget.NewSelect(options, func(value string) {
		day, _ = strconv.Atoi(value)
	})
	slHour := widget.NewSelect(hrOptions, func(value string) {
		hour, _ = strconv.Atoi(value)
	})
	slDuration := widget.NewSelect(options, func(value string) {
		duration, _ = strconv.Atoi(value)
	})
	slPersons := widget.NewSelect(prOptions, func(value string) {
		persons, _ = strconv.Atoi(value)
	})
	btMkRes := widget.NewButton("Make reservation", func() {
		err := rs.MakeReservation(*loggedName, day, hour, duration, persons)
		if err != nil {
			log.Print(err)
			dlg := dialog.NewError(err, w)
			dlg.Show()
			return
		}
		ttl := "Reservation created"
		msg := fmt.Sprintf("Reservation at day %v, from %v hr for %v hrs \nand for %v persons was successfully created!",
			day, hour, duration, persons)
		dlg := dialog.NewInformation(ttl, msg, w)
		dlg.Show()

	})
	txtDay := widget.NewLabel("Day")
	txtHour := widget.NewLabel("Hour")
	txtDur := widget.NewLabel("Duration")
	txtPersons := widget.NewLabel("Persons")
	resFrmCnt := container.New(layout.NewFormLayout(), txtDay, slDay, txtHour, slHour,
		txtDur, slDuration, txtPersons, slPersons)
	resCnt := container.New(layout.NewVBoxLayout(), resFrmCnt, btMkRes)

	name := widget.NewEntry()
	pw := widget.NewPasswordEntry()
	btLgn := widget.NewButton("Login", func() {
		err := rs.Login(name.Text, pw.Text)
		if err != nil {
			log.Print(err)
			dlg := dialog.NewError(err, w)
			dlg.Show()
			return
		}
		w.SetTitle(fmt.Sprintf("Restaurant reservation system - logged customer: %v", name.Text))
		*loggedName = name.Text
		name.SetText("")
		pw.SetText("")
	})
	btRgt := widget.NewButton("Register", func() {
		err := rs.Register(name.Text, pw.Text)
		if err != nil {
			log.Print(err)
			dlg := dialog.NewError(err, w)
			dlg.Show()
			return
		}
	})
	btLgt := widget.NewButton("Logout", func() {
		err := rs.Logout(*loggedName)
		if err != nil {
			log.Print(err)
		}
		w.SetTitle("Restaurant reservation system")
		*loggedName = ""
	})
	cstCnt := container.New(layout.NewVBoxLayout(), name, pw, btLgn, btRgt, btLgt)

	btGt := widget.NewButton("Show reservations", func() {
		rss, err := rs.GetReservations(*loggedName)
		if err != nil {
			log.Print(err)
			dlg := dialog.NewError(err, w)
			dlg.Show()
			return
		}
		dlg := dialog.NewInformation("Your reservations", string(*rss), w)
		dlg.Show()
	})
	tables := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}
	var dayCncl, table int
	slDayCncl := widget.NewSelect(options, func(value string) {
		dayCncl, _ = strconv.Atoi(value)
	})
	slTable := widget.NewSelect(tables, func(value string) {
		table, _ = strconv.Atoi(value)
	})
	txtDayCncl := widget.NewLabel("Day")
	txtTable := widget.NewLabel("Table")
	btCncl := widget.NewButton("Cancel reservation", func() {
		err := rs.CancelReservation(*loggedName, dayCncl, table)
		if err != nil {
			log.Print(err)
			dlg := dialog.NewError(err, w)
			dlg.Show()
			return
		}
		dlg := dialog.NewInformation("Cancellation complete", fmt.Sprintf("Your reservation at day %v at table %v was cancelled",
			dayCncl, table), w)
		dlg.Show()
	})
	cnclCnt := container.New(layout.NewFormLayout(), txtDayCncl, slDayCncl, txtTable, slTable)
	thrdCnt := container.New(layout.NewVBoxLayout(), btGt, cnclCnt, btCncl)

	mainCnt := container.New(layout.NewAdaptiveGridLayout(3), cstCnt, resCnt, thrdCnt)

	w.SetContent(mainCnt)
	w.ShowAndRun()
}
