package api

import (
	"eshaanagg/lfx/database/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllOrgs(c *gin.Context) {
	client := handlers.New()
	defer client.Close()

	year := c.Param("year")
	term := c.Param("term")

	// if year== nil && term==nil {
	// 	c.IndentedJSON(http.StatusOK, client.GetAllParentOrgs())
	// }
	// else if year==nil {
	// 	c.IndentedJSON(http.StatusOK, client.GetAllParentOrgsByTerm())
	// }
	// else if(term==nil){
	// 	c.IndentedJSON(http.StatusOK, client.GetAllParentOrgsByYear())
	// }
	// else{
	// 	c.IndentedJSON(http.StatusOK, client.GetAllParentOrgs())
	// }

	if year=="" && term=="" {
        c.IndentedJSON(http.StatusOK, client.GetAllParentOrgs())
    } else if year=="" {
        c.IndentedJSON(http.StatusOK, client.GetAllParentOrgsByTerm())
    } else if term==""{
    	c.IndentedJSON(http.StatusOK, client.GetAllParentOrgsByYear())
    }else{
		c.IndentedJSON(http.StatusOK, client.GetAllParentOrgsByTermYear())
	}


	
}
