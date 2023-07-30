import { writable } from 'svelte/store';

let defaultMessages = [
    {"sender": "Lex", "content": "fizz"},
    {"sender": "Karmak", "content": "buzz"},
]

export const messages = writable(defaultMessages);
export const username = writable("user-" + Math.floor(Math.random() * 100000))
