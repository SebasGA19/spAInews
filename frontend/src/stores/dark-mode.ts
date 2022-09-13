import { writable } from 'svelte/store';
import { browser } from '$app/environment'

export const darkMode = writable((browser && (localStorage.getItem('dark-mode')) || 'light'));
export const unsubscribeDarkMode = darkMode.subscribe(value => browser && (localStorage.setItem('dark-mode', value)));