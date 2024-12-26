package species

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Species struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

var db *sql.DB

// InitDB initializes the database connection for species
func InitDB(database *sql.DB) {
	if database == nil {
		log.Fatal("Database connection is nil from species")
	}
	db = database
}

// Handler สำหรับ GET /species
func GetAllSpecies(c echo.Context) error {
	if db == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database not initialized"})
	}

	rows, err := db.Query("SELECT id, name, description, location FROM species")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to query database"})
	}
	defer rows.Close()

	var speciesList []Species
	for rows.Next() {
		var species Species
		if err := rows.Scan(&species.Id, &species.Name, &species.Description, &species.Location); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to scan row"})
		}
		speciesList = append(speciesList, species)
	}

	return c.JSON(http.StatusOK, speciesList)
}

// Handler สำหรับ GET /species/:id
func GetSpecies(c echo.Context) error {
	if db == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database not initialized"})
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var species Species
	err = db.QueryRow("SELECT id, name, description, location FROM species WHERE id = $1", id).Scan(
		&species.Id, &species.Name, &species.Description, &species.Location,
	)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Species not found"})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to query database"})
	}

	return c.JSON(http.StatusOK, species)
}
