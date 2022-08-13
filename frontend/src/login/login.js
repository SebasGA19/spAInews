import React from "react";
import { auth_api, news_api } from "../config";

/* Login */
export function login(username, password) {
    let headers = new Headers();
    headers.set('Authorization', 'Basic ' + btoa(username + ":" + password));
    fetch(auth_api + "/session",{method:'GET',headers: headers,})
    .then((response) => {
        return response.json()
      })
      .then((session) => {
        console.log(session)
      })
  }
 
