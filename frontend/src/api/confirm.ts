import {api_baseurl} from "../config";

export async function confirmRegister(confirmCode: string) {
    const response = await fetch(
        api_baseurl + "/confirm/registration",
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

export async function confirmEmail(confirmCode: string) {
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