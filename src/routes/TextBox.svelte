<script lang="ts">
    import { onMount } from 'svelte';
    import { get } from 'svelte/store';
    import {messages, username} from './stores.js'
    export let messageContent: string = ""

    let socket


    onMount(() => {
        console.log(WebSocket)
        socket = new WebSocket("ws://localhost:8080/add");
        socket.addEventListener("message", (event) => {
            messages.update((currentValue) =>{
                const newMessage = JSON.parse(event.data)
                return [...currentValue, newMessage]
            })
        });

        socket.addEventListener("error", (event) => {
            console.error("Error from server:", event.data);
        });

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

        // messages.update((currentValue) => [...currentValue, message])
        // messageContent = ""
        //
        // socket.emit('message', 'Hello from Svelte!');

    }
</script>

<input placeholder="Type your message..." bind:value={messageContent} />
<button on:click={sendMessage}>Send</button>
