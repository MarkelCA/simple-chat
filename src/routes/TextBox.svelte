<script lang="ts">
    import { get } from 'svelte/store';
    import {messages, username} from './stores.js'

    export let messageContent: string

    function sendMessage() {
        console.log(messageContent)
        if (messageContent == "" || messageContent == undefined) {
            console.log("saonteu")
            alert("The message content can't be null")
            return
        }
        const message = {
            'sender' : get(username),
            'content': messageContent
        }
        messages.update((currentValue) => [...currentValue, message])
        messageContent = ""
    }
</script>

<input placeholder="Type your message..." bind:value={messageContent} />
<button on:click={sendMessage}>Send</button>
