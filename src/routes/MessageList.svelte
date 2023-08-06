<script>
    import { onMount } from 'svelte';
    import {messages} from './stores.js'

    let messageList;

    messages.subscribe((value) => {
        messageList = value
    })

    onMount(async () => {
        const res = await fetch('http://127.0.0.1:8000/list');
        messageList = await res.json();
    });



</script>

<h2>Messages:</h2>
{#each messageList as msg}
    <div class="message">
         <p class="sender">{msg.sender}:</p>
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
