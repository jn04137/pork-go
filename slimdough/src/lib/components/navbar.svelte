<script>
  import { Link } from 'svelte-routing'
  import { createQuery } from '@tanstack/svelte-query'
  import axios from 'axios'

  import AuthButton from '../components/login_signup_button.svelte'

  const isLoggedIn = async () => {
    try {
      const res = await axios.get("/auth/isloggedin")
      return res.data
    } catch(err) {
      console.error(err)
      return err.response.data
    }
  }

  const refetchIntervalMs = 300000;
  const query = createQuery({
    queryKey: ['isLoggedIn'],
    queryFn: () => isLoggedIn(),
    refetchInterval: refetchIntervalMs
  })

  export let navs = [
    { endpoint: '/', page: 'Home' },
    //{ endpoint: '/about', page: 'About' },
    { endpoint: '/contact', page: 'Contact' }
  ]

</script>

<div class='flex justify-center shadow-lg'>
  <div class='flex justify-between py-2 w-[1000px] items-center'>
    <div class='text-lg font-bold'><Link to="/">knostash</Link></div>
    <div class="flex space-x-4 items-center text-sm">
      <nav class='space-x-4'>
        {#each navs as nav}
          <Link to={nav.endpoint}>{nav.page}</Link>
        {/each}
      </nav>
      {#if $query.isSuccess}
        <AuthButton loggedIn={$query.data.isLoggedIn}/>
      {/if}
    </div>
  </div>
</div>

<style>

</style>

