import { api_baseurl } from "../config";

export async function latestnews(pages) {

    const response = await fetch(
        api_baseurl + `/news/${pages}`,
        {
            method: 'GET',
            mode: 'cors',
            cache: 'no-cache'
        }   
    )
    const data = await response.json();
    if (response.status === 200) {
        return data;
    }
    throw data["message"];
}