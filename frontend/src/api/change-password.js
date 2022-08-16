import { auth_api } from "../config";

export async function change_password(oldPassword, newPassword) {
    const payload = {
        'old-password' : oldPassword,
        'new-password' : newPassword
    };
    const response = await fetch(
        auth_api + "/password",
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