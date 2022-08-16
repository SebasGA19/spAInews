import { auth_api } from "../config";

export async function register(username, email, password) {
    const payload = {
        username: username,
        email: email,
        password: password
    };
    const response = await fetch(
        auth_api + "/register",
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