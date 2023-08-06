<script lang="ts">
    import { onMount } from 'svelte';
    import { get } from 'svelte/store';
    import {messages, username} from './stores.js'
    export let messageContent: string = ""
    let socket

    onMount(() => {
        socket = new WebSocket(`ws://localhost:${import.meta.env.VITE_SERVER_PORT}/add`);
        socket.addEventListener("message", (event) => {
            messages.update((currentValue) =>{
                const newMessage = JSON.parse(event.data)
                return [...currentValue, newMessage]
            })
        });

        socket.addEventListener("error", (event) => {
            console.error("Error from server:", event.data);
        });

        document.getElementById("message-input").focus();
    });

    function sendMessage() {
        if (messageContent == "" || messageContent == undefined) {
            alert("The message content can't be null")
            return
        }

        const message = {
            'sender' : get(username),
            'content': messageContent
        }

        socket.send(JSON.stringify(message));
    }
</script>

<form id="message-form">
    <input id="message-input" placeholder="Type your message..." bind:value={messageContent} />
    <button on:click={sendMessage}>Send</button>
</form>
