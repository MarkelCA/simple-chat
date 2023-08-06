<script lang="ts">
    import { onMount } from 'svelte';
    import { get } from 'svelte/store';
    import {messages, username} from './stores.js'
    import io from 'socket.io-client';
    let socket;

    export let messageContent: string = ""

    onMount(() => {
        console.log('mounted')
        // Connect to your localhost server here
        socket = io('http://localhost:8080');
        console.log(socket)

        // Optional: Add event listeners for various Socket.io events
        socket.on('connect', () => {
            console.log('Connected to server');
        });

        socket.on('message', (data) => {
            console.log('Received message:', data);
        });
        socket.on("connect_error", (err) => {
            console.error(err);
        })
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

        messages.update((currentValue) => [...currentValue, message])
        messageContent = ""

        socket.emit('message', 'Hello from Svelte!');

    }
</script>

<input placeholder="Type your message..." bind:value={messageContent} />
<button on:click={sendMessage}>Send</button>
