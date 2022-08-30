<script>
import {session} from '../stores/session'

</script>

{#if $session === ''}
<script>
    window.location.href = "/";
</script>
{/if}