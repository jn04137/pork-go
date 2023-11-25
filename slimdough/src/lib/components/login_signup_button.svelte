<script>
  import { Link } from 'svelte-routing'
  export let loggedIn
</script>

{#if loggedIn === true}
    <Link to="/myprofile">
      My Profile
    </Link>
{:else}
    <Link to="/login">
      <button class="bg-blue-500 text-white px-2 py-0.5 rounded-2xl shadow">Login / Signup</button>
    </Link>
{/if}
