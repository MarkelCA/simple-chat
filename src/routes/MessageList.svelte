<script>
    import { onMount } from 'svelte';
    import {messages} from './stores.js'

    let messageList;

    messages.subscribe((value) => {
        messageList = value
    })

    onMount(async () => {
        const res = await fetch(`http://127.0.0.1:${import.meta.env.VITE_SERVER_PORT}/list`);
        messageList = await res.json();
        messages.set(messageList)
    });



</script>

<h2>Messages:</h2>
{#each messageList as msg}
    <div class="message">
         <p><span class="sender">{msg.sender}</span> (<span class="msg-time">{msg.time}</span>):</p>
         <p class="msg-content">{msg.content}</p>
    </div>
{/each}

<style>
    .sender {
        font-weight: bold
    }
    .message {
        margin: 10px
    }
    .message p {
        margin: 1px
    }
</style>
