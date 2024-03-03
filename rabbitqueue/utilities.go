package rabbitqueue

import (
	"bytes"
	"net/http"
	"strings"
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

func CreateUser(username string) error {
	uri := strings.Join([]string{"http://guest:guest@localhost:15672/api/users/", username}, "")
	data := []byte(`{"password":"password","tags":"none"}`)
	err := sendData(uri, data)
	return err
}

func SetUserandvHostPermisssion(username string, vhostname string) error {
	uri := strings.Join([]string{"http://guest:guest@localhost:15672/api/permissions/", vhostname, "/", username}, "")
	data := []byte(`{"configure":".*","write":".*","read":".*"}`)
	err := sendData(uri, data)
	return err
}
func SetUserandvHostTopicPermisssion(username string, vhostname string) error {
	uri := strings.Join([]string{"http://guest:guest@localhost:15672/api/topic-permissions/", vhostname, "/", username}, "")
	data := []byte(`{"exchange":"amq.topic","write":"^a","read":".*"}`)
	err := sendData(uri, data)
	return err
}

func CreatevHosts(vhostname string) error {
	data := []byte(`{"description":"virtual host description", "tags":"accounts,production"}`)
	uri := strings.Join([]string{"http://guest:guest@localhost:15672/api/vhosts/", vhostname}, "")
	err := sendData(uri, data)
	return err
}

func CreateExchange(vhostname string) error {
	data := []byte(`{"type":"direct","auto_delete":false,"durable":true,"internal":false}`)
	uri := strings.Join([]string{"http://guest:guest@localhost:15672/api/exchanges/", vhostname, "/PLExchange"}, "")
	err := sendData(uri, data)
	return err
}

func CreateQueue(vhostname string) error {
	data := []byte(`{"auto_delete":false,"durable":true`)
	uri := strings.Join([]string{"http://guest:guest@localhost:15672/api/queues/", vhostname, "/PLQueue"}, "")
	err := sendData(uri, data)
	return err
}

func BindExchangeandQueue(vhostname string) error {
	data := []byte(`{"routing_key":"PLRoutingKey"}`)
	uri := strings.Join([]string{"http://guest:guest@localhost:15672/api/bindings/vhost/", vhostname, "/e/PLExchange/q/PLQueue"}, "")
	err := postData(uri, data)
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
