import { auth_api } from "../config";

export async function confirmRegister(confirmCode)  {
    const response = await fetch(
        auth_api + "/confirm/registration",
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