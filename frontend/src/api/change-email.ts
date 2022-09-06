import { api_baseurl } from "../config";

export async function updateEmail(password: string, newEmail: string, session: string) {
    const payload = {
        'password' : password,
        'new-email' : newEmail
    };
    const response = await fetch(
        api_baseurl + "/email",
        {
            method: 'POST',
            mode: 'cors',
            cache: 'no-cache',
            headers: {
                'Session': session,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(payload)
        }
    );
    if (response.status === 200){
        return;
    }
    const data = await response.json();
    throw data["message"];
}
