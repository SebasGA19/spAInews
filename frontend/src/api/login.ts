import {api_baseurl} from "../config";

export async function login(username: string, password: string): Promise<string> {
    const headers = new Headers();
    headers.set('Authorization', 'Basic ' + btoa(username + ":" + password));

    const response = await fetch(api_baseurl + "/session", {method: 'GET', headers: headers});
    const data = await response.json();
    if (response.status === 200) {
        return data["session"];
    }
    throw data["message"];
}

export async function logout(session: string) {
    const response = await fetch(
        api_baseurl + "/session",
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