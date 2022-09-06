import type { News } from "src/common/types";
import { api_baseurl } from "../config";

export async function news(pages: number): Promise<News> {

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

export async function search(page: number): Promise<News> {

    const response = await fetch(
        api_baseurl + `/search/${page}`,
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