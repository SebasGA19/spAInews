import {error} from "@sveltejs/kit"

export async function load({params, data}) {
    if (params.code !== undefined) {
        return {code: params.code}
    }
    throw error('403', "No reset code provided")
}