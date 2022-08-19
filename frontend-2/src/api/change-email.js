import { auth_api } from "../config";

export async function change_email(password, newEmail, session) {
    const payload = {
        'password' : password,
        'new-email' : newEmail
    };
    const response = await fetch(
        auth_api + "/email",
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

export async function confirm_change_email(confirmCode)  {
    const response = await fetch(
        auth_api + "/confirm/email",
        {
            method: 'POST',
            mode: 'cors',
            cache: 'no-cache',
            headers: {
              'Confirm-Code': confirmCode
            },
        }
    );
    if (response.status === 200) {
        return;
    }
    const data = await response.json();
    throw data["message"];
}