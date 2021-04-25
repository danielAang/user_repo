package main

func (a *App) MountRoutes() {
	a.Router.HandleFunc("/users/{id}", a.Update).Methods("PUT")
	a.Router.HandleFunc("/users/{id}", a.GetUserById).Methods("GET")
	a.Router.HandleFunc("/users/{id}", a.DeleteById).Methods("DELETE")
	a.Router.HandleFunc("/users", a.GetUsers).Methods("GET")
	a.Router.HandleFunc("/users", a.CreateUser).Methods("POST")
}
