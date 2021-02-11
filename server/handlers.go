package server

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"os"
)

func getUser(c echo.Context) error {
	cc := c.(*ServerContext)
	// Get request param
	userID := c.Param("id")

	// Open JSON file
	jsonFile, err := os.Open(fileDB)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer jsonFile.Close()

	// Read JSON file
	body, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//...
	var users Users
	json.Unmarshal(body, &users)
	for _, u := range users.Users {
		if u.ID == userID {
			return c.JSON(http.StatusOK, u)
		}
	}
	return c.JSON(http.StatusNotFound, userID)
}

func getAllUsers(c echo.Context) error {
	// Open JSON file
	jsonFile, err := os.Open(fileDB)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer jsonFile.Close()

	body, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// Also you can output full body at once
	var users Users
	json.Unmarshal(body, &users)
	var response []byte
	response, err = json.Marshal(users.Users)

	return c.JSON(http.StatusOK, string(response))
}

func remove(slice []User, i int) []User {
	return append(slice[:i], slice[i+1:]...)
}

func deleteUser(c echo.Context) error {
	// Get request param
	userID := c.Param("id")

	// Open JSON file
	jsonFile, err := os.Open(fileDB)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer jsonFile.Close()

	// Parse JSON file
	body, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var users Users
	json.Unmarshal(body, &users)

	// Find and delete user
	for i, u := range users.Users {
		if u.ID == userID {
			users.Users = remove(users.Users, i)
			newBody, err := json.Marshal(users)
			if err != nil {
				fmt.Println(err)
				return err
			}
			err = ioutil.WriteFile(fileDB, newBody, 0644)
			return c.JSON(http.StatusOK, userID)
		}
	}

	return c.JSON(http.StatusNotFound, userID)
}

func addUser(c echo.Context) error {
	// Get request param
	user := &User{}
	if err := c.Bind(user); err != nil {
		fmt.Println(err)
		return err
	}

	// Open JSON file
	jsonFile, err := os.Open(fileDB)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer jsonFile.Close()

	// Parse JSON file
	body, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var users Users
	json.Unmarshal(body, &users)

	// Check for unic
	if !isUniqueID(users.Users, user.ID) {
		return c.JSON(http.StatusBadRequest, user)
	}

	// Add new user
	users.Users = append(users.Users, *user)
	newBody, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ioutil.WriteFile(fileDB, newBody, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func updateUser(c echo.Context) error {
	// Get request param
	user := &User{}
	if err := c.Bind(user); err != nil {
		fmt.Println(err)
		return err
	}

	// Open JSON file
	jsonFile, err := os.Open(fileDB)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer jsonFile.Close()

	// Parse JSON file
	body, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var users Users
	json.Unmarshal(body, &users)

	for i, u := range users.Users {
		if u.ID == user.ID {
			users.Users[i] = *user
			newBody, err := json.Marshal(users)
			if err != nil {
				fmt.Println(err)
				return err
			}
			err = ioutil.WriteFile(fileDB, newBody, 0644)
			if err != nil {
				fmt.Println(err)
				return err
			}
			return c.JSON(http.StatusOK, user)
		}
	}

	return c.JSON(http.StatusNotFound, user)
}

func isUniqueID(users []User, id string) bool {
	for _, v := range users {
		if v.ID == id {
			return false
		}
	}
	return true
}
