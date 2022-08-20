import { api_baseurl } from "../config";

export async function account(session) {

    const response = await fetch(
        api_baseurl + `/account`,
        {
            method: 'GET',
            mode: 'cors',
            cache: 'no-cache',
            headers: {
                'Session': session
            },
        }   
    )
    const data = await response.json();
    if (response.status === 200) {
        return data;
    }
    throw data["message"];
}