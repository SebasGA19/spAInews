import { api_baseurl } from "../config";

export async function register(username, email, password) {
    const payload = {
        username: username,
        email: email,
        password: password
    };
    const response = await fetch(
        api_baseurl + "/register",
        {
            method: 'PUT',
            mode: 'cors',
            cache: 'no-cache',
            headers: {
              'Content-Type': 'application/json'
            },
            body: JSON.stringify(payload)
        }
    );
    if (response.status === 200) {
        return;
    }
    const data = await response.json();
    throw data["message"];
}