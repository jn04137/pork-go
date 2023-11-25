<script>
  import Layout from '../lib/components/layout.svelte'
  import axios from 'axios';
  import { navigate } from 'svelte-routing'

  $: login = {
    "username": "",
    "password": ""
  }

  $: signup = {
    "username": "",
    "email": "",
    "password": "",
    "passwordMatch": ""
  }

  const handleLogin = async () => {
    try {
      await axios.post("/auth/login", login, {
        withCredentials: true
      })
      navigate("/", { replace: true })
    } catch(e) {
      console.error(e)
      alert("Credentials failed")
    }
  }

  const handleSignup = async () => {
    try {
      if(signup.password != signup.passwordMatch) {
        alert("Passwords do not match")
        throw new Error("Passwords do not match")
      }
      await axios.post("/auth/signup", signup, {
        withCredentials: true
      })
    } catch(e) {
      console.error(e)
    }

    // Should navigate to a page where it verifies that a confirmation email was sent
    navigate("/", { replace: true})
  }

  // TODO Have to add matching signupPassword logic before making the signup request 

  const inputClass = "border border-grey-300 px-2 py-2 rounded"
  const buttonClass = "bg-blue-500 rounded-2xl text-white py-1"
</script>

<Layout>
  <body slot="content" class="flex justify-center w-full">
  <div class="flex flex-col justify-center">
    <div class="bg-white rounded space-y-10 px-8 py-8 h-fit w-[350px]">
      <form class="flex flex-col space-y-3" on:submit|preventDefault={handleLogin}>
        <h1 class="text-2xl">Login</h1>
        <input 
          type="text" 
          placeholder="username"
          name="username"
          class={inputClass}
          bind:value={login.username}
        /> 
        <input 
          type="password" 
          placeholder="password"
          name="password"
          class={inputClass}
          bind:value={login.password}
        /> 
        <button class={buttonClass} type="submit">Login</button>
      </form>
      <form class="flex flex-col space-y-3" on:submit|preventDefault={handleSignup}>
        <h1 class="text-2xl">Signup</h1>
        <input 
          type="text" 
          placeholder="email" 
          class={inputClass}
          bind:value={signup.email}
        /> 
        <input 
          type="text" 
          placeholder="username" 
          class={inputClass}
          bind:value={signup.username}
        /> 
        <input 
          type="password" 
          placeholder="password" 
          class={inputClass}
          bind:value={signup.password}
        /> 
        <input 
          type="password" 
          placeholder="match password" 
          class={inputClass}
          bind:value={signup.passwordMatch}
        /> 
        <button class={buttonClass} type="submit">Signup</button>
      </form>
  </div>
  </body>
</Layout>

