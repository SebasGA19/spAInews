import { api_baseurl } from "../config";

export async function change_email(password, newEmail) {
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
                'Session': localStorage.getItem('session'),
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
        api_baseurl + "/confirm/email",
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