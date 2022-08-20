import { api_baseurl } from "../config";

export async function login(username, password) {
    const headers = new Headers();
    headers.set('Authorization', 'Basic ' + btoa(username + ":" + password));
    
    const response = await fetch(api_baseurl + "/session",{method:'GET',headers: headers});
    const data = await response.json();
    if (response.status === 200) {
        return data["session"];
    }
    throw data["message"];
}