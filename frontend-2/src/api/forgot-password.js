import { auth_api } from "../config";

export async function forgot_password(email){
    const payload = {
        'email': email
    };
    const response = await fetch(
        auth_api + "/reset/password",
        {
            method: 'POST',
            mode: 'cors',
            cache: 'no-cache',
            headers: {
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

export async function forgot_password_confirm(newPassword,confirmCode){
    const payload = {
        'new-password': newPassword
    };
    const response = await fetch(
        auth_api + "/confirm/reset/password",
        {
            method: 'POST',
            mode: 'cors',
            cache: 'no-cache',
            headers: {
                'Confirm-Code':confirmCode,
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