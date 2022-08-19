import { auth_api } from "../config";

export async function login(username, password) {
    const headers = new Headers();
    headers.set('Authorization', 'Basic ' + btoa(username + ":" + password));
    
    const response = await fetch(auth_api + "/session",{method:'GET',headers: headers});
    const data = await response.json();
    if (response.status === 200) {
        return data["session"];
    }
    throw data["message"];
}

export async function logout(session) {
    const response = await fetch(
        auth_api + "/session",
        {
            method: 'DELETE',
            mode: 'cors',
            cache: 'no-cache',
            headers: {
              'Session': session
            },
        }
    );
    if (response.status === 200) {
        return;
    }
    const data = await response.json();
    throw data["message"];
}