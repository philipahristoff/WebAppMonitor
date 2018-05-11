package main

import "net/http"

type Status int

const 
(
    DOWN
    UP
)

//Encapsulates the details of the target site.
type site struct
{
	url string
	lastStatus Status
}

//Sends GET request then checks if the resulting code is 200(UP) or anything else(DOWN).
//For future development we can save the status code directly instead of just the status itself.
//That way we'd be capable of monitoring the type of error in greater detail.
func (s Site) Status() (Status, error) 
{
    resp, err := http.Get(s.url)
    status := s.last_status

	if (err == nil) && (resp.StatusCode == 200) 
	{
        status = UP
	} 
	else 
	{
        status = DOWN
    }

    return status, err
}