package controllers

import (
  "github.com/sirupsen/logrus"
  "github.com/gin-gonic/gin"
  "github.com/gorilla/csrf"

  "github.com/charmixer/aapui/config"
  "github.com/charmixer/aapui/environment"
)

func ShowAuthorizations(env *environment.State, route environment.Route) gin.HandlerFunc {
  fn := func(c *gin.Context) {

    log := c.MustGet(environment.LogKey).(*logrus.Entry)
    log = log.WithFields(logrus.Fields{
      "func": "ShowAuthorizations",
    })

    c.HTML(200, "authorizations.html", gin.H{
      "title": "Authorizations",
      csrf.TemplateTag: csrf.TemplateField(c.Request),
      "links": []map[string]string{
        {"href": "/public/css/dashboard.css"},
      },
      "idpUiUrl": config.GetString("idpui.public.url"),
      "aapUiUrl": config.GetString("aapui.public.url"),
    })
  }
  return gin.HandlerFunc(fn)
}
