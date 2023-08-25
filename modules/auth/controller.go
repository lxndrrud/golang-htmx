package auth

import (
	"log"
	"net/http"
	"strconv"
)

type authController struct {
	authService *authService
}

func newAuthController(authService *authService) *authController {
	return &authController{authService}
}

func (ac authController) LoginPage(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("id-user"); err != http.ErrNoCookie {
		http.Redirect(w, r, "/products/", http.StatusPermanentRedirect)
		return
	}
	template, err := ac.authService.LoginPage()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}
	err = template.Execute(w, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}
}

func (ac authController) LoginAction(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}
	login := r.Form.Get("login")
	password := r.Form.Get("password")
	user, err := ac.authService.LoginAction(login, password)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}
	if user == nil {
		http.Redirect(w, r, "/products/", http.StatusMovedPermanently)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "login",
		Value: user.Login,
		Path:  "/",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "id-user",
		Value: strconv.FormatUint(user.Id, 10),
		Path:  "/",
	})
	r.Header.Add("Content-Type", "text/html")
	http.Redirect(w, r, "/products/", http.StatusMovedPermanently)
}

func (ac authController) LogoutAction(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("login")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/products/", http.StatusPermanentRedirect)
		return
	}
	_, err = r.Cookie("id-user")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/products/", http.StatusPermanentRedirect)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:   "login",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "id-user",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	r.Header.Add("Content-Type", "text/html")
	http.Redirect(w, r, "/products/", http.StatusMovedPermanently)
}
