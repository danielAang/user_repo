package main

func (a *App) MountRoutes() {
	a.Router.HandleFunc("/users/{id}", a.GetUserById).Methods("GET")
	a.Router.HandleFunc("/users", a.CreateUser).Methods("POST")
}
