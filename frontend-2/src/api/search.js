import { api_baseurl } from "../config";

export async function search(target) {

    const response = await fetch(
        api_baseurl + `/search/${target}`,
        {
            method: 'GET',
            mode: 'cors',
            cache: 'no-cache',
            headers: {
                'Content-Type': 'application/json'
            },
            
        }
    )
    const data = await response.json();
    if (response.status === 200){
        return data;
    }
    throw data["message"];
}