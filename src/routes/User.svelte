<script>
    import { get } from 'svelte/store';
    import {username} from './stores.js'

    let defaultUsername = get(username);

    function changeUsername(newUsername) {
        username.set(newUsername)
    }
</script>

<span>User: </span><input on:input={(textfield) => changeUsername(textfield.target.value)} value={defaultUsername}/>

