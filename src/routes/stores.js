import { writable } from 'svelte/store';


export const messages = writable([]);
export const username = writable("user-" + Math.floor(Math.random() * 100000))
