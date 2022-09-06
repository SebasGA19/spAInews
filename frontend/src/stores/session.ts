import {writable} from 'svelte/store';
import {browser} from '$app/environment'

export const session: any = writable(browser && (localStorage.getItem('session') || ''));
export const unsubscribeSession = session.subscribe(value => browser && (localStorage.setItem('session', value)));

export const sessionUsername: any = writable(browser && (localStorage.getItem('sessionUsername') || ''));
export const unsubscribeSessionUsername = sessionUsername.subscribe(value => browser && (localStorage.setItem('sessionUsername', value)));