package rabbitqueue

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/rohit123sinha456/digitalSignage/config"
)

func sendData(urlstring string, databyte []byte) error {
	uri := urlstring
	data := databyte
	req, err := http.NewRequest(http.MethodPut, uri, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func postData(urlstring string, databyte []byte) error {
	uri := urlstring
	data := databyte
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
func deleteData(urlstring string, databyte []byte) error {
	uri := urlstring
	data := databyte
	req, err := http.NewRequest(http.MethodDelete, uri, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func CreateUser(username string) error {
	var rabbitadmin string = config.GetEnvbyKey("APPRABBITADMIN")
	var rabbitadminpwd string = config.GetEnvbyKey("APPRABBITADMINPASS")
	var rabbituri string = config.GetEnvbyKey("APPRABBITURL2")
	uri := strings.Join([]string{"http://", rabbitadmin, ":", rabbitadminpwd, "@", rabbituri, "api/users/", username}, "")
	data := []byte(`{"password":"password","tags":"none"}`)
	err := sendData(uri, data)
	fmt.Println("User CReated in Rabbit")
	fmt.Println(err)

	return err
}

func DeleteUser(username string) error {
	var rabbitadmin string = config.GetEnvbyKey("APPRABBITADMIN")
	var rabbitadminpwd string = config.GetEnvbyKey("APPRABBITADMINPASS")
	var rabbituri string = config.GetEnvbyKey("APPRABBITURL2")
	uri := strings.Join([]string{"http://", rabbitadmin, ":", rabbitadminpwd, "@", rabbituri, "api/users/", username}, "")
	data := []byte(`{"password":"password","tags":"none"}`)
	err := deleteData(uri, data)
	fmt.Println("User Deleted in Rabbit")
	fmt.Println(err)

	return err
}

func SetUserandvHostPermisssion(username string, vhostname string) error {
	var rabbitadmin string = config.GetEnvbyKey("APPRABBITADMIN")
	var rabbitadminpwd string = config.GetEnvbyKey("APPRABBITADMINPASS")
	var rabbituri string = config.GetEnvbyKey("APPRABBITURL2")
	uri := strings.Join([]string{"http://", rabbitadmin, ":", rabbitadminpwd, "@", rabbituri, "api/permissions/", vhostname, "/", username}, "")
	data := []byte(`{"configure":".*","write":".*","read":".*"}`)
	err := sendData(uri, data)
	fmt.Println("User Premission set in Rabbit")
	fmt.Println(err)

	return err
}
func SetUserandvHostTopicPermisssion(username string, vhostname string) error {
	var rabbitadmin string = config.GetEnvbyKey("APPRABBITADMIN")
	var rabbitadminpwd string = config.GetEnvbyKey("APPRABBITADMINPASS")
	var rabbituri string = config.GetEnvbyKey("APPRABBITURL2")
	uri := strings.Join([]string{"http://", rabbitadmin, ":", rabbitadminpwd, "@", rabbituri, "api/topic-permissions/", vhostname, "/", username}, "")
	data := []byte(`{"exchange":"amq.topic","write":"^a","read":".*"}`)
	err := sendData(uri, data)
	fmt.Println("User VHOST topic premission set in rabbit in Rabbit")
	fmt.Println(err)
	return err
}

func CreatevHosts(vhostname string) error {
	var rabbitadmin string = config.GetEnvbyKey("APPRABBITADMIN")
	var rabbitadminpwd string = config.GetEnvbyKey("APPRABBITADMINPASS")
	var rabbituri string = config.GetEnvbyKey("APPRABBITURL2")
	data := []byte(`{"description":"virtual host description", "tags":"accounts,production"}`)
	uri := strings.Join([]string{"http://", rabbitadmin, ":", rabbitadminpwd, "@", rabbituri, "api/vhosts/", vhostname}, "")
	err := sendData(uri, data)
	fmt.Println("CreatevHosts")
	fmt.Println(err)
	return err
}

func DeletevHosts(vhostname string) error {
	var rabbitadmin string = config.GetEnvbyKey("APPRABBITADMIN")
	var rabbitadminpwd string = config.GetEnvbyKey("APPRABBITADMINPASS")
	var rabbituri string = config.GetEnvbyKey("APPRABBITURL2")
	data := []byte(`{"description":"virtual host description", "tags":"accounts,production"}`)
	uri := strings.Join([]string{"http://", rabbitadmin, ":", rabbitadminpwd, "@", rabbituri, "api/vhosts/", vhostname}, "")
	err := deleteData(uri, data)
	fmt.Println("DeletevHosts")
	fmt.Println(err)
	return err
}


func CreateExchange(vhostname string) error {
	var rabbitadmin string = config.GetEnvbyKey("APPRABBITADMIN")
	var rabbitadminpwd string = config.GetEnvbyKey("APPRABBITADMINPASS")
	var rabbituri string = config.GetEnvbyKey("APPRABBITURL2")
	data := []byte(`{"type":"direct","auto_delete":false,"durable":true,"internal":false}`)
	uri := strings.Join([]string{"http://", rabbitadmin, ":", rabbitadminpwd, "@", rabbituri, "api/exchanges/", vhostname, "/PLExchange"}, "")
	err := sendData(uri, data)
	fmt.Println("CreateExchange")
	fmt.Println(err)
	return err
}

func DeleteExchange(vhostname string) error {
	var rabbitadmin string = config.GetEnvbyKey("APPRABBITADMIN")
	var rabbitadminpwd string = config.GetEnvbyKey("APPRABBITADMINPASS")
	var rabbituri string = config.GetEnvbyKey("APPRABBITURL2")
	data := []byte(`{"type":"direct","auto_delete":false,"durable":true,"internal":false}`)
	uri := strings.Join([]string{"http://", rabbitadmin, ":", rabbitadminpwd, "@", rabbituri, "api/exchanges/", vhostname, "/PLExchange"}, "")
	err := deleteData(uri, data)
	fmt.Println("DeleteExchange")
	fmt.Println(err)
	return err
}

func CreateQueue(vhostname string) error {
	var rabbitadmin string = config.GetEnvbyKey("APPRABBITADMIN")
	var rabbitadminpwd string = config.GetEnvbyKey("APPRABBITADMINPASS")
	var rabbituri string = config.GetEnvbyKey("APPRABBITURL2")
	data := []byte(`{"auto_delete":false,"durable":true`)
	uri := strings.Join([]string{"http://", rabbitadmin, ":", rabbitadminpwd, "@", rabbituri, "api/queues/", vhostname, "/PLQueue"}, "")
	err := sendData(uri, data)
	fmt.Println("CreateQueue")
	fmt.Println(err)
	return err
}
func DeleteQueue(vhostname string) error {
	var rabbitadmin string = config.GetEnvbyKey("APPRABBITADMIN")
	var rabbitadminpwd string = config.GetEnvbyKey("APPRABBITADMINPASS")
	var rabbituri string = config.GetEnvbyKey("APPRABBITURL2")
	data := []byte(`{"auto_delete":false,"durable":true`)
	uri := strings.Join([]string{"http://", rabbitadmin, ":", rabbitadminpwd, "@", rabbituri, "api/queues/", vhostname, "/PLQueue"}, "")
	err := deleteData(uri, data)
	fmt.Println("DeleteQueue")
	fmt.Println(err)
	return err
}

func BindExchangeandQueue(vhostname string) error {
	var rabbitadmin string = config.GetEnvbyKey("APPRABBITADMIN")
	var rabbitadminpwd string = config.GetEnvbyKey("APPRABBITADMINPASS")
	var rabbituri string = config.GetEnvbyKey("APPRABBITURL2")
	data := []byte(`{"routing_key":"PLRoutingKey"}`)
	uri := strings.Join([]string{"http://", rabbitadmin, ":", rabbitadminpwd, "@", rabbituri, "api/bindings/vhost/", vhostname, "/e/PLExchange/q/PLQueue"}, "")
	err := postData(uri, data)
	fmt.Println("BindExchangeandQueue")
	fmt.Println(err)
	return err
}

func SetupUserandvHost(username string, vhostname string) error {
	err := CreateUser(username)
	if err != nil {
		return err
	}
	err = CreatevHosts(vhostname)
	if err != nil {
		return err
	}
	err = SetUserandvHostPermisssion(username, vhostname)
	if err != nil {
		return err
	}
	err = SetUserandvHostTopicPermisssion(username, vhostname)
	if err != nil {
		return err
	}
	err = CreateExchange(vhostname)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserandvHost(username string, vhostname string) error {
	err := DeleteExchange(vhostname)
	if err != nil {
		return err
	}
	err = DeletevHosts(vhostname)
	if err != nil {
		return err
	}
	err = DeleteUser(username)
	if err != nil {
		return err
	}
	return nil
}
