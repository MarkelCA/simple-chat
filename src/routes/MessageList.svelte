<script>
    import {messages} from './stores.js'

    let messageList;

    messages.subscribe((value) => {
        messageList = value
    })

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
